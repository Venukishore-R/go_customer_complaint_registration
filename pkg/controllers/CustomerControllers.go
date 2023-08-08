package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/mappers"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/models"
	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/services"
)

type Customer struct {
	customer *services.CustomerService
}

func NewCustomerController(c *services.CustomerService) *Customer {
	return &Customer{
		customer: c,
	}
}

func (c *Customer) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong method"))
		return
	}

	Customer := mappers.Decode(w, r)
	if Customer == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cannot able to parse"))

		return
	}

	result := c.customer.CustomerRegisterService(Customer)
	if result != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot able to register a user"))
		log.Println(result)
		return
	}
	mappers.Decoder(w, r, []*models.Customer{Customer})
}

func (c *Customer) Login(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Now().Add(time.Hour * 24),
		},
	)
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Method mismatch"))
		return
	}
	customer := mappers.Decode(w, r)
	if customer == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot able to parse"))
		return
	}

	err := c.customer.CusotmerLoginService(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("UNauthorised user or Register first"))
		fmt.Println(err)
		return
	}

	token, err := c.customer.GenerateToken(customer)
	if err != nil || token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Cannot able to generate token"))
		fmt.Println(err)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(time.Hour * 24),
		},
	)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Logined successfully"))
	return

}

func (c *Customer) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
