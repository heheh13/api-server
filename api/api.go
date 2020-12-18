package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "text")
	params := mux.Vars(r)
	for _, user := range data.Users {
		if user.ID == params["id"] {
			w.Write([]byte("hello"))
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	return

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, user := range data.Users {
		if user.ID == params["id"] {
			fmt.Printf("%T", r.Body)
			// var variable map[string]interface{}
			x := user
			json.NewDecoder(r.Body).Decode(&x)

			// fmt.Println()
			// fmt.Println()
			// fmt.Println(x)
			return

		}
	}
}

//StartServer starts the sever
func StartServer(Port int) {

	log.Printf("-------------starting sever at %d -------\n", Port)

	router := mux.NewRouter()
	router.HandleFunc("/", homepage)
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Port), router))
}
