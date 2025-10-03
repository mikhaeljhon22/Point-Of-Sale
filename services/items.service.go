package services
import (
	"gorm.io/gorm"
	"sync"
	"POS/models"
	"fmt"
	"strconv"

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
   if(find.RowsAffected == 0){
	s.db.Create(&items)
	fmt.Println("ini")
	return "success to add item"
   }else{
	return "item already exists"
   }
}
func (s *ItemsService) ShowAllProducts(userID int) []models.ItemsAdd{
	var items []models.ItemsAdd
	
	s.db.Where("user_id = ? ", userID).Find(&items)
	return items
} 	
func (s *ItemsService) OrderingAdd(orderItems models.TotalProducts) string{
	var item models.ItemsAdd
	var order models.TotalProducts

	find := s.db.Where("item_name = ?", orderItems.Item_name).First(&item)
	fmt.Println("row affected", find.RowsAffected)
	if(find.RowsAffected == 1){
		itemID := item.ID
		random_code := item.Random_code

		check := s.db.Where("id = ? AND stock >= ?", itemID, orderItems.Amount).First(&item)
		
		if(check.RowsAffected == 1){
		checkIfexistsProduct := s.db.Where("item_id = ? AND deleted = ?", itemID, false).First(&order)

		if(checkIfexistsProduct.RowsAffected == 0){
	priceInt,_:= strconv.Atoi(item.Price)
	tp := &models.TotalProducts{
	UserID: orderItems.UserID,
    Item_name:  orderItems.Item_name,
    Amount:     orderItems.Amount,
    ItemID:     itemID,
    Random_code: random_code,
	Price: priceInt  * orderItems.Amount,
      }
      s.db.Create(tp)
	  s.db.Model(&item).Where("id = ?", item.ID).Update("stock", item.Stock - orderItems.Amount)

	}else{
		return "already exists product"
	}
	}else{
		return "amount not enough"
	}
	return "success add item"
	}else{
		return "not found product"
	}
}

func (s *ItemsService) TotalingProducts(userID int) float64{
	var order models.TotalProducts
    var totaling float64
	s.db.Raw("SELECT SUM(price::numeric) FROM total_products WHERE user_id = ? AND deleted = ? ", userID, false).Scan(&totaling)
	s.db.Model(&order).Where("user_id = ? AND deleted = ?", userID,false).Update("deleted", true)
	return totaling
}