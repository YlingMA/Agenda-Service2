package cmd

import (
	"fmt"

	"github.com/YlingMA/Agenda-Service2/cli/service"
	"github.com/spf13/cobra"
)

var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting -t [Title] -p [Participator] -s [StartTime] -e [EndTime]",
	Short: "create a new meeting",
	Long:  `usage of using this command is to create a new meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("Title")
		participators, _ := cmd.Flags().GetStringSlice("Participators")
		startTime, _ := cmd.Flags().GetString("StartTime")
		endTime, _ := cmd.Flags().GetString("EndTime")
		err := service.Create_meeting(title, participators, startTime, endTime)
		if err != nil {
			fmt.Println(err)
			str1 = fmt.Sprint(err)
		} else {
					fmt.Println("Create meeting successfully!")
					str1 = "Create meeting successfully!"
		}
	},
}

func init() {
	RootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringP("Title", "t", "", "title of the meeting")
	createMeetingCmd.Flags().StringSliceP("Participators", "p", []string{}, "participators of the meeting")
	createMeetingCmd.Flags().StringP("StartTime", "s", "", "startTime of the meeting")
	createMeetingCmd.Flags().StringP("EndTime", "e", "", "endTime of the meeting")
}
