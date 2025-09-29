package test 
import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"bytes"
	"POS/controllers"
	"POS/services"
	"POS/configs"

)
type Response struct {
	Message string `json:"message"`
}
func TestSignupHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		"POST",
		"/api/signup",
		bytes.NewBufferString(`{"username":"tes", "email":"mikhaeljhon22@gmail.com"}`),
	)
	
	db :=  configs.ConnectPostgre()
	svc := services.NewUserService(db)
	ctrl := controllers.NewUserController(svc)

	ctrl.SignUpAddUser(w, req)

	if w.Code != 200 {
		t.Fatal("bad status:", w.Code)
	}

	var r struct{ Message string }
	if err := json.NewDecoder(w.Body).Decode(&r); err != nil {
		t.Fatal("failed decode response:", err)
	}

	if r.Message != "false" {
		t.Fatal("bad message:", r.Message)
	}
}
