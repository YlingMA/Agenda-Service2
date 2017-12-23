package entities

import (
	"log"

	"github.com/jinzhu/gorm"
	//driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//interface of service
type serv interface {
	//create table
	load()
}

var (
	gormDb *gorm.DB
	//service for user,currentuser,meeting and meeting participators
	services []serv
)

func addServ(s serv) {
	services = append(services, s)
}
func init() {
	db, err := gorm.Open("sqlite3", "./agenda.db")
	checkErr(err)
	gormDb = db

	var err2 interface{}
	loadcomplete := make(chan bool)
	for _, s := range services {
		go func(s serv) {
			defer func() {
				if e := recover(); e != nil {
					err2 = e
				}
				loadcomplete <- true
			}()
			s.load()
		}(s)
	}
	for _ = range services {
		<-loadcomplete
	}
	if err2 != nil {
		log.Fatal(err2)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
