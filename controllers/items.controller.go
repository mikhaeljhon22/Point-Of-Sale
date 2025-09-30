package controllres 
import (
	"POS/services"
)
type ItemsController struct {
	i *services.ItemsService
}

func NewItemsController(i *services.ItemsService) *ItemsController{
	return &ItemsController{i:i,}
}