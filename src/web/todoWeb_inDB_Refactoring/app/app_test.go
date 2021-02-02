package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
	"web/todoWeb_inDB_Refactoring/model"

	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	os.Remove("./test.db")
	ah := MakeHandler()
	defer ah.Close()

	ts := httptest.NewServer(ah)
	defer ts.Close()

	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo"}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	var todo model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(t, err)
	assert.Equal(t, todo.Name, "Test todo")
	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo2"}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(t, err)
	assert.Equal(t, todo.Name, "Test todo2")
	id2 := todo.ID

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	todos := []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(t, err)
	assert.Equal(t, len(todos), 2)
	for _, v := range todos {
		if v.ID == id1 {
			assert.Equal(t, "Test todo", v.Name)
		} else if v.ID == id2 {
			assert.Equal(t, "Test todo2", v.Name)
		} else {
			assert.Error(t, fmt.Errorf("testID should be id1 or id2"))
		}
	}

	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id1) + "?complete=true")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	todos = []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(t, err)
	assert.Equal(t, len(todos), 2)
	for _, v := range todos {
		if v.ID == id1 {
			assert.True(t, v.Completed)
		}
	}

	req, _ := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	todos = []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(t, err)
	assert.Equal(t, len(todos), 1)
	for _, v := range todos {
		assert.Equal(t, v.ID, id2)
	}
}
