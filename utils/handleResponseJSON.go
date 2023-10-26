package utils

import "encoding/json"

type Response struct {
	Message string `json:"message"`
}

func HandleResponseJSON(message string) []byte {
	result, _ := json.Marshal(Response{Message: message})
	return result
}
