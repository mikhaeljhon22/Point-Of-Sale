package models 

type TotalProducts struct {
	ID int `gorm:"primaryKey"`
	UserID int 
	ItemID int
	Random_code string 
	Item_name string 
	Amount int 
	Price int		
	Deleted bool `gorm:"default:false"`
	User UsersPos `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
	Item ItemsAdd `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
}