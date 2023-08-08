package models

type Complaint struct {
	Customer  Customer
	Complaint string `json:"complaint"`
}
