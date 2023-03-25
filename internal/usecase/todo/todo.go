package todo

import (
	"github.com/gofrs/uuid"
	"github.com/zulfikar4568/go-todo/internal/model/dto"
	"github.com/zulfikar4568/go-todo/internal/model/entity"
	"github.com/zulfikar4568/go-todo/internal/repository/db"
	successresponse "github.com/zulfikar4568/go-todo/pkg/util/response/success_response"
)

type (
	usecase struct {
		Repository db.IDBRepository
	}

	IUsecaseTodo interface {
		List(page uint32, limit uint32) (*[]entity.Todo, *successresponse.ListMeta, error)
		Create(body *dto.CreateTodoRequest) (*dto.CreateTodoResponse, error)
		GetByID(id uuid.UUID) (*dto.GetTodoResponse, error)
		UpdateByID(id uuid.UUID, body *dto.UpdateTodoRequest) (*dto.UpdateTodoResponse, error)
	}
)

func (u *usecase) UpdateByID(id uuid.UUID, body *dto.UpdateTodoRequest) (*dto.UpdateTodoResponse, error) {
	query := entity.Todo{
		ID: id,
	}

	result, err := u.Repository.GetRepositoryTodo().Update(&query, &entity.Todo{Name: body.Name, Description: body.Description})
	if err != nil {
		return nil, err
	}

	return &dto.UpdateTodoResponse{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
	}, nil
}

func (u *usecase) GetByID(id uuid.UUID) (*dto.GetTodoResponse, error) {
	query := entity.Todo{
		ID: id,
	}

	result, err := u.Repository.GetRepositoryTodo().Get(&query)
	if err != nil {
		return nil, err
	}

	return &dto.GetTodoResponse{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
	}, nil
}

func (u *usecase) Create(body *dto.CreateTodoRequest) (*dto.CreateTodoResponse, error) {
	id, _ := uuid.NewV4()

	data := &entity.Todo{
		ID:          id,
		Name:        body.Name,
		Description: body.Description,
	}

	result, err := u.Repository.GetRepositoryTodo().Create(data)
	if err != nil {
		return nil, err
	}

	return &dto.CreateTodoResponse{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
	}, nil
}

func (u *usecase) List(page uint32, limit uint32) (*[]entity.Todo, *successresponse.ListMeta, error) {
	todos, meta, err := u.Repository.GetRepositoryTodo().List(page, limit)
	if err != nil {
		return nil, nil, err
	}

	// Bussines Logic

	return todos, meta, nil
}

func NewUsecaseTodo(r db.IDBRepository) IUsecaseTodo {
	return &usecase{r}
}
