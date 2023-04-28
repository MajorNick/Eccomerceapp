package cmd

import (
	cache "Europe/internal/cache"
	"fmt"
)



func save_product(id string,name string,price int) {
	product,err := cache.GetProductById(id)
	if err != nil{
		if err == cache.ErrorNotExists{
			cache.AddInDataStore(id,name,price)
			fmt.Println("Product Succesfully added!")
		}else{
			fmt.Println(err)
		}
		return
	}
	product.Price = price
	fmt.Println("Product Price Succesfully Changed!")
}
func order_product(){

}

func get_quantity_of_product(id string){
	prod,err:= cache.GetProductById(id)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	fmt.Printf("Quantity of product: %s is: %d\n",id,prod.Quantity)
}
func get_average_price(){

}
func get_product_profit(){

}
func get_fewest_product(){

}
func get_most_popular_product(){
	
}

