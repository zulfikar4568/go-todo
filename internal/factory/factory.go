package factory

import (
	"fmt"

	Config "github.com/zulfikar4568/go-todo/internal/config"
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

	fmt.Println(config.GetServiceName() + " Loaded!")
}

func NewFactory() IFactory {
	config := make(map[string]interface{})
	return &appFactory{config}
}
