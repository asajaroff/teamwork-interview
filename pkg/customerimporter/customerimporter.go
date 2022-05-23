package customerimporter

import "fmt"

type Customer struct {
	fistName  string
	lastName  string
	email     string
	gender    string
	ipAddress string
}

func GetCustomerName(c *Customer) {
	fmt.Println(c.fistName)
}
