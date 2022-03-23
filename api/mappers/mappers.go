package mappers

import (
	"errors"

	"github.com/pedromsmoreira/turbo-todo/api/data"
	"github.com/pedromsmoreira/turbo-todo/api/dto"
	"github.com/pedromsmoreira/turbo-todo/api/models"
)

func FromModelToDto(model *models.Todo) (*dto.TodoData, error) {
	return nil, errors.New("mapping error from model to dto")
}

func FromDtoToModel(dto *dto.TodoData) (*models.Todo, error) {
	return nil, errors.New("mapping error from dto to model")
}

func FromModelToData(model *models.Todo) (*data.Todo, error) {
	return nil, errors.New("mapping error from model to data")
}

func FromDataToModel(data *data.Todo) (*models.Todo, error) {
	return nil, errors.New("mapping error from data to model")
}
