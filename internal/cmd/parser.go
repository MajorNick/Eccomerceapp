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
	switch arguments[0]{
	case "save_product":
		
		price,err :=strconv.Atoi(arguments[3])
		if err!=nil{
			fmt.Println("ERROR IN save_product command,  Price must be integer")
		}else{
			save_product(arguments[1],arguments[2],price)
		}

	case "purchase_product":

	case "order_product":
		
	case "get_quantity_of_product":

		get_quantity_of_product(arguments[1])
	case "get_average_price":

	case "get_product_profit":

	case "get_fewest_product":

	case "get_most_popular_product":

	}


}