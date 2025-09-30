package controllers 

import  (
	"encoding/json"
	"net/http"
	"POS/services"
	"github.com/olebedev/emitter"
	"POS/models"
	"fmt"
	"time"
	"github.com/kataras/jwt"
	"strings"
	"context"
	"POS/configs"
	"github.com/google/uuid"
)

type UserController struct {
  s *services.UserService
  mailService *services.MailService
}

func NewUserController(s *services.UserService,
	mailService *services.MailService,) *UserController{
	return &UserController{s: s, 
		mailService: mailService,}
}

var SharedKey = []byte("sercrethatmaycontainch@r$32chars")

type TokenClaims struct {
	TokenClaims string `json:"tokenClaims"`
	UserID int `json:"userID"`
}

func (c *UserController) SignUpAddUser(w http.ResponseWriter, r *http.Request){
	
	type Request struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	
	type Response struct {
		Message string `json:"message"`
	}
	
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	
	e := &emitter.Emitter{}
	e.On("Signup", func(event *emitter.Event){
		data := event.Args
		password := data[2].(string)
		
		hashPw,_ := services.HashPassword(password)
		user := models.UsersPos{
			Username: data[0].(string),
			Email: data[1].(string),
			Password: hashPw,
		}
		signup := c.s.SignUpAddUser(user)
		uuid := uuid.New()
		meUuid := uuid.String()
		//send mail confirmation
		c.s.StoreCodeVerif(data[0].(string),meUuid)
		c.mailService.SendEmail(data[1].(string),"Confirmation Verification","http://localhost:8081/verif" + meUuid)
		
		resp := Response{Message: fmt.Sprintf("%v", signup)}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
		e.Emit("Signup",req.Username,req.Email, req.Password)


	}
	
	
	func (c *UserController) Signin(w http.ResponseWriter, r *http.Request){
		type Response struct {
			Message string `json:"message"`
		}
		
		type Request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		
		user := models.UsersPos{
			Username: req.Username,
			Password: req.Password,
		}
	
		userID:= c.s.FindUserID(req.Username)

		fmt.Println("user id is", userID)
	myClaims := TokenClaims{
		TokenClaims: req.Username,
		UserID: userID,
	}
	
	token, err := jwt.Sign(jwt.HS256, SharedKey, myClaims, jwt.MaxAge(24* 10 * time.Hour))
	if err != nil {
		panic(err)
	}
	
	login := c.s.SigninUser(user)
	if(login == true){
		w.Header().Set("Authorization", "Bearer " + string(token))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("success login")
		}else{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Failed to login")
			
		}
	}
	
	func (c *UserController) Profile(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Username string `json:"username"`
		Email    string `json:"email"`
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

	username := claims.TokenClaims
	dataProfile := c.s.ProfileUser(username)

	rdb := configs.NewRedisClient()
	key := "profile:" + username

	res, err := rdb.Get(context.Background(), key).Result()
	if res == "" {
		value := fmt.Sprintf("%s,%s,%s",
			dataProfile.Username,
			dataProfile.Email,
		)

		op1 := rdb.Set(context.Background(), key, value, 10*24*time.Hour)
		if err := op1.Err(); err != nil {
			http.Error(w, "unable to cache data", http.StatusInternalServerError)
			return
		}

		res = value
	} else if err != nil {
		http.Error(w, "redis error", http.StatusInternalServerError)
		return
	}

	resp := Response{
		Username: dataProfile.Username,
		Email:    dataProfile.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (c *UserController) Verification(w http.ResponseWriter, r *http.Request){
	const prefix = "/verif/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.NotFound(w, r)
		return
	}

	pathVerif := strings.TrimPrefix(r.URL.Path, prefix)
	pathVerif = strings.TrimSuffix(pathVerif, "/")

	c.s.VerifyCode(pathVerif)	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success to create account")
}