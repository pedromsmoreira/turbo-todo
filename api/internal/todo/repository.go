package todo

import (
	"errors"
)

type TodoRepository interface {
	Get(id string) (*todo, error)
	Add(todo *todo) (*todo, error)
}

type inMemoryTodoRepository struct {
	todos map[string]*todo
}

func NewInMemoryTodoRepository() TodoRepository {
	return &inMemoryTodoRepository{
		todos: map[string]*todo{},
	}
}

func (imr *inMemoryTodoRepository) Get(id string) (*todo, error) {
	return nil, errors.New("not implemented")
}

func (imr *inMemoryTodoRepository) Add(todo *todo) (*todo, error) {
	return nil, errors.New("not implemented")
}

type CockroachDbTodoRepository struct{}

func NewCockroachDbTodoRepository(username string, password string) TodoRepository {
	return &CockroachDbTodoRepository{}
}
func (imr *CockroachDbTodoRepository) Get(id string) (*todo, error) {
	return nil, errors.New("not implemented")
}

func (imr *CockroachDbTodoRepository) Add(todo *todo) (*todo, error) {
	return nil, errors.New("not implemented")
}
