package usecase

import "time"

// output data <DS>
type TodoOutputData struct {
	Id string
	Title string
	Description string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoResponse struct {
	Id string
	Title string
	Description string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// output Boundary <I>
type TodoOutputPortInterface interface {
	Convert(TodoOutputData) (*TodoResponse, error)
}
