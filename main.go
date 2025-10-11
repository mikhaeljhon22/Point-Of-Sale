package main 
import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"POS/configs"
	"POS/entitys"
	"POS/services"
	"POS/controllers"
	"POS/middlewares"
  "POS/memory"
  "POS/db"
)


func main(){
 err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }


  database := configs.ConnectPostgre()
  database.AutoMigrate(&entitys.UsersPos{},&entitys.ItemsAdd{},
     &entitys.TotalProducts{}, &entitys.Person{})

  service := services.NewUserService(database)
  mailService := services.NewMailService()

  dbSave := db.NewSqlDB(database)
  itemService := services.NewItemsService(dbSave)
  repo := memory.NewMemory(database)
  dddService := services.NewDDDService(repo) 
  controller := controllers.NewUserController(service,mailService,dddService)
  itemsController := controllers.NewItemsController(itemService)

  http.HandleFunc("/api/signup", controller.SignUpAddUser)
  http.HandleFunc("/api/signin", controller.Signin)
 // http.HandleFunc("/api/profile", middlewares.JWTVerif(controller.Profile))
  http.HandleFunc("/verif/",controller.Verification)
  http.HandleFunc("/api/ddd", controller.TestDDD)
  
  routes := map[string]http.HandlerFunc{
    "/api/profile": controller.Profile,
    "/api/add/item": itemsController.ItemAdd,
    "/api/get/all/product": itemsController.GetAllProducts,
    "/api/add/order": itemsController.OrderingAdd,
    "/api/totaling": itemsController.TotalingProducts,
    "/api/best/selling": itemsController.BestSellingProducts,
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