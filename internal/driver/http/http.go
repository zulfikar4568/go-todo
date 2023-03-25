package http

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/fgrosse/goldi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zulfikar4568/go-todo/internal/config"
	"github.com/zulfikar4568/go-todo/internal/delivery/rest"
	repository "github.com/zulfikar4568/go-todo/internal/repository/db"
	"github.com/zulfikar4568/go-todo/internal/usecase"
	"github.com/zulfikar4568/go-todo/pkg/log"
	errorresponse "github.com/zulfikar4568/go-todo/pkg/util/response/error_response"
	"github.com/zulfikar4568/go-todo/pkg/util/validator"
	"gorm.io/gorm"
)

type (
	Driver struct {
		ImmutableConfig config.IImmutableConfig
	}

	IDriverHTTP interface {
		Start(db *gorm.DB) (*echo.Echo, error)
	}
)

var (
	once     = &sync.Once{}
	logger   = log.NewAppLogger("[driver][http]")
	instance = &echo.Echo{}
)

func (d *Driver) Start(db *gorm.DB) (*echo.Echo, error) {
	var e error = nil

	if db == nil {
		logger.Error("missing dependencies database!")

		return nil, fmt.Errorf("dependencies connection is undefined")
	}

	once.Do(func() {
		instance = echo.New()

		// Setup
		instance.Pre(middleware.AddTrailingSlash())
		instance.Validator = validator.New()
		instance.HTTPErrorHandler = customErrorHandler

		// Default Routes
		instance.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, d.ImmutableConfig.GetServiceName())
		})

		// Delivery
		rest.NewDelivery(d.getUseCase(db)).Load(instance)

		// start http server
		err := instance.Start(fmt.Sprintf(":%d", d.ImmutableConfig.GetServiceHTTPPort()))
		if err != nil {
			logger.Error("failed to start http server", err)

			e = err
		}
	})

	return instance, e
}

func (d *Driver) getUseCase(db *gorm.DB) *usecase.Usecase {
	registry := goldi.NewTypeRegistry()
	container := goldi.NewContainer(registry, make(map[string]interface{}))

	container.RegisterType("app.http.repository", repository.NewRepository, d.ImmutableConfig, db)
	container.RegisterType("app.http.usecase", usecase.NewUsecase, d.ImmutableConfig, "@app.http.repository")

	return container.MustGet("app.http.usecase").(*usecase.Usecase)
}

func customErrorHandler(err error, ctx echo.Context) {
	httpStatus := http.StatusInternalServerError
	errorData := errorresponse.NewErrorResponse("internal server error", errorresponse.ErrorData{Code: httpStatus, Params: err.Error()})

	echoError := &echo.HTTPError{}

	if errors.As(err, &echoError) {
		httpStatus = echoError.Code
		errorData = errorresponse.NewErrorResponse("internal server error", errorresponse.ErrorData{Code: httpStatus, Params: echoError})
	}

	_ = ctx.JSON(httpStatus, errorData)
}

// NewDriver is a constructor that return interface of http driver
func NewDriver(cfg config.IImmutableConfig) IDriverHTTP {
	return &Driver{
		cfg,
	}
}
