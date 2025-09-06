package cmd

import (
	"fmt"
	"net/url"

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

var getUserListsCmd = &cobra.Command{
	Use:   "user-lists",
	Short: "Fetch a user's lists by ID or username",
	Run: func(cmd *cobra.Command, args []string) {
		userID, _ := cmd.Flags().GetInt("id")
		username, _ := cmd.Flags().GetString("username")

		if userID == 0 && username == "" {
			fmt.Println("Error: either --id or --username is required")
			return
		}

		var (
			lists interface{}
			err   error
		)

		if userID != 0 {
			lists, err = apiClient.GetUserListsByID(userID)
		} else {
			lists, err = apiClient.GetUserListsByName(username)
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printJSON(lists)
	},
}

var getListCmd = &cobra.Command{
	Use:   "list",
	Short: "Fetch a specific list by ID or by username and list name",
	Run: func(cmd *cobra.Command, args []string) {
		listID, _ := cmd.Flags().GetInt("id")
		username, _ := cmd.Flags().GetString("username")
		listName, _ := cmd.Flags().GetString("listname")

		if listID == 0 && (username == "" || listName == "") {
			fmt.Println("Error: either --id or both --username and --listname are required")
			return
		}

		var (
			list interface{}
			err  error
		)

		if listID != 0 {
			list, err = apiClient.GetListByID(listID)
		} else {
			list, err = apiClient.GetListByName(username, listName)
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printJSON(list)
	},
}

var getListItemsCmd = &cobra.Command{
	Use:   "list-items",
	Short: "Fetch items from a list by ID or by username and list name",
	Run: func(cmd *cobra.Command, args []string) {
		listID, _ := cmd.Flags().GetInt("id")
		username, _ := cmd.Flags().GetString("username")
		listName, _ := cmd.Flags().GetString("listname")

		if listID == 0 && (username == "" || listName == "") {
			fmt.Println("Error: either --id or both --username and --listname are required")
			return
		}

		var (
			items interface{}
			err   error
		)

		// For this command, we'll just use an empty params for now.
		// You could extend this to take CLI flags for pagination, sorting, etc.
		params := url.Values{}

		if listID != 0 {
			items, err = apiClient.GetListItems(listID, params)
		} else {
			items, err = apiClient.GetListItemsByName(username, listName, params)
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printJSON(items)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getMyLimitsCmd)
	getCmd.AddCommand(getMyListsCmd)
	getCmd.AddCommand(getUserListsCmd)
	getCmd.AddCommand(getListCmd)
	getCmd.AddCommand(getListItemsCmd)

	getUserListsCmd.Flags().Int("id", 0, "User ID")
	getUserListsCmd.Flags().String("username", "", "Username")

	getListCmd.Flags().Int("id", 0, "List ID")
	getListCmd.Flags().String("username", "", "Username of the list owner")
	getListCmd.Flags().String("listname", "", "Name/slug of the list")

	getListItemsCmd.Flags().Int("id", 0, "List ID")
	getListItemsCmd.Flags().String("username", "", "Username of the list owner")
	getListItemsCmd.Flags().String("listname", "", "Name/slug of the list")
}
