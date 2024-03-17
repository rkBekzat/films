package handler

import (
	"encoding/json"
	"net/http"
)

func sendErr(w http.ResponseWriter, code int, err error) {
	resp := StatusResponse{
		Success: false,
		Data:    err.Error(),
	}
	w.WriteHeader(code)
	res, _ := formResponse(resp)
	w.Write(res)
}

type StatusResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func formResponse(response interface{}) ([]byte, error) {
	responseStatus := StatusResponse{
		Success: true,
		Data:    response,
	}
	res, err := json.Marshal(responseStatus)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func sendResponse(data interface{}, w http.ResponseWriter) error {
	res, err := formResponse(data)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return nil
}
