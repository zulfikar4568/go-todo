package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/zulfikar4568/go-todo/internal/delivery/rest/v1/todo"
	"github.com/zulfikar4568/go-todo/internal/usecase"
)

type (
	Delivery struct {
		Usecase *usecase.Usecase
	}

	IDeliveryRest interface {
		Load(e *echo.Echo)
	}
)

func (d *Delivery) Load(e *echo.Echo) {
	api := e.Group("/api")

	v1 := api.Group("/v1")

	todoV1 := v1.Group("/todo")

	todo.NewDelivery(d.Usecase).Route(todoV1)
}

func NewDelivery(u *usecase.Usecase) IDeliveryRest {
	return &Delivery{u}
}
