package successresponse

import baseresponse "github.com/zulfikar4568/go-todo/pkg/util/response/base_response"

type (
	ListMeta struct {
		Count     int `json:"count"`
		Total     int `json:"total"`
		Page      int `json:"page"`
		TotalPage int `json:"totalPage"`
	}
)

func NewSuccessResponse[M any | ListMeta](message string, result any, meta M) *baseresponse.BaseResponse {
	return &baseresponse.BaseResponse{
		Success: true,
		Message: message,
		Result:  result,
		Error:   nil,
		Meta:    meta,
	}
}
