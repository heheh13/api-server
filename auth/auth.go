package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//BasicAuth is for authenticating basic user
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := "username"
		pass := "password"
		username, password, authOk := r.BasicAuth()
		if !authOk {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		if username != user || password != pass {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	}
}

var secretKey = []byte("mySecretKey")

//GenerateJWT will genearate jwt
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "heheh"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyJWT verify the jwt
func VerifyJWT(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err

}

//AuthenticatedWithJWT is
func AuthenticatedWithJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {

			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return secretKey, nil
			})
			if err != nil {
				fmt.Println(err)
			}
			if token.Valid {
				next.ServeHTTP(response, request)
			}

		} else {
			log.Println("not authorized")
		}
	}

}
