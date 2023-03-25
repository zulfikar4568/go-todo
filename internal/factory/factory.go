package factory

import (
	Config "github.com/zulfikar4568/go-todo/internal/config"
	Http "github.com/zulfikar4568/go-todo/internal/driver/http"
	Postgres "github.com/zulfikar4568/go-todo/internal/driver/postgres"
	"github.com/zulfikar4568/go-todo/internal/factory/deps"
)

type (
	appFactory struct {
		depsConfig map[string]interface{}
	}

	IFactory interface {
		Load()
	}
)

func (f *appFactory) Load() {
	dependencies := deps.NewDependencyFactory(f.depsConfig)

	config := dependencies.MustGet("app.config").(Config.IImmutableConfig)
	http := dependencies.MustGet("app.http").(Http.IDriverHTTP)
	postgres := dependencies.MustGet("app.postgres").(Postgres.IDriverPostgres)

	pi, err := postgres.Start()
	if err != nil {
		panic("cannot start db connection")
	}

	_, err = http.Start(pi[config.GetDatabaseConfig()[0].ID])
	if err != nil {
		panic("cannot start http server")
	}
}

func NewFactory() IFactory {
	config := make(map[string]interface{})
	return &appFactory{config}
}
