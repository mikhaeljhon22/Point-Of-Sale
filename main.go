package main 
import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"POS/configs"
	"POS/models"
	"POS/services"
	"POS/controllers"
	"POS/middlewares"
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
  http.HandleFunc("/api/signin", controller.Signin)
  http.HandleFunc("/api/profile", middlewares.JWTVerif(controller.Profile))
  configs.NewRedisClient()
	// http.HandleFunc("/", RouteHandler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}