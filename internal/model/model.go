package internal


type Product struct{
	Quantity int
	Id string
	Price int
	Name string
	Ordered int
	PurchHistory *[]*Trans
	OrderHistory *[]*Trans
}
type Trans struct{
	Price int
	Quantity int
}