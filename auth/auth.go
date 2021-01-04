package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

////BasicAuth is for authenticating basic user
//func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
//	now := func(response http.ResponseWriter, request *http.Request) {
//
//		//will fetch from db
//		user := "username"
//		pass := "password"
//		username, password, authOk := request.BasicAuth()
//		if !authOk {
//			http.Error(response, "not authorized", http.StatusUnauthorized)
//			return
//		}
//		if username != user || password != pass {
//			http.Error(response, "not authorized", http.StatusUnauthorized)
//			return
//		}
//		next.ServeHTTP(response, request)
//
//	}
//
//	return now
//}
//Stores the secretkey for jwt check
var secretKey []byte

//UserName get the user name from the environment file
var UserName string

//Password get the user name from the environment file
var Password string

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

//
////AuthenticatedWithJWT is
//func AuthenticatedWithJWT(next http.HandlerFunc) http.HandlerFunc {
//	return func(response http.ResponseWriter, request *http.Request) {
//		if request.Header["Token"] != nil {
//
//			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
//
//				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//					return nil, fmt.Errorf("there was an error")
//				}
//
//				return secretKey, nil
//
//			})
//			if err != nil {
//				fmt.Println(err)
//			}
//			fmt.Println(token.Valid, token.Method, token.Claims, token.Header, token.Signature)
//			if token.Valid {
//				next.ServeHTTP(response, request)
//			}
//
//		} else {
//			log.Println("not authorized")
//		}
//	}
//
//}

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
	fmt.Println(UserName, Password)
	if username == UserName && password == Password {
		return true
	}
	return false
}
func validateCookie(value string) (bool, error) {
	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return secretKey, nil
	})
	if err != nil {
		log.Println(err)
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, nil
}

//IsAuthenticated is a function
func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		//has cookies?
		for _, cookie := range request.Cookies() {
			if cookie.Name == "Token" {
				ok, err := validateCookie(cookie.Value)
				if err != nil {
					log.Println(err)
				}
				if ok {
					log.Println("serve because of cookies")
					next.ServeHTTP(response, request)
					return
				}
			}
		}
		//has token?
		ok, err := hasJWT(request)
		if err != nil {
			log.Fatal("there was a error finding token")
		}
		if ok {
			log.Println("serve because of has jwt token")
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
			cookie := http.Cookie{
				Name:  "Token",
				Value: token,
				Path:  "/",
			}
			http.SetCookie(response, &cookie)
			fmt.Println(token)
			next.ServeHTTP(response, request)
			return

		}
		response.WriteHeader(http.StatusUnauthorized) //unauthorized
		return
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	//UserName get the user name from the environment file
	UserName = os.Getenv("UserName")

	//Password get the user name from the environment file
	Password = os.Getenv("Password")

	//SecretKey stores the key to validate jwt
	secretKey = []byte(os.Getenv("SecretKey"))
}
