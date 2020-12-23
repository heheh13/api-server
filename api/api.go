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

//the comments may unnecessary but for the sake of understandig :)

// its a demo homepage with the purpose of something on the home page
func homepage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "welcome to home page!")
	// fmt.Println(data.Users)
}

// getUsers gets all users form the database. it receive the request
// request contains a lot of field including method,url,body-> type of io reader
// body can be read using ioutil.readAll(request.Body) or can be decoded with json/encoding

func getUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	//json.newcoder recives a io writer and return a new encoder that wite to response
	//encode function encode the data into the response writer interface
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data.Users)
}

//getUser returns a specific user based on request
func getUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	//A syntax for reading url param form request
	// keys ,ok := request.URL.Query()  can receive the query string from urls
	//names are stored in map of routes variables

	params := mux.Vars(request)

	for _, user := range data.Users {
		if user.ID == params["id"] {
			// w.Write([]byte("hello"))
			json.NewEncoder(response).Encode(user) //write encoded user
			response.WriteHeader(http.StatusOK)
			return
		}
	}
	response.WriteHeader(http.StatusNoContent)
	return

}

// creating a new user
// didn't defined a actual way
// for the sake of learning

func createUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	var newUser data.Profile
	// reqBody, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(reqBody, &newUser)

	json.NewDecoder(request.Body).Decode(&newUser) // decode normal go struct format
	data.Users = append(data.Users, newUser)       //update in demo db
	fmt.Println(newUser)

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(newUser) //write an optional created data to response
}

//update an existing userpayload
// maybe the i need to receive the whole paylaod
// or can  create a patch request

func updateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	fmt.Println("update user called ", params)

	// reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Printf("%+v", string(reqBody))

	for idx, user := range data.Users {
		if user.ID == params["id"] {
			//need to verify payload and update?
			//can do a lot of things here
			// but now writing a simple implementation
			updatedUser := user
			json.NewDecoder(request.Body).Decode(&updatedUser)

			// Replaced with same id
			updatedUser.ID = user.ID
			fmt.Println(updatedUser)

			response.WriteHeader(http.StatusCreated)
			data.Users[idx] = updatedUser
			json.NewEncoder(response).Encode(updatedUser)
			fmt.Println(data.Users)
			return
		}
	}
	// sends an empty response
	response.WriteHeader(http.StatusNoContent)
}

// delete user
func deleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-Type", "application/json")
	params := mux.Vars(request)
	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			fmt.Println(data.Users)
			json.NewEncoder(response).Encode(data.Users)
			response.WriteHeader(http.StatusOK)
			return
		}
	}
	response.WriteHeader(http.StatusNoContent)

}

var router = mux.NewRouter()

//StartServer starts the sever
func StartServer(Port int) {

	log.Printf("-------------starting sever at %d -------\n", Port)

	router.HandleFunc("/", auth.IsAuthenticated(homepage)) //login
	router.HandleFunc("/api/users", auth.IsAuthenticated(getUsers)).Methods("GET")
	router.HandleFunc("/api/users/{id}", auth.IsAuthenticated(getUser)).Methods("GET")
	router.HandleFunc("/api/users", auth.IsAuthenticated(createUser)).Methods("POST")
	router.HandleFunc("/api/users/{id}", auth.IsAuthenticated(updateUser)).Methods("PUT")
	router.HandleFunc("/api/users/{id}", auth.IsAuthenticated(deleteUser)).Methods("DELETE")

	token, err := auth.GenerateJWT()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Port), router))
}
