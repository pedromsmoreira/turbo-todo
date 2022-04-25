package api

import (
	"net/http"
	"testing"
)

func TestHealthschecks(t *testing.T) {
	t.Run("status endpoint returns up", func(t *testing.T) {
		given, when, then := newApiStage(t, "localhost:5000")

		given.
			anHttpClientIsCreated()

		when.
			statusEndpointIsQueried()

		then.
			shouldReturnStatusCode(http.StatusOK).
			and().
			returnsStatus("up")
	})

	t.Run("ping endpoint returns pong", func(t *testing.T) {
		given, when, then := newApiStage(t, "localhost:5000")

		given.
			anHttpClientIsCreated()

		when.
			pingEndpointIsQueried()

		then.
			shouldReturnStatusCode(http.StatusOK).
			and().
			returnsPong()
	})

}
