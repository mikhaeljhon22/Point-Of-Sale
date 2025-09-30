package middlewares 
import (
	"net/http"
	"github.com/kataras/jwt"
	"strings"

)

var sharedKey = []byte("sercrethatmaycontainch@r$32chars")

type TokenClaims struct {
	TokenClaims string `json:"tokenClaims"`
}

func JWTVerif(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
	partHeader := strings.Split(authHeader, "Bearer ")
	jwtToken := partHeader[1]


	_, err := jwt.Verify(jwt.HS256, sharedKey, []byte(jwtToken))
	if err != nil {
		panic(err)
	}
	
		next(w, r)
	}
}