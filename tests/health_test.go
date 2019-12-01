package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	data, _ := json.Marshal(loginPayload)
	if req, err := http.NewRequest(
		"GET", "/ping", bytes.NewBufferString(string(data))); err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		req.Header.Set("Content-Type", "text/plain")
		resp := httptest.NewRecorder()
		routing.ServeHTTP(resp, req)
		fmt.Println(resp.Body)
		assert.Equal(t, resp.Code, http.StatusOK)
		assert.Equal(t, resp.Body.String(), "Pong")
	}
}
