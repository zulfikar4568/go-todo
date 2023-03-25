package todo

import (
	"github.com/zulfikar4568/go-todo/internal/model/entity"
	"github.com/zulfikar4568/go-todo/pkg/util/pagination"
	successresponse "github.com/zulfikar4568/go-todo/pkg/util/response/success_response"
	"gorm.io/gorm"
)

type (
	repository struct {
		db *gorm.DB
	}

	IRepositoryTodo interface {
		List(page uint32, limit uint32) (*[]entity.Todo, *successresponse.ListMeta, error)
		Create(data *entity.Todo) (*entity.Todo, error)
		Get(query *entity.Todo) (*entity.Todo, error)
		Update(query *entity.Todo, data *entity.Todo) (*entity.Todo, error)
	}
)

func (r *repository) Update(query *entity.Todo, data *entity.Todo) (*entity.Todo, error) {
	exec := r.db.Debug()
	err := exec.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&query).Find(&data).Error; err != nil {
			return err
		}

		if err := tx.Save(&data).Error; err != nil {
			return err
		}

		return nil
	})

	return data, err
}

func (r *repository) Get(query *entity.Todo) (*entity.Todo, error) {
	result := entity.Todo{}

	if err := r.db.Debug().Where(&query).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repository) Create(data *entity.Todo) (*entity.Todo, error) {
	if err := r.db.Debug().Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) List(page uint32, limit uint32) (*[]entity.Todo, *successresponse.ListMeta, error) {
	result := []entity.Todo{}
	meta := successresponse.ListMeta{}
	var total int64

	exec := r.db.Debug()

	err := exec.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&entity.Todo{}).Count(&total).Error; err != nil {
			return err
		}

		if page != 0 && limit != 0 {
			if err := tx.Limit(int(limit)).Offset(pagination.GetOffset(int(page), int(limit))).Find(&result).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Find(&result).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		panic("Transaction failed!")
	}

	count, totalPage := pagination.ParseMeta(int(total), int(limit), result)

	if total > 0 {
		meta = successresponse.ListMeta{
			Count:     count,
			Total:     int(total),
			Page:      int(page),
			TotalPage: totalPage,
		}
	}

	return &result, &meta, nil
}

func NewRepositoryTodo(conn *gorm.DB) IRepositoryTodo {
	return &repository{conn}
}
