package api

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestHealthsChecks(t *testing.T) {
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

func TestTodoEndpoints(t *testing.T) {
	t.Run("list returns array with one todos item", func(t *testing.T) {
		given, when, then := newApiStage(t, "localhost:5000")
		id := uuid.New().String()

		given.
			anHttpClientIsCreated().
			and().
			todoIsCreatedForId(id)

		when.
			listEndpointIsQueriedForId(id)

		then.
			shouldReturnStatusCode(http.StatusOK).
			and().
			shouldListWithOneItem()
	})
}
