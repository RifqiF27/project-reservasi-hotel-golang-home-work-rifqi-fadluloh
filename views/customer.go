package views

import (
	"fmt"
	"main/services"
	"strings"
)

func ViewAllCustomer() {
	customer := services.ViewAllCustomer()
	fmt.Println("All Customer")
	fmt.Println(strings.Repeat("-", 40))

	for _, c := range customer {
		fmt.Printf("ID: %d, Username: %s, Phone Number: %s\n",
			c.ID, c.Username, c.Username)
		fmt.Println(strings.Repeat("-", 40))

	}
}
