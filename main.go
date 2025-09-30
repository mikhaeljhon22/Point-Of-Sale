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
  db.AutoMigrate(&models.UsersPos{},&models.ItemsAdd{})

  service := services.NewUserService(db)
  mailService := services.NewMailService()
  itemService := services.NewItemsService(db)
  controller := controllers.NewUserController(service,mailService)
  itemsController := controllers.NewItemsController(itemService)

  http.HandleFunc("/api/signup", controller.SignUpAddUser)
  http.HandleFunc("/api/signin", controller.Signin)
 // http.HandleFunc("/api/profile", middlewares.JWTVerif(controller.Profile))
  http.HandleFunc("/verif/",controller.Verification)
  http.HandleFunc("/api/add/item", itemsController.ItemAdd)
  
  routes := map[string]http.HandlerFunc{
    "/api/profile": controller.Profile,
    "api/add/item": itemsController.ItemAdd,
  }

  for path, handler := range routes{
    http.HandleFunc(path, middlewares.JWTVerif(handler))
  }

  configs.NewRedisClient()
	// http.HandleFunc("/", RouteHandler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}