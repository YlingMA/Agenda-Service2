package service

import (
	"fmt"
	//"time"

	"github.com/Caroline1997/Service-Agenda/cli/entity"
)

// log's command

// Login POST /v1/user/login
func Login(name string, password string) (err error) {
	var flag bool
	flag, err = entity.Check_Login()
	if err != nil {
		return err
	}
	if flag == true {
		var currentUser string
		currentUser, err = entity.GetCurrentUser()
		err = fmt.Errorf("There exist someone logged as '%s', please logout!", currentUser)
		return err
	}
	err = entity.Login(name, password)
	if err != nil {
		return err
	}
	return nil
}

// Logout POST /v1/user/logout
func Logout() (err error) {
	var flag bool
	flag, err = entity.Check_Login()
	if err != nil {
		return err
	}
	if flag == true {
		err, _ = entity.Logout()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Please log in first")
	}
	return nil
}

// User's command

func check_user_valid(user *entity.User) error {
	if len(user.Username) == 0 {
		return fmt.Errorf("Please enter your username!")
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("Please enter your password!")
	}
	return nil
}

// Register POST /v1/users
func Register(username string, password string, email string, phone string) (err error) {
	User1 := &entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	err = check_user_valid(User1)
	if err != nil {
		return err
	}
	err = entity.CreateUser(User1)
	if err != nil {
		return err
	}
	return nil
}

// DELETE DELETE /v1/user/account
func Delete_user(username string, password string) (err error) {
	var flag bool
	flag, err = entity.Check_Login()
	if err != nil {
		return err
	}
	if flag == true {
		err = entity.DeleteUser()
		if err != nil {
			return err
		}
	}
	return nil
}

// ListAllUsers GET /v1/users
func List_all_users() ([]entity.User, error) {
	return entity.ListAllUsers()
}

// Meeting's command

func check_meeting_valid(meeting *entity.Meeting) error {
	if len(meeting.Title) == 0 {
		return fmt.Errorf("Please enter your title!")
	}
	if len(meeting.Participators) == 0 {
		return fmt.Errorf("Please enter at least one participator!")
	}
	return nil
}

// Create Meeting POST /v1/meetings
func Create_meeting(title string, participators []string, startTime string, endTime string) (err error) {
	var flag bool
	flag, err = entity.Check_Login()
	if err != nil {
		return err
	}
	if flag == true {
		sponsor, err := entity.GetCurrentUser()
		if err != nil {
			return err
		}
		Meeting1 := &entity.Meeting{
			Sponsor:       sponsor,
			Participators: participators,
			StartTime:     startTime,
			EndTime:       endTime,
			Title:         title,
		}
		err = check_meeting_valid(Meeting1)
		if err != nil {
			return err
		}
		err = entity.CreateMeeting(Meeting1)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// Query Meetings GET /v1/meetings{?startDate,endDate}
/*func Query_meeting(startTime string, endTime string) (meetings []entity.Meeting, err error) {
	var flag bool
	flag, err = entity.Check_Login()
	if err != nil {
		return nil, err
	}
	if flag == true {
		// check start time
		if len(startTime) == 0 {
			return nil, fmt.Errorf("Please enter startTime!")
		}
		// 2006-01-02 15:04:05 is time format
		_, err := time.Parse("2006-01-02 15:04:05", startTime)
		if err != nil {
			return nil, fmt.Errorf("Please enter correct startTime format!")
		}
		// check end time
		if len(endTime) == 0 {
			return nil, fmt.Errorf("Please enter endTime!")
		}
		_, err = time.Parse("2006-01-02 15:04:05", endTime)
		if err != nil {
			return nil, fmt.Errorf("Please enter correct endTime format!")
		}
		// compare
		if startTime > endTime {
			return nil, fmt.Errorf("startTime should not larger than endTime!")
		}
		meetings, err = entity.QueryMeeting(startTime, endTime)
		if err != nil {
			return nil, err
		}
		return meetings, nil
	}
	return nil, nil
}*/
