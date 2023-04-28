package internal
import(
	model "Europe/internal/model"
	"errors"
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
		 
		PurchHistory: make([]*model.Trans,2),
		OrderHistory: make([]*model.Trans,2),
	}
	data[id] = &prod
}

