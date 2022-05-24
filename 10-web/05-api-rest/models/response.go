package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Data          interface{} `json:"data"`
	Message       string      `json:"message"`
	contentType   string
	responseWrite http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		responseWrite: w,
		contentType:   "application/json",
	}
}

func (resp *Response) Send() {
	resp.responseWrite.Header().Set("Content-Type", resp.contentType)
	resp.responseWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.responseWrite, string(output))
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Message = "OK"
	response.Send()
}

func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource not found"
}

func SendNotFound(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity not found"
}

func SendUnprocessableEntity(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}
