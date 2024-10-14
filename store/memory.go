package store

import (
	"strings"

	"github.com/pmarley/htmx-starter/types"
)

type InMemoryStore struct {
	lastId int
	todos  []types.Todo
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

func (ms *InMemoryStore) AllTodos() []types.Todo {
	return ms.todos
}

func (ms *InMemoryStore) CreateTodo(description string) types.Todo {
	ms.lastId++
	todo := types.Todo{Id: ms.lastId, Description: description, Done: false}
	ms.todos = append(ms.todos, todo)
	return todo
}

func (ms *InMemoryStore) indexOf(id int) int {
	for idx, todo := range ms.todos {
		if todo.Id == id {
			return idx
		}
	}
	return -1
}

func (ms *InMemoryStore) ToggleTodo(id int) types.Todo {
	idx := ms.indexOf(id)
	if idx == -1 {
		return types.Todo{}
	}

	ms.todos[idx].Done = !ms.todos[idx].Done
	return ms.todos[idx]
}

func (ms *InMemoryStore) DeleteTodo(id int) {
	idx := ms.indexOf(id)
	ms.todos = append(ms.todos[:idx], ms.todos[idx+1:]...)
}

func (ms *InMemoryStore) Filter(expr string) (res []types.Todo) {
	for _, t := range ms.todos {
		if strings.Contains(strings.ToLower(t.Description), strings.ToLower(expr)) || strings.Contains(strings.ToUpper(t.Description), strings.ToUpper(expr)) {
			res = append(res, t)
		}
	}
	return
}
