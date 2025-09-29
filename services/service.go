package services 
import (
	"gorm.io/gorm"
	"sync"
	"POS/models"

)

type UserService struct {
	db *gorm.DB
	mutex sync.Mutex
}
func NewUserService(db *gorm.DB) *UserService{
	return &UserService{db: db,}
}

func (s *UserService) SignUpAddUser(user models.UsersPos) bool {
	var existingUser models.UsersPos
	findUsername :=  s.db.Where("username = ?", user.Username).Or("email = ?", user.Email).Find(&existingUser)
	if(findUsername.RowsAffected == 1){
		return false
	}else{
		s.mutex.Lock()
		defer s.mutex.Unlock()
		 s.db.Create(&user)
		 return true
	}
}

func (s *UserService) SigninUser(user models.UsersPos){
	
}