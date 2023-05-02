package internal

import (
	model "Europe/internal/model"
	"errors"
	"math"
)
// map for storing data of products
var data map[string]*model.Product

var ErrorNotExists error = errors.New("data: No Matching Product found")

func Init(){
	data = make(map[string]*model.Product)
}
func GetProductMap()map[string]*model.Product{
	return data
}

func GetProductById(id string)(*model.Product,error){
	product,ok := data[id]
	if !ok{
		return nil,ErrorNotExists
	}
	return product,nil
}

func AddInDataStore(id string,name string,price int) {

	data[id] = model.CreateProdInstance(id,price,name)
}
// iterating over map and finding product which has a minimal number of quantity
func FindFewest() (string,error){
	answ := ""
	quant := math.MaxInt64
	for _,v := range data{
		if quant>v.GetQuantity(){
			quant = v.GetQuantity()
			answ = v.GetName()
		}
	}
	if quant == math.MaxInt64{
		return "",errors.New("products not added")
	}else{
		return answ,nil
	}
}
// iterating over map and finding product which has a maximal number of ordered
func FindPopular() (string,error){
	answ := ""
	quant := math.MinInt64
	for _,v := range data{
		if quant<v.GetQuantity(){
			quant = v.GetNumberOfOrdered()
			answ = v.GetName()
		}
	}
	if quant == math.MinInt64{
		return "",errors.New("products not added")
	}else{
		return answ,nil
	}
}

