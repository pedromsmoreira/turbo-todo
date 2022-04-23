package todo

import (
	"errors"
)

func FromModelToDto(model *todo) (*TodoDto, error) {
	return nil, errors.New("mapping error from model to dto")
}

func FromDtoToModel(dto *TodoDto) (*todo, error) {
	return nil, errors.New("mapping error from dto to model")
}

func FromModelToData(model *todo) (*todoData, error) {
	return nil, errors.New("mapping error from model to data")
}

func FromDataToModel(data *todoData) (*todo, error) {
	return nil, errors.New("mapping error from data to model")
}
