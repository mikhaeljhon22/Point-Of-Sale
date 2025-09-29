package controllers 

import  (
	"encoding/json"
	"net/http"
	"POS/services"
	"github.com/olebedev/emitter"
	"POS/models"
	"fmt"
)

type UserController struct {
  s *services.UserService
}

func NewUserController(s *services.UserService) *UserController{
	return &UserController{s: s,}
}

func (c *UserController) SignUpAddUser(w http.ResponseWriter, r *http.Request){

	type Request struct {
		Username string `json:"username"`
		Email string `json:"email"`
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

			user := models.UsersPos{
				Username: data[0].(string),
				Email: data[1].(string),
			}
			signup := c.s.SignUpAddUser(user)
resp := Response{Message: fmt.Sprintf("%v", signup)}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		})
		e.Emit("Signup",req.Username,req.Email)

}
