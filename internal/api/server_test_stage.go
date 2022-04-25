package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

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
}

func (as *apiStage) and() *apiStage {
	return as
}

func newApiStage(t *testing.T, url string) (*apiStage, *apiStage, *apiStage) {
	as := &apiStage{
		t:    t,
		host: fmt.Sprintf("http://%s", url),
	}

	return as, as, as
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
	return as
}

func (as *apiStage) listEndpointIsQueriedForId(id string) *apiStage {
	return as
}

func (as *apiStage) shouldListWithOneItem() *apiStage {
	return as
}
