package internal


// product Struct and fn's to completely encapsulate this struct

type Product struct{
	quantity int
	id string
	price int
	name string
	ordered int
	purchHistory *[]*Trans
	orderHistory *[]*Trans
	
}
type Trans struct{
	price int
	quantity int
}
func (t Trans)GetTransPrice()int{
	return t.price
}
func (t Trans)GetTransQuantity()int{
	return t.quantity
}
func CreateProdInstance(id string,price int,name string)*Product{
	prod := Product{
		id:id,
		name:name,
		price:price,
		quantity:0,
		ordered: 0,
		purchHistory: &[]*Trans{},
		orderHistory: &[]*Trans{},
	}
	return &prod
}
func (p *Product)GetId()string{
	return p.id
}
func (p *Product)GetName()string{
	return p.name
}

func (p *Product)GetPrice()int{
	return p.price
}
func (p *Product)ChangePrice(price int,name string){
	p.price = price
	p.name = name
}
func (p *Product)GetQuantity()int{
	return p.quantity
}
func (p *Product)GetNumberOfOrdered()int{
	return p.ordered
}
func (p *Product)AddNumberOfOrdered(k int){
	p.ordered += k
}
func (p *Product)ChangeQuantity(k int){
	p.quantity=k
}

func (p *Product)AppendOrderHistory(quantity int){
	*p.orderHistory =append(*p.orderHistory,&Trans{quantity: quantity, price: p.GetPrice()})
	
}
func (p *Product)AppendPurchaseHistory(price ,quantity int){
	*p.purchHistory =append(*p.purchHistory,&Trans{quantity: quantity, price: price})

}
func (p* Product)AverageOrderedPrice()float64{
	cnt := 0
	sum := 0
	for _, v := range *p.orderHistory {
		cnt += v.quantity
		sum += v.price * v.quantity
	}
	var avg float64 = float64(sum)/float64(cnt)
	return avg
}
func (p* Product)AveragePurchasePrice()float64{
	cnt := 0
	sum := 0
	for _, v := range *p.purchHistory {
		cnt += v.quantity
		sum += v.price * v.quantity
	}
	var avg float64 = float64(sum)/float64(cnt)
	return avg
}
func (p* Product)GetOrdersHistory()[]*Trans{
	return *p.orderHistory
}

