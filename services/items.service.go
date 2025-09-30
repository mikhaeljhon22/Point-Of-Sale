package services
import (
	"gorm.io/gorm"
	"sync"
	"POS/models"
	"fmt"

)

type ItemsService struct{
	db *gorm.DB
	mutex sync.Mutex
}

func NewItemsService(db *gorm.DB) *ItemsService{
	return &ItemsService{db: db,}
}

func (s *ItemsService) AddItem(items models.ItemsAdd) string{
find := s.db.Where("item_name = ?", items.Item_name).First(&items)
fmt.Println(find.RowsAffected)
   if(find.RowsAffected == 0){
	s.db.Create(&items)
	fmt.Println("ini")
	return "success to add item"
   }else{
	return "item already exists"
   }
}

	