package cmd

import (
	"fmt"
	"testing"
)

//正常注册
func TestRegister1(t *testing.T) {
	fmt.Println("-------正常*注册*操作测试--------")
	//fmt.Println("Register")
	registerCmd.Flags().Set("user", "root")
	registerCmd.Flags().Set("password", "pass")
	registerCmd.Flags().Set("mail", "mail")
	registerCmd.Flags().Set("phone", "phone")
	registerCmd.Run(registerCmd, nil)
	if (str1 != "Register successfully!") {
		t.Errorf("Register should be successful but failed because %s!",str1)
	}
}
//重复注册
func TestRegister2(t *testing.T) {
	fmt.Println("-------重复*注册*操作测试--------")
	//fmt.Println("Register")
	registerCmd.Flags().Set("user", "myl")
	registerCmd.Flags().Set("password", "pass")
	registerCmd.Flags().Set("mail", "mail")
	registerCmd.Flags().Set("phone", "phone")
	registerCmd.Run(registerCmd, nil)
	if (str1 != "the input username have been used, please try another one") {
		t.Errorf("Register should be failed but %s!",str1)
	}
}

//正常登出
func TestLogout1(t *testing.T) {
	fmt.Println("-------正常*登出*操作测试--------")
	//fmt.Println("Logout")
	logoutCmd.Run(logoutCmd, nil)
	if (str1 != "Logout successfully!") {
		t.Errorf("Logout should be failed not because of %s!",str1)
	}
}
//正常登录
func TestLogin1(t *testing.T) {
	fmt.Println("-------正常*登录*操作测试1--------")
	//fmt.Println("Login")
	loginCmd.Flags().Set("name", "root")
	loginCmd.Flags().Set("password", "pass")
	loginCmd.Run(loginCmd, nil)
	if (fmt.Sprint(str1) != "Login successfully!") {
		t.Errorf("Login should be successful,but failed because %s!",str1)
	}
}
//删除用户操作测试
func TestDelete1(t *testing.T) {
	fmt.Println("-------正常*删除用户*操作测试--------")
	//fmt.Println("DeleteUser")
	deleteCmd.Flags().Set("name", "root")
	deleteCmd.Flags().Set("password", "pass")
	deleteCmd.Run(deleteCmd, nil)
	if (fmt.Sprint(str1) != "Delete user successfully!") {
		t.Errorf("Delete user should be successful,but failed because %s!",str1)
	}
}
//正常创建会议
func TestCreateMeeting1(t *testing.T) {
	fmt.Println("-------正常*创建会议*操作测试--------")
	//fmt.Println("Register")
	createMeetingCmd.Flags().Set("Title", "titleofmeeting")
	createMeetingCmd.Flags().Set("Participators", "myl")
	createMeetingCmd.Flags().Set("StartTime", "2017-12-20/17:10")
	createMeetingCmd.Flags().Set("EndTime", "2017-12-22/23:01")
	createMeetingCmd.Run(createMeetingCmd, nil)

	if (str1 != "Create meeting successfully!" ) {
		t.Errorf("Create meeting should be successful but failed because %s!",str1)
	}
}
//异常登录
func TestLogin2(t *testing.T) {
	fmt.Println("-------错误账号密码*登录*操作测试--------")
	//fmt.Println("Login")
	loginCmd.Flags().Set("name", "root")
	loginCmd.Flags().Set("password", "pass")
	loginCmd.Run(loginCmd, nil)
	if (str1 != "the username and password are not correct" ) {
		t.Errorf("Login should be failed because username and password are not correct,but %s!",str1)
	}
}
//未登录就登出
func TestLogout2(t *testing.T) {
	fmt.Println("-------登陆前先*登出*操作测试--------")
	//fmt.Println("Logout")
	logoutCmd.Run(logoutCmd, nil)
	if (str1 != "Please log in first" ) {
		t.Errorf("Logout should be failed because you didn't log in first,not because %s!",str1)
	}
}
