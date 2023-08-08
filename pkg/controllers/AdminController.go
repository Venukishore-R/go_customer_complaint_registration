package controllers

import (
	"net/http"
	"time"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/mappers"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/services"
)

type AdminController struct {
	admin *services.AdminServices
}

func NewAdminController(c *services.AdminServices) *AdminController {
	return &AdminController{
		admin: c,
	}
}

func (c *AdminController) Login(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,
		&http.Cookie{
			Name:    "admin-token",
			Value:   "",
			Expires: time.Now().Add(time.Hour * 24),
		},
	)
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Method mismatch"))
		return
	}
	admin := mappers.AdminDecode(w, r)
	if admin == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot able to parse"))
		return
	}
	err := c.admin.LoginService(admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("UNauthorised user or Register first"))
		return
	}
	token, err := c.admin.GenerateToken(admin)
	if err != nil || token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Cannot able to generate token"))
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "admin-token",
			Value:   token,
			Expires: time.Now().Add(time.Hour * 24),
		},
	)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Logined successfully"))
	return
}

func (c *AdminController) GetAllComplaint(w http.ResponseWriter, r *http.Request) {
	complaints, err := c.admin.GetAllComplaintService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	mappers.EncodeComplaint(w, r, complaints)
}

// func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	todoId, err := strconv.Atoi(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Something wrong, Check the request body"))
// 		return
// 	}
// 	var todo models.TodoList
// 	if r.Method == "PUT" {
// 		err := json.NewDecoder(r.Body).Decode(&todo)
// 		if err != nil {
// 			http.Error(w, "Couldn't parse", http.StatusBadRequest)
// 			return
// 		}
// 	}
// 	err = c.todo.UpdateTodoModels(todo, todoId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte("Something wrong, Check the request body"))
// 		return
// 	}
// 	mappers.Decoder(w, r, []*models.TodoList{&todo})
// }

// func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	todoId, err := strconv.Atoi(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Something wrong, Check the request body"))
// 		return
// 	}
// 	err = c.todo.DeleteTodo(todoId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Cannot find task"))
// 		return
// 	}
// 	w.Write([]byte("Deleted successfully"))
// }
