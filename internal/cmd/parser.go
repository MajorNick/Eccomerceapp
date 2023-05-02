package cmd

import (
	"bufio"
	"fmt"
	"strconv"

	"os"
	"strings"
)


func ParseConsole ()(string,error){
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.TrimSpace(line)
	arguments := strings.Fields(line)
	callFunction(arguments)
	return line, nil
}

func callFunction(arguments []string){
	if len(arguments)==0{
		
		return
	}
	switch arguments[0]{
	case "save_product":
		if wrong_arguments(len(arguments),4){
			return
		}
		
		price,err :=strconv.Atoi(arguments[3])
		if err!=nil{
			fmt.Println("ERROR IN save_product command,  Price must be integer")
		}else{
			save_product(arguments[1],arguments[2],price)
		}

	case "purchase_product":
		if wrong_arguments(len(arguments),4){
			return
		}
		quantity,err :=strconv.Atoi(arguments[2])
		if err!=nil{
			fmt.Println("ERROR IN purchase_product command,  quantity must be integer")
		}
		price,err := strconv.Atoi(arguments[3])
		if err!=nil{
			fmt.Println("ERROR IN purchase_product command,  price must be integer")
		}
		purchase_product(arguments[1],quantity,price)
	case "order_product":
		if wrong_arguments(len(arguments),3){
			return
		}
		quantity,err :=strconv.Atoi(arguments[2])
		if err!=nil{
			fmt.Println("ERROR IN order_product command,  quantity must be integer")
		}else{
			order_product(arguments[1],quantity)
		}
		
	case "get_quantity_of_product":
		if wrong_arguments(len(arguments),2){
			return
		}
		get_quantity_of_product(arguments[1])
	case "get_average_price":
		if wrong_arguments(len(arguments),2){
			return
		}

		k := get_average_price(arguments[1])
		fmt.Printf("Average price of ordered Product with ID:%s is: %2f\n\n", arguments[1], k)
	case "get_product_profit":
		if wrong_arguments(len(arguments),2){
			return
		}
		get_product_profit(arguments[1])
	case "get_fewest_product":
		if wrong_arguments(len(arguments),1){
			return
		}
		get_fewest_product()
	case "get_most_popular_product":
		if wrong_arguments(len(arguments),1){
			return
		}
		get_most_popular_product()
	case "get_orders_report":
		if wrong_arguments(len(arguments),1){
			return
		}
		get_orders_report()
	case "export_orders_report":
		if wrong_arguments(len(arguments),2){
			return
		}
		export_orders_report(arguments[1])
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Enter Valid Command")
		Help()
	}


}
func Help(){
	fmt.Println(`
Use this commands:
save_product {product_id} {product_name} {price}
purchase_product {product_id} {quantity} {price}
order_product {product_id} {quantity}
get_quantity_of_product {product_id}
get_average_price {product_id}
get_product_profit {product_id}
get_fewest_product
get_most_popular_product
get_orders_report
export_orders_report {path}
`)
}

func wrong_arguments(current,expected int) bool{
	if current != expected{
		fmt.Println("Wrong number of Arguments")

		return true
	}
	return false 
}