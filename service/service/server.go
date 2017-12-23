package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	n.Use(&checkHandler{})
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	//User check log in, log in, log out, delete
	mx.HandleFunc("/v1/user/login", checkLoginHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/user/login", loginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/user/logout", logoutHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/user/account", deleteUserHandler(formatter)).Methods("DELETE")
	//for users to register
	mx.HandleFunc("/v1/users", registerHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/users", listUsersHandler(formatter)).Methods("GET")
	//meetings
	mx.HandleFunc("/v1/meetings", cmHandler(formatter)).Methods("POST")
}
