package controllers

import (
	"encoding/json"
	"POS/services"
	"net/http"
	"github.com/google/uuid"
	"POS/entitys"
	"github.com/kataras/jwt"
	"strings"
	"strconv"
)

type ItemsController struct {
	i *services.ItemsService
}

func NewItemsController(i *services.ItemsService) *ItemsController {
	return &ItemsController{i: i}
}
func (c *ItemsController) ItemAdd(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Name  string `json:"name"`
		SKU   string `json:"sku"`
		Price int    `json:"price"`
		Stock int `json:"stock"`
	}
	type Response struct {
		Message string `json:"message"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// validasi JWT
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
	if err := verifiedToken.Claims(&claims); err != nil {
		http.Error(w, "invalid claims", http.StatusUnauthorized)
		return
	}
	userID := claims.UserID

	meUuid := uuid.New().String()

	done := make(chan string)
	go func() {
		items := entitys.ItemsAdd{
			UserID:      userID,
			Item_name:   req.Name,
			SKU:         req.SKU,
			Stock: req.Stock,
			Price:       strconv.Itoa(req.Price),
			Random_code: meUuid,
		}
		result := c.i.AddItems(items)
		done <- result
	}()

	respMessage := <-done
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: respMessage})
}

func (c *ItemsController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
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

func (c *ItemsController) OrderingAdd(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ItemName string `json:"itemName"`
		Amount   int    `json:"amount"`
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
	orderItems := entitys.TotalProducts{
		Item_name: req.ItemName,
		Amount:    req.Amount,
		UserID:    userID,
	}
	order := c.i.OrderingAdd(orderItems)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: order})
}

func (c *ItemsController) TotalingProducts(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Message float64 `json:"message"`
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
	total := c.i.TotalingProducts(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: total})
}


func (c *ItemsController) BestSellingProducts(w http.ResponseWriter, r *http.Request){
	type Response struct {
		Message []entitys.TotalProducts `json:"message"`
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
	bestSelling := c.i.BestSellingProducts(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: bestSelling})
}