package services 

import (
	"POS/entitys"
	"POS/ports"
)

type ItemsService struct {
	repo ports.ItemsRepository
}
func NewItemsService(repo ports.ItemsRepository) *ItemsService{
	return &ItemsService{repo: repo,}
}

func (i *ItemsService) AddItems(e entitys.ItemsAdd) string {
	return i.repo.AddItems(e)
}

func (i *ItemsService) ShowAllProducts(userID int)  []entitys.ItemsAdd{
	return i.repo.ShowAllProducts(userID)
}
func (i *ItemsService) OrderingAdd(orderItems entitys.TotalProducts) string{
	return i.repo.OrderingAdd(orderItems)
}
func (i *ItemsService) TotalingProducts(userID int) float64{
	return i.repo.TotalingProducts(userID)
}

func (i *ItemsService) BestSellingProducts(userID int) []entitys.TotalProducts{
	return i.repo.BestSellingProducts(userID)
}