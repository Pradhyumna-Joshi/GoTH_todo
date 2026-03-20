package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Something went wrong: %s ", err.Error()))
		return
	}

	err := h.service.CreateTodo(r.Context(), toTodo(payload))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"status": "success",
	})
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
	utils.WriteJSON(w, http.StatusOK, resp)
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

	err := h.service.UpdateTodo(r.Context(), id, toTodo(payload))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"status": "success",
	})

}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	_id := r.PathValue("id")
	id, _ := strconv.Atoi(_id)

	err := h.service.DeleteTodo(r.Context(), id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"status": "success",
	})

}

//dto

func toTodo(p TodoPayload) service.Todo {

	return service.Todo{
		Title:       p.Title,
		Description: p.Description,
		IsComplete:  p.IsComplete,
	}
}
