package usecase

import (
	"github.com/zulfikar4568/go-todo/internal/config"
	"github.com/zulfikar4568/go-todo/internal/repository/db"
	"github.com/zulfikar4568/go-todo/internal/usecase/todo"
)

type (
	Usecase struct {
		Config       config.IImmutableConfig
		DBRepository db.IDBRepository
		Todo         todo.IUsecaseTodo
	}
)

func NewUsecase(
	cfg config.IImmutableConfig,
	dbRepository db.IDBRepository,
) *Usecase {
	return &Usecase{
		Config:       cfg,
		DBRepository: dbRepository,
		Todo:         todo.NewUsecaseTodo(dbRepository),
	}
}
