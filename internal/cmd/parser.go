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
		if len(arguments) != 4{
			fmt.Println("Wrong number of Arguments")
		}
		
		price,err :=strconv.Atoi(arguments[3])
		if err!=nil{
			fmt.Println("ERROR IN save_product command,  Price must be integer")
		}else{
			save_product(arguments[1],arguments[2],price)
		}

	case "purchase_product":
		if len(arguments) != 4{
			fmt.Println("Wrong number of Arguments")
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
		quantity,err :=strconv.Atoi(arguments[2])
		if err!=nil{
			fmt.Println("ERROR IN order_product command,  quantity must be integer")
		}else{
			order_product(arguments[1],quantity)
		}
		
	case "get_quantity_of_product":
		if len(arguments) != 2{
			fmt.Println("Wrong number of Arguments")
		}
		get_quantity_of_product(arguments[1])
	case "get_average_price":
		if len(arguments) != 2{
			fmt.Println("Wrong number of Arguments")
		}
		get_average_price(arguments[1])
	case "get_product_profit":
		if len(arguments) != 2{
			fmt.Println("Wrong number of Arguments")
		}
		get_product_profit(arguments[1])
	case "get_fewest_product":
		if len(arguments) != 1{
			fmt.Println("Wrong number of Arguments")
		}
		get_fewest_product()
	case "get_most_popular_product":

	default:
		fmt.Println("Enter Valid Command")
	}


}