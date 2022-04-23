package api

import (
	"net/http"
	"testing"
)

func TestServerHealthcheckIsWorking(t *testing.T) {
	given, when, then := newApiStage()

	given.
		anHttpClientIsCreated()

	when.
		healthEndpointIsQueried()

	then.
		shouldReturnStatusCode(http.StatusOK).
		and().
		returnsStatus("up")
}
