// Maintainer: Lucian Maly <lmaly@redhat.com>
package cmd

import (
	"errors"
	"fmt"

	"github.com/luckylittle/mdblist-cli/internal/client"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update resources in MDBList.",
}

var updateListNameCmd = &cobra.Command{
	Use:   "list-name <new-name>",
	Short: "Updates the name of a list.",
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

var updateListItemsCmd = &cobra.Command{
	Use:   "list-items",
	Short: "You can modify static list by adding or removing items.",
	RunE: func(cmd *cobra.Command, args []string) error {
		listID, _ := cmd.Flags().GetInt("id")
		if listID == 0 {
			return errors.New("--id is required")
		}

		action, _ := cmd.Flags().GetString("action")
		if action != "add" && action != "remove" {
			return errors.New("--action must be either 'add' or 'remove'")
		}

		movieTmdbIDs, _ := cmd.Flags().GetIntSlice("movie-tmdb")
		movieImdbIDs, _ := cmd.Flags().GetStringSlice("movie-imdb")
		showTmdbIDs, _ := cmd.Flags().GetIntSlice("show-tmdb")
		showImdbIDs, _ := cmd.Flags().GetStringSlice("show-imdb")

		if len(movieTmdbIDs) == 0 && len(movieImdbIDs) == 0 && len(showTmdbIDs) == 0 && len(showImdbIDs) == 0 {
			return errors.New("at least one movie or show ID must be provided")
		}

		items := client.ModifyListRequest{}
		for _, id := range movieTmdbIDs {
			items.Movies = append(items.Movies, map[string]interface{}{"tmdb": id})
		}
		for _, id := range movieImdbIDs {
			items.Movies = append(items.Movies, map[string]interface{}{"imdb": id})
		}
		for _, id := range showTmdbIDs {
			items.Shows = append(items.Shows, map[string]interface{}{"tmdb": id})
		}
		for _, id := range showImdbIDs {
			items.Shows = append(items.Shows, map[string]interface{}{"imdb": id})
		}

		response, err := apiClient.ModifyListItems(listID, action, items)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmt.Printf("List items updated successfully (action: %s).\n", action)
		printData(response)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateListNameCmd)
	updateCmd.AddCommand(updateListItemsCmd)

	updateListNameCmd.Flags().Int("id", 0, "List ID")
	updateListNameCmd.Flags().String("username", "", "Username of the list owner")
	updateListNameCmd.Flags().String("listname", "", "Current name/slug of the list")

	updateListItemsCmd.Flags().IntP("id", "i", 0, "List ID (required)")
	updateListItemsCmd.Flags().StringP("action", "a", "", "Action to perform: 'add' or 'remove' (required)")
	updateListItemsCmd.Flags().IntSlice("movie-tmdb", []int{}, "TMDb ID of a movie to add/remove")
	updateListItemsCmd.Flags().StringSlice("movie-imdb", []string{}, "IMDb ID of a movie to add/remove")
	updateListItemsCmd.Flags().IntSlice("show-tmdb", []int{}, "TMDb ID of a show to add/remove")
	updateListItemsCmd.Flags().StringSlice("show-imdb", []string{}, "IMDb ID of a show to add/remove")
	updateListItemsCmd.MarkFlagRequired("id")
	updateListItemsCmd.MarkFlagRequired("action")
}
