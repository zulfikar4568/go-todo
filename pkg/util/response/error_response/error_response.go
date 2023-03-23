package error_response

import (
	"github.com/zulfikar4568/go-todo/pkg/util/response/base_response"
)

type (
	ErrorData struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Params  any    `json:"params"`
	}
)

func NewErrorResponse(message string, errorVal ErrorData) *base_response.BaseResponse {
	return &base_response.BaseResponse{
		Success: false,
		Message: message,
		Result:  map[string]string{},
		Error:   errorVal,
		Meta:    nil,
	}
}
