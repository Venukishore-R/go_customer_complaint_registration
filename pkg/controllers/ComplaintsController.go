package controllers

import (
	"net/http"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/mappers"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/models"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/services"
)

type ComplaintController struct {
	complaint *services.ComplaintService
}

func NewComplaintController(c *services.ComplaintService) *ComplaintController {
	return &ComplaintController{
		complaint: c,
	}
}

func (c *ComplaintController) RegisterComplaint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong method"))
		return
	}

	Complaint := mappers.ComplaintDecode(w, r)
	if Complaint == "" {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Cannot able to parse"))
		return
	}

	err := c.complaint.RegisterComplaintService(w, r, Complaint)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Cannot able to register your complaint"))
		return
	}

	mappers.EncodeComplaint(w, r, []*models.Complaint{&models.Complaint{Complaint: Complaint}})
}
