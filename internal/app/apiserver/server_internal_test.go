package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pitshifer/valera-acceptor/internal/app/store/teststore"
	"gotest.tools/assert"
)

func TestServer_HandleDeviceCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/devices", nil)
	s := NewServer(teststore.New())
	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
