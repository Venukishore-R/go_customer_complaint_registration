package mappers

import (
	"encoding/json"
	"net/http"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/models"
)

func Decode(w http.ResponseWriter, r *http.Request) *models.Customer {
	var Customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&Customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cannot able to parse"))
		return nil
	}
	return &Customer
}

func Decoder(w http.ResponseWriter, r *http.Request, c []*models.Customer) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Cannot able to parse"))
		return
	}

	w.Header().Add("content-type", "Application/json")
	w.Write(jsonData)
	return
}

// Complaint controller decoder
func ComplaintDecode(w http.ResponseWriter, r *http.Request) string {
	var Complaint models.Complaint
	err := json.NewDecoder(r.Body).Decode(&Complaint)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cannot able to parse"))
		return ""
	}
	return Complaint.Complaint
}

func EncodeComplaint(w http.ResponseWriter, r *http.Request, c []*models.Complaint) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Cannot able to parse"))
		return
	}

	w.Header().Add("content-type", "Application/json")
	w.Write(jsonData)
	return
}

// Admin mappers
func AdminDecode(w http.ResponseWriter, r *http.Request) *models.Admin {
	var Admin models.Admin
	err := json.NewDecoder(r.Body).Decode(&Admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cannot able to parse"))
		return nil
	}
	return &Admin
}

// func AdminComplaintsEncoder(w http.ResponseWriter, r *http.Request, []*models.Complaint) {

// }
// func EncodeComplaint(w http.ResponseWriter, r *http.Request, c []*models.Complaint) {
// 	jsonData, err := json.Marshal(c)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadGateway)
// 		w.Write([]byte("Cannot able to parse"))
// 		return
// 	}

// 	w.Header().Add("content-type", "Application/json")
// 	w.Write(jsonData)
// 	return
// }
