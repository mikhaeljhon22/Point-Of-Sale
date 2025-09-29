package main 
import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"POS/configs"
	"POS/models"
	"POS/services"
	"POS/controllers"
)


func main(){
 err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }


  db := configs.ConnectPostgre()
  db.AutoMigrate(&models.UsersPos{})

  service := services.NewUserService(db)
  controller := controllers.NewUserController(service)

  http.HandleFunc("/api/signup", controller.SignUpAddUser)
	// http.HandleFunc("/", RouteHandler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}