package utils

import "encoding/json"

func Response(data any, err error) []byte {
	response := data
	if err != nil {
		response = map[string]any{
			"data":  data,
			"error": err.Error(),
		}
	}
	jsonResponse, _ := json.Marshal(response)
	return jsonResponse
}
