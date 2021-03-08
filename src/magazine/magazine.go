package magazine

type Subscriber struct {
	Name       string
	Rate       float64
	Active     bool
	HomeAdress Address
}

type Employee struct {
	Name       string
	Salary     float64
	HomeAdress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
