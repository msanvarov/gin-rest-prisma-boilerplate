package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/msanvarov/gin-rest-prisma-boilerplate/forms"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/prisma-client"
	"github.com/stretchr/testify/assert"
)

/**
* Test Registration
* Must return response code 200
 */
func TestRegistration(t *testing.T) {
	data, _ := json.Marshal(registrationPayload)
	if req, err := http.NewRequest(
		"POST", "/api/v1/register", bytes.NewBufferString(string(data))); err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		routing.ServeHTTP(resp, req)
		fmt.Println(resp.Body)
		assert.Equal(t, resp.Code, http.StatusOK)
	}
}

/**
* Test Registration with duplicate username
* Must return response code 406
 */
func TestDuplicateUsername(t *testing.T) {
	data, _ := json.Marshal(registrationPayload)
	if req, err := http.NewRequest(
		"POST", "/api/v1/register", bytes.NewBufferString(string(data))); err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		routing.ServeHTTP(resp, req)
		fmt.Println(resp.Body)
		assert.Equal(t, resp.Code, http.StatusNotAcceptable)
	}
}

/**
* Test Registration with duplicate username
* Must return response code 406
 */
func TestInvalidPayload(t *testing.T) {
	data, _ := json.Marshal(forms.RegistrationForm{
		Name:     "not-valid-payload",
		Username: "no-email-or-password-fields",
	})
	if req, err := http.NewRequest(
		"POST", "/api/v1/register", bytes.NewBufferString(string(data))); err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		routing.ServeHTTP(resp, req)
		fmt.Println(resp.Body)
		assert.Equal(t, resp.Code, http.StatusBadRequest)
	}
}

/**
* Test Login
* Test user login and store the cookie on local variable [cookie]
* Must return response code 200
 */
func TestLogin(t *testing.T) {
	data, _ := json.Marshal(loginPayload)
	if req, err := http.NewRequest(
		"POST", "/api/v1/login", bytes.NewBufferString(string(data))); err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		routing.ServeHTTP(resp, req)
		fmt.Println(resp.Body)
		cookie = resp.Header().Get("Set-Cookie")
		assert.Equal(t, resp.Code, http.StatusOK)
	}
}

/**
* Test fetching session data
* Given the cookie, the web app will query redis for session data to return
 */
func TestFetchingSessionData(t *testing.T) {
	if req, err := http.NewRequest("GET", "/api/v1/session", nil); err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		req.Header.Set("Cookie", cookie)
		resp := httptest.NewRecorder()
		fmt.Println(resp.Body)
		routing.ServeHTTP(resp, req)
		assert.Equal(t, resp.Code, 200)
	}
}

/**
* Test logging out a user
* Must return response code 200
 */
func TestLogout(t *testing.T) {
	if req, err := http.NewRequest("POST", "/api/v1/logout", nil); err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		req.Header.Set("Cookie", cookie)
		resp := httptest.NewRecorder()
		fmt.Println(resp.Body)
		routing.ServeHTTP(resp, req)
		assert.Equal(t, resp.Code, 200)
	}
}

/**
* Test teardown of test account
* Must return a not nill operation status
 */
func TestTeardown(t *testing.T) {
	ctx := context.Background()
	deletedUser, err := client.DeleteUser(prisma.UserWhereUniqueInput{
		Username: &registrationPayload.Username,
	}).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		assert.Fail(t, err.Error())
		return
	} else {
		assert.NotEqual(t, nil, deletedUser)
	}
}
