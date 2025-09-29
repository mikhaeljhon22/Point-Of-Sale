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

)

type UserController struct {
  s *services.UserService
}

func NewUserController(s *services.UserService) *UserController{
	return &UserController{s: s,}
}

var sharedKey = []byte("sercrethatmaycontainch@r$32chars")

type TokenClaims struct {
	TokenClaims string `json:"tokenClaims"`
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
	
	myClaims := TokenClaims{
		TokenClaims: req.Username,
	}

	token, err := jwt.Sign(jwt.HS256, sharedKey, myClaims, jwt.MaxAge(15*time.Minute))
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