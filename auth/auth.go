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
	now := func(w http.ResponseWriter, r *http.Request) {

		//will fetch from db
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

	return now
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

				if xx, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				} else {
					fmt.Println(xx)
				}

				return secretKey, nil
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(token.Valid, token.Method, token.Claims, token.Header, token.Signature)
			if token.Valid {
				next.ServeHTTP(response, request)
			}

		} else {
			log.Println("not authorized")
		}
	}

}

func hasJWT(request *http.Request) (bool, error) {
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
			return true, nil
		}
	}
	return false, nil

}
func hasBasicAuth(request *http.Request) bool {
	username, password, authOk := request.BasicAuth()
	if !authOk {
		return false
	}
	fmt.Println(username, password, authOk)
	if username == "username" && password == "password" {
		return true
	}
	return false
}

func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		//has a token?
		ok, err := hasJWT(request)
		if err != nil {
			fmt.Errorf("there was a error")
		}
		if ok {
			next.ServeHTTP(response, request)
			return
		}
		//basic auth
		basicAuth := hasBasicAuth(request)
		fmt.Println(basicAuth)
		if basicAuth {
			token, err := GenerateJWT()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(token)
			request.Header.Set("Token", token)
			next.ServeHTTP(response, request)

		} else {
			response.WriteHeader(401) //unauthorized
			return
		}
		//give user a token
	}
}
