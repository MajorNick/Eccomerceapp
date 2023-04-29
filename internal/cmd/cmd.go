package cmd

import (
	cache "Europe/internal/cache"
	model "Europe/internal/model"
	"fmt"
)

func save_product(id string, name string, price int) {
	product, err := cache.GetProductById(id)
	if err != nil {
		if err == cache.ErrorNotExists {
			cache.AddInDataStore(id, name, price)
			fmt.Println("Product Succesfully added!")
		} else {
			fmt.Println(err)
		}
		return
	}
	product.Price = price
	fmt.Println("Product Price Succesfully Changed!")
}
func order_product(id string, quantity int) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if prod.Quantity < quantity {
		fmt.Printf("There is only %d left on balance, we are temporarily unable to sell you the remaining %d\n", prod.Quantity, quantity-prod.Quantity)
		prod.Quantity = 0
	} else {
		prod.Quantity -= quantity
		fmt.Printf("You have succesfuly bought product with id: %s with quantity of:%d\n", id, quantity)
	}

	*prod.OrderHistory = append(*prod.OrderHistory, &model.Trans{Quantity: quantity, Price: prod.Price})
}
func purchase_product(id string, quantity int, price int) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	*prod.PurchHistory = append(*prod.PurchHistory, &model.Trans{Quantity: quantity, Price: price})
	prod.Quantity += quantity
	fmt.Printf("A purchase of %d units of product with ID %s was made at a price of %d each.\n", quantity, id, price)
}

func get_quantity_of_product(id string) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quantity of product: %s is: %d\n", id, prod.Quantity)
}
func get_average_price(id string) float64 {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	cnt := 0
	sum := 0
	for _, v := range *prod.OrderHistory {
		cnt += v.Quantity
		sum += v.Price * v.Quantity
	}
	var k float64 = float64(sum / cnt)
	fmt.Printf("Average price of ordered Product with ID:%s is: %2f\n", id, k)
	return k
}
func get_product_profit(id string) {
	prod, err := cache.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	order_profit := get_average_price(id)
	var purch_lose float64
	cnt, sum := 0, 0
	for _, v := range *prod.PurchHistory {
		cnt += v.Quantity
		sum += v.Quantity * v.Price
	}
	purch_lose = float64(sum / cnt)
	ordered := 0
	for _, v := range *prod.OrderHistory {
		ordered += v.Quantity
	}
	profit_per_1 := order_profit - purch_lose
	fmt.Printf("Average price of Purchased Product with ID:%s is: %2f\n", id, purch_lose)
	fmt.Println("Profit is: ", profit_per_1*float64(ordered))
}
func get_fewest_product() {
	fewest,err := cache.FindFewest()
	if err != nil{
		fmt.Println(fewest)
	}
	fmt.Println("currently fewest number of product has: ",fewest)
}
func get_most_popular_product() {

}
