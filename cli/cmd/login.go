package cmd

import (
	"fmt"

	"github.com/YlingMA/Agenda-Service2/cli/service"
	"github.com/spf13/cobra"
)

var str1 string
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "for user to log in",
	Long:  `the usage is to log in`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		err := service.Login(name, password)
		if err != nil {
			str1 = fmt.Sprint(err)
			fmt.Println(err)
		} else {
			fmt.Println("Login successfully!")
			str1 = "Login successfully!"
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("name", "n", "", "user_name")
	loginCmd.Flags().StringP("password", "p", "", "password")
}
