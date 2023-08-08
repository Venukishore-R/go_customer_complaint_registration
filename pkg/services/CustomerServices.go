package services

import (
	"time"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	customer *models.CustomerRepository
}

func Hashing(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func DeHashing(hashed_password, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err
}

func (c *CustomerService) CustomerRegisterService(cus *models.Customer) error {
	db := models.Dbconn()
	pass, _ := Hashing(cus.Password)
	sqlStatement := `
	INSERT INTO customers (name, email, password, phone)
	VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, cus.Name, cus.Email, pass, cus.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) CusotmerLoginService(cus *models.Customer) error {
	customerDetails, err := c.GetDb(cus.Name)
	if err != nil {
		return err
	}
	err = DeHashing(customerDetails.Password, cus.Password)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) GetDb(cname string) (*models.Customer, error) {
	db := models.Dbconn()

	var customerDetails models.Customer

	query := `SELECT * FROM customers WHERE name=$1`
	err := db.QueryRow(query, cname).Scan(&customerDetails.Id, &customerDetails.Name, &customerDetails.Email, &customerDetails.Password, &customerDetails.Phone)

	if err != nil {
		return nil, err
	}
	return &customerDetails, nil
}

func (c *CustomerService) GenerateToken(cus *models.Customer) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &models.Claims{
		Username: cus.Name,
		Email:    cus.Email,
		Phone:    cus.Phone,
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
