package model

import "net/http"

type BaseResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	Error      []string
}

func (b BaseResponse) InternalServerError(err string) BaseResponse {
	b.StatusCode = http.StatusInternalServerError
	b.Message = string("internal server error")
	b.Data = false
	b.Error = []string{err}

	return b
}
