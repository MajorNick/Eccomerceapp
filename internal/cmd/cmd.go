package cmd

import (
	cache "Europe/internal/cache"
	_ "Europe/internal/model"
	"fmt"
)

func save_product(id string, name string, price int) {
	product, err := cache.GetProductById(id)
	if err != nil {
		if err == cache.ErrorNotExists {
			cache.AddInDataStore(id, name, price)
			fmt.Println("Product Succesfully added!")
			fmt.Println()
		} else {
			fmt.Println(err)
		}
		return
	}
	product.ChangePrice(price,name)
	fmt.Println("Product Price Succesfully Changed!")
	fmt.Println()
}
func order_product(id string, quantity int) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if prod.GetQuantity() < quantity {
		fmt.Printf("There is only %d left on balance, we are temporarily unable to sell you the remaining %d\n\n", prod.GetQuantity(), quantity-prod.GetQuantity())
		prod.ChangeQuantity(0) 
	} else {
		prod.ChangeQuantity(prod.GetQuantity()-quantity) 
		fmt.Printf("You have succesfuly bought product with id: %s with quantity of:%d\n\n", id, quantity)
	}
	prod.AddNumberOfOrdered(quantity) 
	prod.AppendOrderHistory(quantity)
	
}
func purchase_product(id string, quantity int, price int) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	prod.AppendPurchaseHistory(price,quantity)
	prod.ChangeQuantity(prod.GetQuantity()+quantity) 
	fmt.Printf("A purchase of %d units of product with ID %s was made at a price of %d each.\n\n", quantity, id, price)
}

func get_quantity_of_product(id string) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quantity of product: %s is: %d\n\n", id, prod.GetQuantity())
}
func get_average_price(id string) float64 {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var k float64 = prod.AverageOrderedPrice()
	return k
}
func get_product_profit(id string) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		return
	}
	order_profit := get_average_price(id)


	purch_lose := prod.AveragePurchasePrice()

	profit_per_1 := order_profit - purch_lose
	fmt.Printf("Average price of Purchased Product with ID:%s is: %2f\n", id, purch_lose)
	fmt.Println("Profit is: ", profit_per_1*float64(prod.GetNumberOfOrdered()))
	fmt.Println()
}
func get_fewest_product() {
	fewest,err := cache.FindFewest()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("currently fewest number of product has: ",fewest)
	fmt.Println()
}
func get_most_popular_product() {
	popular,err := cache.FindPopular()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("currently Most Popular Product Is: ",popular)
	fmt.Println()
}

func get_orders_report(){
	data := cache.GetProductMap()
	for id,val := range data{
		fmt.Printf("Orders Report for product with id: %s\n",id)
		orderHistory := val.GetOrdersHistory()
		COGS := val.AveragePurchasePrice()
		for i,trans := range orderHistory{
			fmt.Printf("\t%d)Product Id: %s,Product Name: %s, Selling Price: %d,Quantity: %d,COGS: %f\n",
			i,
			val.GetId(),
			val.GetName(),
			trans.GetTransPrice(),
			trans.GetTransQuantity(),
			COGS)
		}
		
	}
}