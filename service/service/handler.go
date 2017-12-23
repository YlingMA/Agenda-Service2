package service

import (
	"net/http"

	"github.com/YlingMA/Agenda-Service2/service/entities"

	"encoding/json"
	"fmt"

	"github.com/unrolled/render"
)

//checkLogin
var (
	notLoginAPI = map[string]bool{
		"/v1/user/login": true,
		"/v1/users":      true,
	}
)

type checkHandler struct{}

func (h *checkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, isLogin := notLoginAPI[r.URL.Path]
	if !isLogin && checklogin() == nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		next(w, r)
	}
}
func checklogin() *entities.User {
	currentuser := entities.CurrentServ.WhoLogin()
	if currentuser == nil {
		return nil
	}

	user := entities.UserServ.FindUser(currentuser.Username)
	if user == nil {
		return nil
	}
	return user
}
func checkLoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if user := checklogin(); user != nil {
			formatter.JSON(w, http.StatusOK, user)
		} else {
			formatter.JSON(w, http.StatusUnauthorized, struct{}{})
		}
	}
}

//loginHandler
func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var reqbody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(req.Body).Decode(&reqbody); err != nil {
			fmt.Print(err)
			w.WriteHeader(400)
			return
		}
		var badData struct {
			Message string `json:"message"`
		}
		user := entities.UserServ.FindUser(reqbody.Username)
		if user == nil || user.Password != reqbody.Password {
			badData.Message = "the username and password are not correct"
			formatter.JSON(w, http.StatusUnauthorized, badData)
		}
		usernow := entities.CurrentUser{
			Username: reqbody.Username,
		}
		var resbody struct {
			Username string `json:"username"`
			Message  string `json:"message"`
		}
		entities.CurrentServ.Add(&usernow)
		resbody.Username = usernow.Username
		resbody.Message = "succeed to log in"
		formatter.JSON(w, http.StatusOK, resbody)
	}
}

//logoutHandler
func logoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := entities.CurrentServ.WhoLogin()
		entities.CurrentServ.Delete(user.Username)
		w.WriteHeader(http.StatusOK)
	}
}

//deleteUserHandler
func deleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		currentuser := entities.CurrentServ.WhoLogin()
		user := entities.UserServ.FindUser(currentuser.Username)
		entities.UserServ.DeleteUser(user)
		logoutHandler(formatter)(w, req)
	}
}

//register
func checkUserName(user *entities.User) error {
	if entities.UserServ.FindUser(user.Username) != nil {
		return fmt.Errorf("the input username have been used, please try another one")
	}
	return nil
}
func registerHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user entities.User
		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			w.WriteHeader(401)
			return
		}

		var badData struct {
			Message string `json:"message"`
		}

		if err := checkUserName(&user); err != nil {
			badData.Message = err.Error()
			formatter.JSON(w, http.StatusBadRequest, badData)
			return
		}

		entities.UserServ.AddUser(&user)
		formatter.JSON(w, http.StatusCreated, struct{}{})
	}
}

//LIST ALL USERS
func listUsersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, entities.UserServ.Allusers())
	}
}

//cmHandler(formatter)).Methods("POST")
func cmHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var meeting entities.Meeting
		if err := json.NewDecoder(req.Body).Decode(&meeting); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//var badData struct {
		//	Message string `json:"message"`
		//}

		currentuser := entities.CurrentServ.WhoLogin()
		user := entities.UserServ.FindUser(currentuser.Username)
		meeting.SponsorName = user.Username
		entities.MeetingServ.Add(&meeting)
		entities.MeetingParticipatorServ.Add(&meeting)

		formatter.JSON(w, http.StatusCreated, struct{}{})
	}
}
