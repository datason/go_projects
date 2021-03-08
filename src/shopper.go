package main

import (
	"fmt"
	"magazine"
	"salut"
)

func main() {
	salut.Hello()
	// type subscriber = magazine.Subscriber
	var s magazine.Subscriber
	s.Name = "Adam Smith"
	s.Rate = 5.99
	s.Active = true
	fmt.Println(s.Name)

	c := magazine.Subscriber{Rate: 123, Name: "Ada Rogovtseva"}
	fmt.Println(c.Name)

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"

	employee.HomeAdress = address
	fmt.Println(employee)
}
