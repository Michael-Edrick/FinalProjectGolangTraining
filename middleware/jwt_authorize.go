package middleware

import (
	"FinalProject/utils"
	"net/http"
)

func IsAuthorized ()func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request){
			header:= r.Header.Get("Authorization")
			
			claims, err := utils.ParseJWT(header)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized"))
				return
			}
			if claims["email"] != ""{
				handler.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized"))
				return 
			}
		})
	}
}

