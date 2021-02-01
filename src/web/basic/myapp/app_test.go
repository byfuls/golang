package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "hello world", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "hello bar", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=byfuls", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "hello byfuls", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name": "first", "last_name": "last", "email": "email"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(t, err)
	assert.Equal(t, "first", user.FirstName)
	assert.Equal(t, "last", user.LastName)
}
