package cmd

import (
	cache "Europe/internal/cache"
	_ "Europe/internal/model"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)
// save_product fn to save product id,name and price in catalogue

func save_product(id string, name string, price int)  {
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
//Place an order for the product, decreasing its balance
//according to the specified quantity.
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
		fmt.Printf("You have successfully sold product with id: %s with quantity of:%d\n\n", id, quantity)
	}
	prod.AddNumberOfOrdered(quantity) 
	prod.AppendOrderHistory(quantity)
	
}
//Purchase a product, increasing its balance
//based on the specified quantity.
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
//get remaining quantity of product
func get_quantity_of_product(id string) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quantity of product: %s is: %d\n\n", id, prod.GetQuantity())
}

//calculating average price of ordered product
func get_average_price(id string) float64 {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var k float64 = prod.AverageOrderedPrice()
	return k
}
//calculating profit in su,by calculating average ordered and purchased price then
// calculating profit per product and multiplyng it to the number of ordered product
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
	//fmt.Printf("Average price of Purchased Product with ID:%s is: %2f\n", id, purch_lose)
	fmt.Println("Profit is: ", profit_per_1*float64(prod.GetNumberOfOrdered()))
	fmt.Println()
}
// get fewest number of product 
func get_fewest_product() {
	fewest,err := cache.FindFewest()
	if err != nil{
		fmt.Println(err)
		return 
	}
	fmt.Println("currently fewest number of product has: ",fewest)
	fmt.Println()
}

// get mostly ordered product 
func get_most_popular_product() {
	popular,err := cache.FindPopular()
	if err != nil{
		fmt.Println(err)
		return
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
			fmt.Printf("\t%d)Product Id: %s,Product Name: %s,Average Purchase Price: %f, Selling Price: %d,Quantity: %d,COGS: %f\n",
			i+1,
			val.GetId(),
			val.GetName(),
			val.AveragePurchasePrice(),
			trans.GetTransPrice(),
			trans.GetTransQuantity(),
			COGS)
		}
		
	}
}

// using built in encoding/csv library to export data in csv file
// turn every order to slice of string and then flush it using writer interface
func export_orders_report(path string){
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
    	fmt.Println("Failed to create file:",err)
	}
	data := cache.GetProductMap()
	writer := csv.NewWriter(file)
	titles := []string{"ID","Name","Average Purchase Price:","Price","Quantity","COGS"}
	writer.Write(titles)
	for _,val := range data{
		
		orderHistory := val.GetOrdersHistory()
		avgPurch := val.AveragePurchasePrice()
		for _,trans := range orderHistory{
			row := []string{val.GetId(),val.GetName(),strconv.FormatFloat(avgPurch,'f',3,64),strconv.Itoa(trans.GetTransPrice()),strconv.Itoa(trans.GetTransQuantity()),strconv.FormatFloat(avgPurch,'f',3,64)}
			writer.Write(row)

		}
		writer.Flush()
		
	}
	
}