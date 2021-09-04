package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RodrigoCMoraes/judges_submeter/internal/submeter"
)

func TestServer(t *testing.T) {
	t.Run("returns mocked response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/submit", nil)
		response := httptest.NewRecorder()

		submeter.Submit(response, request)

		got := response.Body.String()
		want := "mocked submission response."

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
