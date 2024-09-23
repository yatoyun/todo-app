package usecase

import "time"

// output data <DS>
type TodoOutputData struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// output Boundary <I>
type TodoOutputPortInterface interface {
	Convert(TodoOutputData) (*TodoResponse, error)
}
