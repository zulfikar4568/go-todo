package baseresponse

type (
	BaseResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Result  any    `json:"result"`
		Error   any    `json:"error"`
		Meta    any    `json:"meta"`
	}

	IBaseResponse interface {
		GetResult() any
		GetError() any
	}
)

func (r *BaseResponse) GetResult() any {
	return r.Result
}

func (r *BaseResponse) GetError() any {
	return r.Error
}

func NewBaseResponse(success bool, message string, result any, errorVal any, meta any) *BaseResponse {
	return &BaseResponse{
		Success: success,
		Message: message,
		Result:  result,
		Error:   errorVal,
		Meta:    meta,
	}
}
