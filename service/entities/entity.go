package entities

import (
	"github.com/jinzhu/gorm"
)

//User type
type User struct {
	Username string `json:"username" gorm:"primary_key;cloumn:username"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Phone    string `json:"phonenumber" gorm:"column:phonenumber"`
}

//Table
func (*User) TableName() string {
	return "users"
}

//User Service
type UserService struct{}

//UserServ
var UserServ = UserService{}

//create table
func (*UserService) load() {
	u := &User{}
	if !gormDb.HasTable(u) {
		gormDb.CreateTable(u)
	}
}

//add user
func (*UserService) AddUser(u *User) {
	tx := gormDb.Begin()
	checkErr(tx.Error)
	//if add user to table fail then rollback
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

//delete user
func (*UserService) DeleteUser(u *User) {
	tx := gormDb.Begin()
	checkErr(tx.Error)
	//if delete user fail then rollback
	if err := tx.Delete(u).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

//find user by name
func (*UserService) FindUser(username string) *User {
	users := make([]User, 0, 0)
	checkErr(gormDb.Where([]string{username}).Find(&users).Error)
	if len(users) == 0 {
		return nil
	}
	return &users[0]
}

//find all users
func (*UserService) Allusers() []User {
	users := make([]User, 0, 0)
	checkErr(gormDb.Find(&users).Error)
	return users
}

//current user type
type CurrentUser struct {
	Username string `gorm:"primary_key;column:username"`
}

//table name of current user table
func (*CurrentUser) TableName() string {
	return "currentuser"
}

//current user service
type CurrentUserService struct{}

//current serv
var CurrentServ = CurrentUserService{}

func (*CurrentUserService) load() {
	c := &CurrentUser{}
	if !gormDb.HasTable(c) {
		gormDb.CreateTable(c)
	}
}

//add current user = login
func (*CurrentUserService) Add(c *CurrentUser) {
	tx := gormDb.Begin()
	checkErr(tx.Error)
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}
	tx.Commit()
}

//delete current user = logout
func (*CurrentUserService) Delete(username string) {
	c := &CurrentUser{Username: username}
	tx := gormDb.Begin()
	checkErr(tx.Error)
	if err := tx.Delete(c).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}
	tx.Commit()
}

//if someone login already
func (*CurrentUserService) WhoLogin() *CurrentUser {
	user := make([]CurrentUser, 0, 0)
	checkErr(gormDb.Find(&user).Error)
	if len(user) == 0 {
		return nil
	}
	return &user[0]
}

//Meeting type
type Meeting struct {
	Title             string                `json:"title" gorm:"primary_key;column:title"`
	Sponsor           User                  `gorm:"column:sponsor;ForeignKey:Username"`
	SponsorName       string                `json:"sponsorname" gorm:"-"`
	ParticipatorsName []string              `json:"participators" gorm:"-"`
	Participators     []MeetingParticipator `gorm:"ForeignKey:Username"`
	StartTime         string                `json:"startTime" gorm:"column:startTime"`
	EndTime           string                `json:"endTime" gorm:"column:endTime"`
}

func (*Meeting) TableName() string {
	return "meeting"
}

type MeetingService struct{}

var MeetingServ = MeetingService{}

func (*MeetingService) load() {
	meet := &Meeting{}
	if !gormDb.HasTable(meet) {
		gormDb.CreateTable(meet)
	}
}
func (*MeetingService) Add(meet *Meeting) {
	tx := gormDb.Begin()
	checkErr(tx.Error)
	if err := tx.Create(meet).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}
	tx.Commit()
}
func (*MeetingService) Find() []Meeting {
	all_meetings := make([]Meeting, 0, 0)
	checkErr(gormDb.Find(&all_meetings).Error)
	return all_meetings
}

//MeetingParticipators combine meeting and participators
type MeetingParticipator struct {
	gorm.Model
	Title    string `json:"title" gorm:"column:title"`
	Username string `json:"username" gorm:"column:username"`
}

func (*MeetingParticipator) TableName() string {
	return "participator"
}

type MeetingParticipatorService struct{}

var MeetingParticipatorServ = MeetingParticipatorService{}

func (*MeetingParticipatorService) load() {
	m := &MeetingParticipator{}
	if !gormDb.HasTable(m) {
		gormDb.CreateTable(m)
	}
}
func (*MeetingParticipatorService) Add(m *Meeting) {
	tx := gormDb.Begin()
	checkErr(tx.Error)
	for _, par := range m.ParticipatorsName {
		meetingpar := &MeetingParticipator{
			Title:    m.Title,
			Username: par,
		}
		if err := tx.Create(meetingpar).Error; err != nil {
			tx.Rollback()
			checkErr(err)
		}
	}
	tx.Commit()
}
func init() {
	addServ(&UserServ)
	addServ(&CurrentServ)
	addServ(&MeetingServ)
	addServ(&MeetingParticipatorServ)
}
