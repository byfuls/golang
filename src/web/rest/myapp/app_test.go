package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "hello world", string(data))
}

func TestUsers(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(data), "No Users")
}

func TestGetUserInfo(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(data), "No User Id:89")
}

func TestCreateUserInfo(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "appplication/json",
		strings.NewReader(`{"first_name": "first", "last_name": "last", "email": "email@gmail.com"}`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, user.ID)

	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, user2.ID)
	assert.Equal(t, user.FirstName, user2.FirstName)
}

func TestDeleteUserInfo(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(data), "No User Id:1")

	resp, err = http.Post(ts.URL+"/users", "appplication/json",
		strings.NewReader(`{"first_name": "first", "last_name": "last", "email": "email@gmail.com"}`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, user.ID)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(data), "Deleted User Id:1")
}

func TestUpdateUser(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id": 1, "first_name": "updated", "last_name": "updated", "email": "updated@gamil.com"}`))
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(data), "No User Id:1")

	resp, err = http.Post(ts.URL+"/users", "appplication/json",
		strings.NewReader(`{"first_name": "first", "last_name": "last", "email": "email@gmail.com"}`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, user.ID)

	updateStr := fmt.Sprintf(`{"id": %d, "first_name": "updated"}`, user.ID)
	req, _ = http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(updateStr))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(t, err)
	assert.Equal(t, updateUser.ID, user.ID)
	assert.Equal(t, "updated", updateUser.FirstName)
	assert.Equal(t, user.LastName, updateUser.LastName)
	assert.Equal(t, user.Email, updateUser.Email)
}

func TestUsers_WithUsersData(t *testing.T) {
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "appplication/json",
		strings.NewReader(`{"first_name": "first", "last_name": "last", "email": "email@gmail.com"}`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	resp, err = http.Post(ts.URL+"/users", "appplication/json",
		strings.NewReader(`{"first_name": "json", "last_name": "park", "email": "jp@gmail.com"}`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/users")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	users := []*User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
}
