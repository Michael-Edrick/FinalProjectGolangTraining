package handler

import (
	"FinalProject/entity"
	"FinalProject/mapper"
	"FinalProject/middleware"
	"FinalProject/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	r              *mux.Router
	commentService entity.CommentServiceInterface
}

func NewCommentHandler(r *mux.Router, commentService entity.CommentServiceInterface) {
	handler := CommentHandler{
		r:              r,
		commentService: commentService,
	}
	s := r.PathPrefix("").Subrouter()
	s.Use(middleware.IsAuthorized())
	s.HandleFunc("/comments", handler.commentPostGetHandler)
	s.HandleFunc("/comments/{Id}", handler.commentUpdateDeleteHandler)
}

func (h CommentHandler) commentPostGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var postComment *entity.Comment
		err := decoder.Decode(&postComment)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//get claims
		header := r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		postComment.UserId = int(claims["userid"].(float64))
		//masuk ke photo service
		res, err := h.commentService.CommentPostService(postComment)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//keluarin response
		response := mapper.PostCommentMapper(res)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(utils.Response(response, err))
	}
	if r.Method == "GET" {
		//get claims
		header := r.Header.Get("Authorization")
		claims, err := utils.ParseJWT(header)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//masuk ke comment service
		var getComments entity.Comment
		getComments.UserId = int(claims["userid"].(float64))
		res, err := h.commentService.CommentGetService(&getComments)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//keluarin response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(utils.Response(res, err))
	}
}

func (h CommentHandler) commentUpdateDeleteHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["Id"]
	if r.Method == "PUT" {
		//baca dr body
		decoder := json.NewDecoder(r.Body)
		var updateComment entity.Comment
		err := decoder.Decode(&updateComment)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Response("", err))
			return
		}
		//convert param
		if id != "" {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(utils.Response("", err))
				return
			} else {
				updateComment.Id = idInt
				//masuk ke comment service
				res, err := h.commentService.CommentUpdateService(&updateComment)
				if err != nil {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					w.Write(utils.Response("", err))
					return
				}
				//keluarin response
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(utils.Response(res, err))
			}
		}
	}
	if r.Method == "DELETE" {
		var deleteComment entity.Comment
		//convert param
		if id != "" {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(utils.Response("", err))
				return
			} else {
				deleteComment.Id = idInt
				//masuk ke comment service
				err = h.commentService.CommentDeleteService(&deleteComment)
				if err == nil {
					response := map[string]string{
						"message": "Your comment has been successfully deleted",
					}
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(utils.Response(response, err))
					return
				}
			}
		}
	}
}
