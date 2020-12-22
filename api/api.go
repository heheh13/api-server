package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heheh13/api-server/auth"
	"github.com/heheh13/api-server/data"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home page!")
	fmt.Println(data.Users)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, user := range data.Users {
		if user.ID == params["id"] {
			// w.Write([]byte("hello"))
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	return

}
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser data.Profile
	// reqBody, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(reqBody, &newUser)
	json.NewDecoder(r.Body).Decode(&newUser)
	data.Users = append(data.Users, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data.Users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// reqBody, _ := ioutil.ReadAll(r.Body)

	// fmt.Printf("%+v", string(reqBody))
	for idx, user := range data.Users {
		if user.ID == params["id"] {
			//need to verify payload and update?
			//can do a lot of things here
			// but now writing a simple implementation
			updatedUser := user
			json.NewDecoder(r.Body).Decode(&updatedUser)

			// Replaced with same id
			updatedUser.ID = user.ID
			fmt.Println(updatedUser)

			w.WriteHeader(http.StatusCreated)
			data.Users[idx] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	// sends an empty response
	w.WriteHeader(http.StatusNoContent)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			fmt.Println(data.Users)
			json.NewEncoder(w).Encode(data.Users)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)

}

var router = mux.NewRouter()

//StartServer starts the sever
func StartServer(Port int) {

	log.Printf("-------------starting sever at %d -------\n", Port)

	router.HandleFunc("/", auth.BasicAuth(homepage))
	router.HandleFunc("/api/users", auth.BasicAuth(getUsers)).Methods("GET")
	router.HandleFunc("/api/users/{id}", auth.BasicAuth(getUser)).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Port), router))
}
