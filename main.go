package main

import (
	"net/http"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/controllers"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/routes"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/services"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	CustomerRepository := services.CustomerService{}
	Customers := controllers.NewCustomerController(&CustomerRepository)

	ComplaintRepository := services.ComplaintService{}
	Complaints := controllers.NewComplaintController(&ComplaintRepository)

	routes.Router(r, Customers, Complaints)

	AdminRepository := services.AdminServices{}
	Admins := controllers.NewAdminController(&AdminRepository)
	routes.AdminRouter(r, Admins)
	http.ListenAndServe(":8000", r)

}
