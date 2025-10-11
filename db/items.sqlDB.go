package db

import (
	"gorm.io/gorm"
	"sync"
	"POS/entitys"
	"fmt"
	"strconv"
)

type sqlDB struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewSqlDB(db *gorm.DB) *sqlDB {
	return &sqlDB{db: db,}
}

func (s *sqlDB) AddItems(items entitys.ItemsAdd) string {
	find := s.db.Where("item_name = ?", items.Item_name).First(&items)
	if find.RowsAffected == 0 {
		s.db.Create(&items)
		fmt.Println("ini")
		return "success to add item"
	} else {
		return "item already exists"
	}
}

func (s *sqlDB) ShowAllProducts(userID int) []entitys.ItemsAdd {
	var items []entitys.ItemsAdd
	s.db.Where("user_id = ?", userID).Find(&items)
	return items
}

func (s *sqlDB) OrderingAdd(orderItems entitys.TotalProducts) string {
	var item entitys.ItemsAdd
	var order entitys.TotalProducts

	find := s.db.Where("item_name = ?", orderItems.Item_name).First(&item)
	fmt.Println("row affected", find.RowsAffected)
	if find.RowsAffected == 1 {
		itemID := item.ID
		random_code := item.Random_code

		check := s.db.Where("id = ? AND stock >= ?", itemID, orderItems.Amount).First(&item)

		if check.RowsAffected == 1 {
			checkIfexistsProduct := s.db.Where("item_id = ? AND deleted = ?", itemID, false).First(&order)

			if checkIfexistsProduct.RowsAffected == 0 {
				priceInt, _ := strconv.Atoi(item.Price)
				tp := &entitys.TotalProducts{
					UserID:      orderItems.UserID,
					Item_name:   orderItems.Item_name,
					Amount:      orderItems.Amount,
					ItemID:      itemID,
					Random_code: random_code,
					Price:       priceInt * orderItems.Amount,
				}
				s.db.Create(tp)
				s.db.Model(&item).Where("id = ?", item.ID).Update("stock", item.Stock-orderItems.Amount)

			} else {
				return "already exists product"
			}
		} else {
			return "amount not enough"
		}
		return "success add item"
	} else {
		return "not found product"
	}
}

func (s *sqlDB) TotalingProducts(userID int) float64 {
	var order entitys.TotalProducts
	var totaling float64
	s.db.Raw("SELECT SUM(price::numeric) FROM total_products WHERE user_id = ? AND deleted = ? ", userID, false).Scan(&totaling)
	s.db.Model(&order).Where("user_id = ? AND deleted = ?", userID, false).Update("deleted", true)
	return totaling
}


	func (s *sqlDB) BestSellingProducts(userID int) []entitys.TotalProducts{
		var result []entitys.TotalProducts 
		s.db.Limit(2).Order("amount DESC").Where("user_id = ?", userID).Find(&result)
		return result
	}