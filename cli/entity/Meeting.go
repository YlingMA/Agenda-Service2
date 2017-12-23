package entity

import (
	"fmt"
	//"net/http"
)

type Meeting struct {
	Sponsor       string   `json:"sponsor"`
	Participators []string `json:"participators"`
	StartTime     string   `json:"startTime"`
	EndTime       string   `json:"endTime"`
	Title         string   `json:"title"`
}

// for create meeting
func CreateMeeting(meeting *Meeting) (err error) {
	var code int
	var requestBody struct {
		Title         string   `json:"title"`
		Participators []string `json:"participators"`
		StartTime     string   `json:"startTime"`
		EndTime       string   `json:"endTime"`
	}
	requestBody.Title = meeting.Title
	requestBody.Participators = meeting.Participators
	requestBody.StartTime = meeting.StartTime
	requestBody.EndTime = meeting.EndTime
	var responseBody struct {
		Message string `json:"message"`
	}
	code, err = request("POST", "/v1/meetings", &requestBody, &responseBody)
	if code == 201 {
		return nil
	}
	if err != nil {
		return err
	}
	err = fmt.Errorf("%s", responseBody.Message)
	return err
}

// for qurey meeting
/*func QueryMeeting(startTime string, endTime string) (responseBody []Meeting, err error) {
	var code int
	code, err = request("GET", "/v1/meetings"+"?startTime="+startTime+"&endTime="+endTime, nil, &responseBody)
	if err != nil {
		return nil, err
	}
	// 200
	if code == http.StatusOK {
		return responseBody, nil
	}
	// 401
	if code == http.StatusUnauthorized {
		return nil, fmt.Errorf("Please log in !")
	}
	return
}*/
