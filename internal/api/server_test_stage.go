package api

import (
	"net/http"

	"github.com/stretchr/testify/suite"
)

type apiStage struct {
	suite.Suite
	http http.Client
}

func (as *apiStage) and() *apiStage {
	return as
}

func newApiStage() (*apiStage, *apiStage, *apiStage) {
	as := &apiStage{
		http: *http.DefaultClient,
	}

	return as, as, as
}

func (as *apiStage) anHttpClientIsCreated() *apiStage {
	return as
}

func (as *apiStage) healthEndpointIsQueried() *apiStage {
	return as
}

func (as *apiStage) shouldReturnStatusCode(statusCode int) *apiStage {
	return as
}

func (as *apiStage) returnsStatus(expStatus string) *apiStage {
	return as
}
