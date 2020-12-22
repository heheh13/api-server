package auth

import "net/http"

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

var secretKey = "mySecretkey"

func generateJWT() {

}
