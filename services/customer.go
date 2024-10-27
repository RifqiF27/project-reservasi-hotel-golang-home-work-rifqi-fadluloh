package services

import (
	"main/models"
)



func ViewAllCustomer() []models.User {
	var customer []models.User
	for _, c := range users {
		if c.Role == "customer"{
			customer = append(customer, c)
		}	
	}
	return customer
}