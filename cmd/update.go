// Maintainer: Lucian Maly <lmaly@redhat.com>
package cmd

import (
	"fmt"

	"github.com/luckylittle/mdblist-cli/internal/client"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update resources on MDBList",
}

var updateListNameCmd = &cobra.Command{
	Use:   "list-name <new-name>",
	Short: "Update a list's name by ID or by username and list name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newName := args[0]
		listID, _ := cmd.Flags().GetInt("id")
		username, _ := cmd.Flags().GetString("username")
		listName, _ := cmd.Flags().GetString("listname")

		if listID == 0 && (username == "" || listName == "") {
			fmt.Println("Error: either --id or both --username and --listname are required")
			return
		}

		var (
			response *client.ListUpdateResponse
			err      error
		)

		if listID != 0 {
			response, err = apiClient.UpdateListNameByID(listID, newName)
		} else {
			response, err = apiClient.UpdateListNameByName(username, listName, newName)
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("List name updated successfully.")
		printJSON(response)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateListNameCmd)

	updateListNameCmd.Flags().Int("id", 0, "List ID")
	updateListNameCmd.Flags().String("username", "", "Username of the list owner")
	updateListNameCmd.Flags().String("listname", "", "Current name/slug of the list")
}
