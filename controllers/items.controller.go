package controllers 
import (
	"encoding/json"
	"POS/services"
	"net/http"
	"github.com/google/uuid"
	"POS/models"
	"github.com/kataras/jwt"
	"strings"
	"github.com/olebedev/emitter"
	"strconv"
)
type ItemsController struct {
	i *services.ItemsService
}

func NewItemsController(i *services.ItemsService) *ItemsController{
	return &ItemsController{i:i,}
}


func (c *ItemsController) ItemAdd(w http.ResponseWriter, r *http.Request){
	type Request struct {
		Name string `json:"name"`
		SKU string `json:"sku"`
		Price int `json:"price"`
	}
	type Response struct {
		Message string `json:"message"`
	}

	
		var req Request 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		authHeader := r.Header.Get("Authorization")
	    partHeader := strings.Split(authHeader, "Bearer ")
	if len(partHeader) < 2 {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	jwtToken := partHeader[1]

	verifiedToken, err := jwt.Verify(jwt.HS256, SharedKey, []byte(jwtToken))
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	var claims TokenClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		http.Error(w, "invalid claims", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID

	uuid := uuid.New()
	meUuid := uuid.String()

	e := &emitter.Emitter{}
	e.On("ItemAdd", func(event *emitter.Event){
	   data := event.Args

	   items := models.ItemsAdd{
		UserID: userID,
		Item_name: data[0].(string),
		SKU: data[1].(string),
		Price: strconv.Itoa(data[2].(int)),
		Random_code: meUuid,
	}

	add := c.i.AddItem(items)
	
	resp := Response{Message: add}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	})

	e.Emit("ItemAdd",req.Name,req.SKU,req.Price)
}

func (c *ItemsController) GetAllProducts(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
	    partHeader := strings.Split(authHeader, "Bearer ")
	if len(partHeader) < 2 {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	jwtToken := partHeader[1]

	verifiedToken, err := jwt.Verify(jwt.HS256, SharedKey, []byte(jwtToken))
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	var claims TokenClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		http.Error(w, "invalid claims", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID
	getAllProducts := c.i.ShowAllProducts(userID)


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAllProducts)
}

func (c *ItemsController) OrderingAdd(w http.ResponseWriter, r *http.Request){
	type Request struct {
		ItemName string `json:"itemName"`
		Amount int `json:"amount"`
	}
	type Response struct {
		Message string `json:"message"`
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "invalid request body", http.StatusBadRequest)
    return
}


	authHeader := r.Header.Get("Authorization")
	    partHeader := strings.Split(authHeader, "Bearer ")
	if len(partHeader) < 2 {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	jwtToken := partHeader[1]

	verifiedToken, err := jwt.Verify(jwt.HS256, SharedKey, []byte(jwtToken))
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	var claims TokenClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		http.Error(w, "invalid claims", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID
	orderItems := models.TotalProducts{
		Item_name: req.ItemName,
		Amount: req.Amount,
		UserID: userID,
	}
	order := c.i.OrderingAdd(orderItems)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: order})
}