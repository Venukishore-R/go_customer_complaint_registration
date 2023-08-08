package services

import (
	"testing"

	"github.com/Venukishore-R/admin_customer_complaint_project/pkg/models"
)

func TestCustomerService_CustomerRegisterService(t *testing.T) {
	type fields struct {
		customer *models.CustomerRepository
	}
	type args struct {
		cus *models.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				customer: &models.CustomerRepository{},
			},

			args: args{
				cus: &models.Customer{
					Name:     "Vk",
					Email:    "vk@gmail.com",
					Phone:    "62132713",
					Password: "12345678",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomerService{
				customer: tt.fields.customer,
			}
			if err := c.CustomerRegisterService(tt.args.cus); (err != nil) != tt.wantErr {
				t.Errorf("CustomerService.CustomerRegisterService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
