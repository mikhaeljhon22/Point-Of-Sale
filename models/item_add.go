package models

import (
    "time"
)

type ItemsAdd struct {
    ID          int             `gorm:"primaryKey" json:"id"`
    UserID      int             `json:"user_id"`
    Item_name   string          `json:"item_name"`
    Stock       int             `json:"stock"`
    SKU         string          `json:"sku"`
    Price       string          `json:"price"`
    Random_code string          `json:"random_code"`
    CreatedAt   time.Time       `json:"created_at"`
    User        UsersPos        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"` // <-- disembunyikan
    Total       []TotalProducts `gorm:"foreignKey:ItemID" json:"total"`
}
