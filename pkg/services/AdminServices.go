package services

import (
	"fmt"
	"time"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type AdminServices struct {
	admin *models.Admin
}

func (c *AdminServices) GenerateToken(ad *models.Admin) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &models.Claims{
		Username: ad.Name,
		Email:    ad.Email,
		Phone:    ad.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(models.Jwtkey)
	if err != nil {
		return "", err
	}
	return TokenString, nil
}
func (c *AdminServices) LoginService(ad *models.Admin) error {
	admimDetails, err := c.GetDb(ad.Name)
	if err != nil {
		return err
	}
	if admimDetails.Password != ad.Password {
		return fmt.Errorf("password mismatch")
	}
	return nil
}
func (c *AdminServices) GetDb(cname string) (*models.Admin, error) {
	db := models.Dbconn()

	var admimDetails models.Admin

	query := `SELECT * FROM admins WHERE name=$1`
	err := db.QueryRow(query, cname).Scan(&admimDetails.Id, &admimDetails.Name, &admimDetails.Email, &admimDetails.Password, &admimDetails.Phone)

	if err != nil {
		return nil, err
	}
	return &admimDetails, nil
}

func (c *AdminServices) GetAllComplaintService() ([]*models.Complaint, error) {
	db := models.Dbconn()
	selDb, err := db.Query("SELECT * FROM complaints")
	if err != nil {
		return nil, err
	}

	allComplaints := &ComplaintService{}
	for selDb.Next() {
		var Complaint models.Complaint
		selDb.Scan(&Complaint.Customer.Id, &Complaint.Customer.Name, &Complaint.Customer.Email, &Complaint.Customer.Phone, &Complaint.Complaint)
		allComplaints.complaints = append(allComplaints.complaints, &Complaint)
	}
	return allComplaints.complaints, nil
}
