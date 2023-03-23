package deps

import (
	"sync"

	"github.com/fgrosse/goldi"
	Config "github.com/zulfikar4568/go-todo/internal/config"
	"github.com/zulfikar4568/go-todo/pkg/log"
)

var (
	once     = &sync.Once{}
	instance = &goldi.Container{}
	logger   = log.NewAppLogger("[factory][deps]")
)

func NewDependencyFactory(config map[string]interface{}) *goldi.Container {
	registry := goldi.NewTypeRegistry()

	once.Do(func() {
		instance = goldi.NewContainer(registry, config)

		instance.RegisterType("app.config", Config.NewImmutableConfig)

		logger.Info("success build dependencies injection")
	})

	return instance
}
