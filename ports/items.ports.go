package ports 

import (
	"POS/entitys"
)
type sqlDB struct {}

type ItemsRepository interface{
	AddItems(entitys.ItemsAdd) string 
	ShowAllProducts(userID int)  []entitys.ItemsAdd
	OrderingAdd(orderItems entitys.TotalProducts) string
	TotalingProducts(userID int) float64
	BestSellingProducts(userID int) []entitys.TotalProducts
}