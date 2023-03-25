package dto

import "github.com/gofrs/uuid"

type (
	CreateTodoRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	CreateTodoResponse struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
	}

	GetTodoResponse struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
	}

	UpdateTodoRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	UpdateTodoResponse struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
	}
)
