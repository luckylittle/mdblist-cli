package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources from MDBList",
}

var getMyLimitsCmd = &cobra.Command{
	Use:   "my-limits",
	Short: "Get information about your API limits",
	Run: func(cmd *cobra.Command, args []string) {
		limits, err := apiClient.GetMyLimits()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printJSON(limits)
	},
}

var getMyListsCmd = &cobra.Command{
	Use:   "my-lists",
	Short: "Fetch your lists",
	Run: func(cmd *cobra.Command, args []string) {
		lists, err := apiClient.GetMyLists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printJSON(lists)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getMyLimitsCmd)
	getCmd.AddCommand(getMyListsCmd)
}
