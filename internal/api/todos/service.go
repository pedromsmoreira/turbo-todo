package todos

import "errors"

type TodoService struct {
	tcr TodoRepository
}

func NewTodoService(tcr TodoRepository) *TodoService {
	return &TodoService{
		tcr: tcr,
	}
}

func (tdsvc *TodoService) Get(id string) (*todo, error) {
	return nil, errors.New("not implemented")
}
