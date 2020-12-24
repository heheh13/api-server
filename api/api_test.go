package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Requests struct {
	Method             string
	URL                string
	Body               io.Reader
	ExpectedStatusCode int
	ExpectedResponse   string
}

func Test_getUsers(t *testing.T) {
	var requests []Requests
	requests = append(requests, Requests{
		Method:             "GET",
		URL:                "http://localhost:8080/api/users",
		Body:               nil,
		ExpectedStatusCode: 200,
		ExpectedResponse:   "",
	})
	fmt.Println("need auth")
	for _, test := range requests {
		req, err := http.NewRequest(test.Method, test.URL, test.Body)
		if err != nil {
			log.Fatal(err)
		}
		res := httptest.NewRecorder()
		getUsers(res, req)
		fmt.Println(res.Body)
		if res.Result().StatusCode != test.ExpectedStatusCode {
			log.Printf(" %d %d", res.Result().StatusCode, test.ExpectedStatusCode)
		}
	}
}

func Test_getUser(t *testing.T) {
	go StartServer(8080)
	fmt.Println("HI")
	var requests []Requests
	requests = append(requests, Requests{
		Method:             "GET",
		URL:                "http://localhost:8080/api/users/1",
		Body:               nil,
		ExpectedStatusCode: 200,
		ExpectedResponse:   "user 1",
	})
	requests = append(requests, Requests{
		Method:             "GET",
		URL:                "http://localhost:8080/api/users/2",
		Body:               nil,
		ExpectedStatusCode: 200,
		ExpectedResponse:   "user 2",
	})
	processTest(t, requests)
}

func Test_createUser(t *testing.T) {
	// go StartServer(8080)
	var requests []Requests
	requests = append(requests, Requests{
		Method:             "POST",
		URL:                "http://localhost:8080/api/users",
		Body:               strings.NewReader(`{"id":"6","name":"heheh"}`),
		ExpectedStatusCode: 201,
		ExpectedResponse:   "",
	})
	processTest(t, requests)

}

func Test_updateUser(t *testing.T) {
	// go StartServer(8080)
	var requests []Requests
	requests = append(requests, Requests{
		Method:             "PUT",
		URL:                "http://localhost:8080/api/users/1",
		Body:               strings.NewReader(`{"name":"heheh"}`),
		ExpectedStatusCode: 201,
		ExpectedResponse:   "",
	})
	processTest(t, requests)

}

func Test_deleteUser(t *testing.T) {
	// go StartServer(8080)
	var requests []Requests
	requests = append(requests, Requests{
		Method:             "DELETE",
		URL:                "http://localhost:8080/api/users/1",
		Body:               nil,
		ExpectedStatusCode: 200,
		ExpectedResponse:   "",
	})
	//requests = append(requests, Requests{
	//	Method:             "DELETE",
	//	URL:                "http://localhost:8080/api/users/1",
	//	Body:               nil,
	//	ExpectedStatusCode: 204,
	//	ExpectedResponse:   "",
	//})
	processTest(t, requests)
}

func processTest(t *testing.T, requests []Requests) {
	client := http.DefaultClient
	for _, req := range requests {
		request, err := http.NewRequest(req.Method, req.URL, req.Body)
		if err != nil {
			log.Fatal(err)
		}

		request.Header.Add("Content-type", "application/json")
		request.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("username:password")))
		fmt.Println("header = ", request.Header, request.URL, request.Method)

		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		var here map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&here)
		fmt.Println(here)

		if resp.StatusCode != req.ExpectedStatusCode {
			log.Printf("expected %d found %d", req.ExpectedStatusCode, resp.StatusCode)
		}
	}
	time.Sleep(time.Second)

}

func Test_deleteUser1(t *testing.T) {
	//go StartServer(8080)

	type args struct {
		Method             string
		URL                string
		Body               io.Reader
		ExpectedStatusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				Method:             http.MethodDelete,
				URL:                "http://localhost:8080/api/users/2",
				Body:               nil,
				ExpectedStatusCode: 200,
			},
		},
	}
	client := http.DefaultClient
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.Method, tt.args.URL, tt.args.Body)
			req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("username:password")))
			if err != nil {
				t.Fatal(err)
			}
			res, err := client.Do(req)
			assert.NoError(t, err)
			assert.Equal(t, res.StatusCode, tt.args.ExpectedStatusCode)
		})
	}
}
