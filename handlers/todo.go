package handlers

import (
	"net/http"
	"strconv"

	"github.com/pmarley/htmx-starter/types"
	"github.com/pmarley/htmx-starter/views"
)

type TodoStore interface {
	AllTodos() []types.Todo
	CreateTodo(description string) types.Todo
	ToggleTodo(id int) types.Todo
	DeleteTodo(id int)
	Filter(expr string) []types.Todo
}

type TodoHandler struct {
	store TodoStore
}

func NewTodoHandler(store TodoStore) *TodoHandler {
	return &TodoHandler{store: store}
}

func (th *TodoHandler) Home(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, views.TodoPage(th.store.AllTodos()))
}

func (th *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) error {
	descripton := r.FormValue("description")

	if descripton == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return nil
	}

	todo := th.store.CreateTodo(descripton)
	return render(w, r, views.Todo(todo))
}

func (th *TodoHandler) ToggleTodo(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(r.PathValue("id"))
	todo := th.store.ToggleTodo(id)
	return render(w, r, views.Todo(todo))
}

func (th *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(r.PathValue("id"))
	th.store.DeleteTodo(id)
	w.Write([]byte(""))
	return nil

}

func (th *TodoHandler) FilterTodos(w http.ResponseWriter, r *http.Request) error {
	filter := r.URL.Query().Get("filter")

	return render(w, r, views.TodoList(th.store.Filter(filter)))
}
