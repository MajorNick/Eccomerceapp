package internal

import (
	model "Europe/internal/model"
	"errors"
	"math"
)
var data map[string]*model.Product
var ErrorNotExists error = errors.New("data: No Matching Product found")
func Init(){
	data = make(map[string]*model.Product)
	
}

func GetProductById(id string)(*model.Product,error){
	product,ok := data[id]
	if !ok{
		return nil,ErrorNotExists
	}
	return product,nil
}

func AddInDataStore(id string,name string,price int) {
	prod := model.Product{
		Id:id,
		Name:name,
		Price:price,
		Quantity:0,
		Ordered: 0,
		PurchHistory: &[]*model.Trans{},
		OrderHistory: &[]*model.Trans{},
	}
	data[id] = &prod
}

func FindFewest() (string,error){
	answ := ""
	quant := math.MaxInt64
	for _,v := range data{
		if quant>v.Quantity{
			quant = v.Quantity
			answ = v.Name
		}
	}
	if quant == math.MaxInt64{
		return "",errors.New("Products Not Added")
	}else{
		return answ,nil
	}
}

func FindPopular() (string,error){
	answ := ""
	quant := math.MinInt64
	for _,v := range data{
		if quant<v.Quantity{
			quant = v.Ordered
			answ = v.Name
		}
	}
	if quant == math.MaxInt64{
		return "",errors.New("Products Not Added")
	}else{
		return answ,nil
	}
}

