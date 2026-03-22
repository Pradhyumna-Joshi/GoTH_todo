package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pradhyumna-Joshi/go_todo_htmx/components"
	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/common"
	"github.com/Pradhyumna-Joshi/go_todo_htmx/internal/todo/service"
	"github.com/Pradhyumna-Joshi/go_todo_htmx/utils"
)

type TodoHandler struct {
	service service.Service
}

func NewTodoHandler(service service.Service) TodoHandler {
	return TodoHandler{service}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing request body"))
		return
	}

	var payload TodoPayload

	payload.Title = r.FormValue("title")
	payload.Description = r.FormValue("description")

	_, err := h.service.CreateTodo(r.Context(), toTodo(payload))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	resp, err := h.service.GetTodos(r.Context(), common.Filter{})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	components.TodoList(resp).Render(r.Context(), w)

}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	filter := common.Filter{
		Sort:      params.Get("sort"),
		Completed: params.Get("completed"),
	}

	resp, err := h.service.GetTodos(r.Context(), filter)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	cmp := components.TodoList(resp)
	cmp.Render(r.Context(), w)
}

func (h *TodoHandler) ToggleTodo(w http.ResponseWriter, r *http.Request) {

	_id := r.PathValue("id")
	id, _ := strconv.Atoi(_id)

	updatedTodo, err := h.service.ToggleTodo(r.Context(), id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	components.TodoItem(updatedTodo).Render(r.Context(), w)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing request body"))
		return
	}

	_id := r.PathValue("id")
	id, _ := strconv.Atoi(_id)

	var payload TodoPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.service.UpdateTodo(r.Context(), id, toTodo(payload))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	resp, err := h.service.GetTodos(r.Context(), common.Filter{})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	components.TodoList(resp).Render(r.Context(), w)

}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	_id := r.PathValue("id")
	id, _ := strconv.Atoi(_id)

	err := h.service.DeleteTodo(r.Context(), id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

}

//dto

func toTodo(p TodoPayload) common.Todo {

	return common.Todo{
		Title:       p.Title,
		Description: p.Description,
		IsComplete:  p.IsComplete,
	}
}
