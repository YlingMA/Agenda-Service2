package entity

import (
	"fmt"
	"net/http"
)

// for log in
func Login(username string, password string) (err error) {
	requestBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{username, password}
	var responseBody struct {
		Message string `json:"message"`
	}
	var code int
	code, err = request("POST", "/v1/user/login", &requestBody, &responseBody)
	if err != nil {
		return err
	}
	// 200
	if code == http.StatusOK {
		return
	}
	err = fmt.Errorf("%s", responseBody.Message)
	return
}

// for log out
func Logout() (err error, flag bool) {
	var code int
	code, err = request("POST", "/v1/user/logout", nil, nil)
	if err != nil {
		return err, false
	}
	// 200
	if code == http.StatusOK {
		return nil, true
	}
	err = fmt.Errorf("%d", code)
	return err, false
}

// check if there exist current user
func Check_Login() (flag bool, err error) {
	var responseBody struct {
		Username string `json:"username"`
	}
	var code int
	code, err = request("GET", "/v1/user/login", nil, &responseBody)
	if err != nil {
		return false, err
	}
	flag = false
	// 200
	if code == http.StatusOK {
		flag = true
		return flag, nil
	}
	return flag, nil
}

// get current username
func GetCurrentUser() (username string, err error) {
	var responseBody struct {
		Username string `json:"username"`
	}
	var code int
	code, err = request("GET", "/v1/user/login", nil, &responseBody)
	if err != nil {
		return "", err
	}
	// 200
	if code == http.StatusOK {
		username = responseBody.Username
	}
	// 401
	if code == http.StatusUnauthorized {
		return "", err
	}
	return username, nil
}
