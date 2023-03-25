package db

import (
	"github.com/zulfikar4568/go-todo/internal/config"
	"github.com/zulfikar4568/go-todo/internal/repository/db/todo"
	"gorm.io/gorm"
)

type (
	repository struct {
		config config.IImmutableConfig
		db     *gorm.DB
		Todo   todo.IRepositoryTodo
	}

	IDBRepository interface {
		GetConn() *gorm.DB
		GetRepositoryTodo() todo.IRepositoryTodo
	}
)

func (r *repository) GetConn() *gorm.DB {
	return r.db
}

func (r *repository) GetRepositoryTodo() todo.IRepositoryTodo {
	return r.Todo
}

func NewRepository(cfg config.IImmutableConfig, conn *gorm.DB) IDBRepository {
	return &repository{
		config: cfg,
		db:     conn,
		Todo:   todo.NewRepositoryTodo(conn),
	}
}
