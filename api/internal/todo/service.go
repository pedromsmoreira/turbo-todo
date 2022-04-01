package todo

type TodoService struct {
	tcr TodoRepository
}

func NewTodoService(tcr TodoRepository) *TodoService {
	return &TodoService{
		tcr: tcr,
	}
}
