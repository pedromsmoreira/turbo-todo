package todos

import (
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"

	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"

	"github.com/jackc/pgx/v4"
)

type TodoRepository interface {
	Get(id string) (*todo, error)
	Add(todo *todo) (*todo, error)
	CloseDB() error
}

type CockroachDbTodoRepository struct {
	conn *pgx.Conn
}

func NewCockroachDbTodoRepository(cfg *configs.Config) (TodoRepository, error) {
	config, err := pgx.ParseConfig(fmt.Sprintf("postgresql://root@%s/defaultdb?sslmode=disable", cfg.Database.Host))
	config.Database = cfg.Database.DbName
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error schema configuration: %v", err))
	}
	conn, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error connecting to the database: %v", err))
	}

	return &CockroachDbTodoRepository{
		conn: conn,
	}, nil
}

func (imr *CockroachDbTodoRepository) CloseDB() error {
	return imr.conn.Close(context.Background())
}

func (imr *CockroachDbTodoRepository) Get(id string) (*todo, error) {
	td := new(todo)
	err := pgxscan.Select(context.Background(),
		imr.conn,
		&td,
		fmt.Sprintf(`SELECT id, title, content, tags, status, version, datecreated FROM todos WHERE id='%s'`, id))

	if err != nil {
		return nil, NewDbResourceNotFound(fmt.Sprintf("todo with id: %s not found", id))
	}

	return td, nil
}

func (imr *CockroachDbTodoRepository) Add(todo *todo) (*todo, error) {
	return nil, errors.New("not implemented")
}
