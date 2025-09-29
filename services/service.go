package services 
import (
	"gorm.io/gorm"
	"sync"
	"POS/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *gorm.DB
	mutex sync.Mutex
}
func NewUserService(db *gorm.DB) *UserService{
	return &UserService{db: db,}
}



func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
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

func (s *UserService) SigninUser(user models.UsersPos) bool{
var foundUsers models.UsersPos
s.db.Where("username = ?", user.Username).First(&foundUsers)
checKpW := CheckPasswordHash(user.Password,foundUsers.Password)
return checKpW
}