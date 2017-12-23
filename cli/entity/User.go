package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phonenumber"`
}

// for register
func CreateUser(user *User) (err error) {
	var code int
	var responseBody struct {
		Message string `json:"message"`
	}
	code, err = request("POST", "/v1/users", user, &responseBody)
	if err != nil {
		return
	}
	// 201
	if code == http.StatusCreated {
		return
	}
	err = fmt.Errorf("%s", responseBody.Message)
	return
}

// for delete user
func DeleteUser() (err error) {
	var code int
	code, err = request("DELETE", "/v1/user/account", nil, nil)
	if err != nil {
		return
	}
	// 20	
	if code == http.StatusOK {
		return
	}
	err = fmt.Errorf("%d", code)
	return
}

func ListAllUsers() (user []User, err error) {
	var code int
	code, err = request("GET", "/v1/users", nil, &user)
	if err != nil {
		return
	}
	// 200
	if code == http.StatusOK {
		return user, nil
	}
	err = fmt.Errorf("%d", code)
	return
}

var (
	scheme = "http"
	host   = "localhost:8080"
	//scheme = "https"
	//host = "private-89d1b-agenda28.apiary-mock.com"
)
var (
	// ErrInternalServerError ..
	ErrInternalServerError = fmt.Errorf("Internal Server Error")
)

func request(method string, api string, reqBodyPtr interface{}, resBodyPtr interface{}) (code int, err error) {
	var reqBodyReader io.Reader
	if reqBodyPtr != nil {

		var byteBody []byte
		if byteBody, err = json.Marshal(reqBodyPtr); err != nil {
			panic(err)
		}
		reqBodyReader = bytes.NewReader(byteBody)
	}

	// build full url
	var fullURL *url.URL
	if fullURL, err = url.Parse(api); err != nil {
		panic(err)
	}
	fullURL.Scheme = scheme
	fullURL.Host = host

	var req *http.Request
	if req, err = http.NewRequest(method, fullURL.String(), reqBodyReader); err != nil {
		panic(err)
	}

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	code = res.StatusCode

	if resBodyPtr != nil {
		err = json.NewDecoder(res.Body).Decode(resBodyPtr)
	}
	if code >= http.StatusInternalServerError {
		err = ErrInternalServerError
	}
	return
}
