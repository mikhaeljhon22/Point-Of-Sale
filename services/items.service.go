package services
import (
	"gorm.io/gorm"
	"sync"
	"POS/models"
)

type ItemsService struct{
	db *gorm.DB
	mutex sync.Mutex
}

func NewItemsService(db *gorm.DB) *ItemsService{
	return &ItemsService{db: db,}
}

func (s *ItemsService) AddItem(items models.ItemsAdd){

}