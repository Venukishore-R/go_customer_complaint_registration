package routes

import (
	"net/http"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/controllers"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/services"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router, c *controllers.Customer, u *controllers.ComplaintController) {
	r.HandleFunc("/register", c.Register).Methods("POST")
	r.HandleFunc("/login", c.Login).Methods("POST")
	r.Handle("/home", services.ValidateTokenMiddleware(http.HandlerFunc(c.Home))).Methods("GET")
	r.Handle("/complaint", services.ValidateTokenMiddleware(http.HandlerFunc(u.RegisterComplaint))).Methods("POST")
	r.HandleFunc("/", c.Home).Methods("GET")
}

func AdminRouter(r *mux.Router, c *controllers.AdminController) {
	r.HandleFunc("/admin/login", c.Login).Methods("POST")
	// r.Handle("/admin/allcomplaints", services.AdminValidateTokenMiddleware(http.HandlerFunc(c.GetAllComplaint))).Methods("GET")
	r.HandleFunc("/admin/allcomplaints", c.GetAllComplaint).Methods("GET")

}
