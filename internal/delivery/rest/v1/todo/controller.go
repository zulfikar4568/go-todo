package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zulfikar4568/go-todo/internal/usecase"
	baseResponse "github.com/zulfikar4568/go-todo/pkg/util/response/base_response"
	errorResponse "github.com/zulfikar4568/go-todo/pkg/util/response/error_response"
	successResponse "github.com/zulfikar4568/go-todo/pkg/util/response/success_response"
)

type (
	Delivery struct {
		Usecase *usecase.Usecase
	}
)

func (d *Delivery) Route(g *echo.Group) {
	g.GET("/", d.List)
}

func (d *Delivery) List(ctx echo.Context) error {
	result, meta, err := d.Usecase.Todo.List(1, 10)
	var msg string
	var res baseResponse.BaseResponse

	if err != nil {
		msg = "Failed Get List"
		res = *errorResponse.NewErrorResponse(msg, errorResponse.ErrorData{Code: 401, Message: msg, Params: err.Error()})
		return ctx.JSON(http.StatusForbidden, res)
	} else {
		msg = "Success Get All Data"
	}

	res = *successResponse.NewSuccessResponse(msg, result, meta)

	return ctx.JSON(http.StatusOK, res)
}

func NewDelivery(u *usecase.Usecase) *Delivery {
	return &Delivery{
		Usecase: u,
	}
}
