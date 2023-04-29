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
		wrong_arguments(len(arguments),4)
		
		price,err :=strconv.Atoi(arguments[3])
		if err!=nil{
			fmt.Println("ERROR IN save_product command,  Price must be integer")
		}else{
			save_product(arguments[1],arguments[2],price)
		}

	case "purchase_product":
		wrong_arguments(len(arguments),4)
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
		wrong_arguments(len(arguments),3)
		quantity,err :=strconv.Atoi(arguments[2])
		if err!=nil{
			fmt.Println("ERROR IN order_product command,  quantity must be integer")
		}else{
			order_product(arguments[1],quantity)
		}
		
	case "get_quantity_of_product":
		wrong_arguments(len(arguments),2)
		get_quantity_of_product(arguments[1])
	case "get_average_price":
		wrong_arguments(len(arguments),2)
		get_average_price(arguments[1])
	case "get_product_profit":
		wrong_arguments(len(arguments),2)
		get_product_profit(arguments[1])
	case "get_fewest_product":
		wrong_arguments(len(arguments),1)
		get_fewest_product()
	case "get_most_popular_product":
		wrong_arguments(len(arguments),1)
		get_most_popular_product()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Enter Valid Command")
	}


}

func wrong_arguments(current,expected int){
	if current != expected{
		fmt.Println("Wrong number of Arguments")
	} 
}