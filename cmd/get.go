// Maintainer: Lucian Maly <lmaly@redhat.com>
package cmd

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources from MDBList.",
}

var getMyLimitsCmd = &cobra.Command{
	Use:   "my-limits",
	Short: "Show information about user limits.",
	Run: func(cmd *cobra.Command, args []string) {
		limits, err := apiClient.GetMyLimits()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(limits)
	},
}

var getMyListsCmd = &cobra.Command{
	Use:   "my-lists",
	Short: "Fetches users lists.",
	Run: func(cmd *cobra.Command, args []string) {
		lists, err := apiClient.GetMyLists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(lists)
	},
}

var getUserListsCmd = &cobra.Command{
	Use:   "user-lists",
	Short: "Fetch a user's lists.",
	Run: func(cmd *cobra.Command, args []string) {
		userID, _ := cmd.Flags().GetInt("id")
		username, _ := cmd.Flags().GetString("username")

		if userID == 0 && username == "" {
			fmt.Println("Error: either user's --id or --username is required!")
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
		printData(lists)
	},
}

var getListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves details of a list.",
	Run: func(cmd *cobra.Command, args []string) {
		listID, _ := cmd.Flags().GetInt("id")
		username, _ := cmd.Flags().GetString("username")
		listName, _ := cmd.Flags().GetString("listname")

		if listID == 0 && (username == "" || listName == "") {
			fmt.Println("Error: either --id or both --username and --listname are required!")
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
		printData(list)
	},
}

var getListItemsCmd = &cobra.Command{
	Use:   "list-items",
	Short: "Fetches items from a specified list.",
	Run: func(cmd *cobra.Command, args []string) {
		listID, _ := cmd.Flags().GetInt("id")
		username, _ := cmd.Flags().GetString("username")
		listName, _ := cmd.Flags().GetString("listname")

		if listID == 0 && (username == "" || listName == "") {
			fmt.Println("Error: either --id or both --username and --listname are required!")
			return
		}

		var (
			items interface{}
			err   error
		)

		// TODO: extend this to take CLI flags for pagination, sorting, etc.
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
		printData(items)
	},
}

var getListChangesCmd = &cobra.Command{
	Use:   "list-changes <list-id>",
	Short: "Returns Trakt IDs for items changed after the last list update.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid list ID provided!")
			return
		}

		changes, err := apiClient.GetListChanges(listID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(changes)
	},
}

var getMediaInfoCmd = &cobra.Command{
	Use:   "media-info <provider> <media-type> <media-id>",
	Short: "Fetch information about a media item",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		provider, mediaType, mediaID := args[0], args[1], args[2]

		// Example of how you could add optional params via flags
		// For now, it's empty.
		params := url.Values{}

		info, err := apiClient.GetMediaInfo(provider, mediaType, mediaID, params)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(info)
	},
}

var getTopListsCmd = &cobra.Command{
	Use:   "top-lists",
	Short: "Outputs the top lists sorted by Trakt likes.",
	Run: func(cmd *cobra.Command, args []string) {
		lists, err := apiClient.GetTopLists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(lists)
	},
}

var getLastActivitiesCmd = &cobra.Command{
	Use:   "last-activities",
	Short: "Fetch the last activity timestamps for sync.",
	Run: func(cmd *cobra.Command, args []string) {
		activities, err := apiClient.GetLastActivities()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(activities)
	},
}

var getWatchlistItemsCmd = &cobra.Command{
	Use:   "watchlist-items",
	Short: "Fetches watchlist items, they are sorted by date added.",
	Run: func(cmd *cobra.Command, args []string) {
		params := url.Values{}
		sort, _ := cmd.Flags().GetString("sort")
		if sort != "" {
			params.Set("sort", sort)
		}

		items, err := apiClient.GetWatchlistItems(params)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(items)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getMyLimitsCmd)
	getCmd.AddCommand(getMyListsCmd)
	getCmd.AddCommand(getUserListsCmd)
	getCmd.AddCommand(getListCmd)
	getCmd.AddCommand(getListItemsCmd)
	getCmd.AddCommand(getListChangesCmd)
	getCmd.AddCommand(getMediaInfoCmd)
	getCmd.AddCommand(getTopListsCmd)
	getCmd.AddCommand(getLastActivitiesCmd)
	getCmd.AddCommand(getWatchlistItemsCmd)

	getUserListsCmd.Flags().Int("id", 0, "User ID")
	getUserListsCmd.Flags().String("username", "", "Username")

	getListCmd.Flags().Int("id", 0, "List ID")
	getListCmd.Flags().String("username", "", "Username of the list owner")
	getListCmd.Flags().String("listname", "", "Name/slug of the list")

	getListItemsCmd.Flags().Int("id", 0, "List ID")
	getListItemsCmd.Flags().String("username", "", "Username of the list owner")
	getListItemsCmd.Flags().String("listname", "", "Name/slug of the list")

	getWatchlistItemsCmd.Flags().String("sort", "", "Sort order (e.g., 'added_at.desc')")
}
