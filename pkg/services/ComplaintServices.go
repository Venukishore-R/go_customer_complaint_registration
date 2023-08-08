package services

import (
	"fmt"
	"net/http"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type ComplaintService struct {
	complaints []*models.Complaint
}

func (c *ComplaintService) RegisterComplaintService(w http.ResponseWriter, r *http.Request, com string) error {
	db := models.Dbconn()
	cookie, err := r.Cookie("token")
	if cookie == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return err
	}
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}

		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	cus, err := c.ExtractUnverifiedClaims(cookie.Value)
	if err != nil {
		return err
	}
	sqlStatement := `INSERT INTO complaints (name, email, phone, complaint) VALUES ($1, $2, $3, $4)`
	db.Exec(sqlStatement, cus.Name, cus.Email, cus.Phone, com)
	return nil
}

func (c *ComplaintService) ExtractUnverifiedClaims(tokenString string) (*models.Customer, error) {
	var name, email, phone string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		name = fmt.Sprint(claims["Username"])
		email = fmt.Sprint(claims["Email"])
		phone = fmt.Sprint(claims["Phone"])
	}
	if name == "" || email == "" || phone == "" {
		return nil, fmt.Errorf("invalid token payload")
	}
	customer := &models.Customer{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	return customer, nil
}

//https://stackoverflow.com/questions/39859244/how-to-extract-the-claims-from-jwt-token
