package todos

import (
	"time"

	"github.com/google/uuid"
)

type todo struct {
	Id          string
	Attributes  *attributes
	DateCreated string
	Version     int64
}

type attributes struct {
	Title   string
	Content string
	Tags    []string
	Status  string
}

// apply builder pattern
func NewTodo() *todo {
	id := uuid.New().String()
	date := time.Now().String()
	version := 1
	a := &attributes{}

	return &todo{
		Id:          id,
		DateCreated: date,
		Version:     int64(version),
		Attributes:  a,
	}
}

type TodoBuilder struct {
	todo   *todo
	errors []error
}

func NewTodoBuilder() *TodoBuilder {
	return &TodoBuilder{todo: NewTodo()}
}

type AttributesBuilder struct {
	TodoBuilder
}

func (tb *TodoBuilder) WithAttributes() *AttributesBuilder {
	return &AttributesBuilder{*tb}
}

func (atb *AttributesBuilder) WithTitle(t string) *AttributesBuilder {
	atb.todo.Attributes.Title = t
	return atb
}

func (atb *AttributesBuilder) WithContent(c string) *AttributesBuilder {
	atb.todo.Attributes.Content = c
	return atb
}

func (atb *AttributesBuilder) WithTag(t string) *AttributesBuilder {
	atb.todo.Attributes.Tags = append(atb.todo.Attributes.Tags, t)
	return atb
}

func (atb *AttributesBuilder) WithStatus(s string) *AttributesBuilder {
	atb.todo.Attributes.Status = s
	return atb
}

func (td *TodoBuilder) Build() (*todo, []error) {
	if len(td.errors) != 0 {
		return nil, td.errors
	}

	return td.todo, nil
}
