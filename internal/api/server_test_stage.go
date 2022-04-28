package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/pedromsmoreira/turbo-todo/internal/api/todos"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"

	"github.com/jackc/pgx/v4"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Status struct {
	Status string `json:"status"`
}

type Ping struct {
	Message string `json:"message"`
}

type apiStage struct {
	suite.Suite
	t        *testing.T
	host     string
	http     *http.Client
	response *http.Response
	dbConn   *pgx.Conn
}

func (as *apiStage) and() *apiStage {
	return as
}

func newApiStage(t *testing.T, url string) (*apiStage, *apiStage, *apiStage) {
	dbConn := getDbConn(t)
	as := &apiStage{
		t:      t,
		host:   fmt.Sprintf("http://%s", url),
		dbConn: dbConn,
	}

	return as, as, as
}

func getDbConn(t *testing.T) *pgx.Conn {
	config, err := pgx.ParseConfig("postgresql://root@localhost:26257/defaultdb?sslmode=disable")
	config.Database = "turbotodo"
	require.Nil(t, err)
	conn, err := pgx.ConnectConfig(context.Background(), config)
	require.Nil(t, err)
	return conn
}

func (as *apiStage) anHttpClientIsCreated() *apiStage {
	as.http = http.DefaultClient
	return as
}

func (as *apiStage) statusEndpointIsQueried() *apiStage {

	response, err := as.http.Get(fmt.Sprintf("%s/status", as.host))
	require.Nil(as.t, err)
	require.NotNil(as.t, response)
	as.response = response

	return as
}

func (as *apiStage) pingEndpointIsQueried() *apiStage {

	response, err := as.http.Get(fmt.Sprintf("%s/ping", as.host))
	require.Nil(as.t, err)
	require.NotNil(as.t, response)
	as.response = response

	return as
}

func (as *apiStage) shouldReturnStatusCode(statusCode int) *apiStage {
	require.Equal(as.t, statusCode, as.response.StatusCode)
	return as
}

func (as *apiStage) returnsPong() *apiStage {
	body, err := ioutil.ReadAll(as.response.Body)
	require.Nil(as.t, err)

	st := new(Ping)
	err = json.Unmarshal(body, st)
	require.Nil(as.t, err)
	require.Equal(as.t, "pong", st.Message)
	return as
}

func (as *apiStage) returnsStatus(expStatus string) *apiStage {
	body, err := ioutil.ReadAll(as.response.Body)
	require.Nil(as.t, err)

	st := new(Status)
	err = json.Unmarshal(body, st)
	require.Nil(as.t, err)
	require.Equal(as.t, expStatus, st.Status)
	return as
}

func (as *apiStage) todoIsCreatedForId(id string) *apiStage {
	td := todos.CreateTodoRequest{
		Data: &todos.Dto{
			DateCreated: time.Now().UTC().Format(time.RFC3339),
			Version:     0,
			Attributes: &todos.Attributes{
				Title:   "New todo",
				Status:  "done",
				Content: "todo content",
				Tags:    []string{"test", "new"},
			},
		},
	}

	err := crdbpgx.ExecuteTx(context.Background(), as.dbConn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(),
			"INSERT INTO todos (id, title, datecreated, content, tags, status, version) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			id, td.Data.Attributes.Title, td.Data.DateCreated, td.Data.Attributes.Content, td.Data.Attributes.Tags, td.Data.Attributes.Status, td.Data.Version)
		require.Nil(as.t, err)
		return nil
	})

	require.Nil(as.t, err)
	return as
}

func (as *apiStage) listEndpointIsQueriedForId(id string) *apiStage {
	return as
}

func (as *apiStage) shouldListWithOneItem() *apiStage {
	return as
}
