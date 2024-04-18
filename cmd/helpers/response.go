package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseWithData struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

func Response(w http.ResponseWriter, code int, message any, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "Request Success"

	if code >= 400 {
		status = "Request Failed"
	}

	if payload != nil {
		response = &ResponseWithData{
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
