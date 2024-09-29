package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	contextusage "github.com/mbilaljawwad/golang-channels-learning/internal/context_usage"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"
	store := &SpyStore{response: data}
	svr := contextusage.Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	// // Adding canceling context
	// cancellingCtx, cancel := context.WithCancel(request.Context())
	// time.AfterFunc(5*time.Millisecond, cancel)
	// request = request.WithContext(cancellingCtx)

	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if !store.cancelled {
		t.Error("Store was not told to cancel")
	}

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
