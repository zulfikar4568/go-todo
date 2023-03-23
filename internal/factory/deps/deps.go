package deps

import (
	"sync"

	"github.com/fgrosse/goldi"
	Config "github.com/zulfikar4568/go-todo/internal/config"
)

var (
	once     = &sync.Once{}
	instance = &goldi.Container{}
)

func NewDependencyFactory(config map[string]interface{}) *goldi.Container {
	registry := goldi.NewTypeRegistry()

	once.Do(func() {
		instance = goldi.NewContainer(registry, config)

		instance.RegisterType("app.config", Config.NewImmutableConfig)
	})

	return instance
}
