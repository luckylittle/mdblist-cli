// Maintainer: Lucian Maly <lmaly@redhat.com>
package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search resources in MDBList.",
}

var searchMediaCmd = &cobra.Command{
	Use:   "media <media-type>",
	Short: "Search for movie, show or both (any).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mediaType := args[0]
		query, _ := cmd.Flags().GetString("query")

		if query == "" {
			fmt.Println("Error: --query flag is required!")
			return
		}

		params := url.Values{}
		params.Set("query", query)

		result, err := apiClient.SearchMedia(mediaType, params)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(result)
	},
}

var searchListsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Search public lists by title.",
	Run: func(cmd *cobra.Command, args []string) {
		query, _ := cmd.Flags().GetString("query")
		if query == "" {
			fmt.Println("Error: --query flag is required!")
			return
		}

		params := url.Values{}
		params.Set("query", query)

		lists, err := apiClient.SearchLists(params)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		printData(lists)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchMediaCmd)
	searchCmd.AddCommand(searchListsCmd)

	searchMediaCmd.Flags().StringP("query", "q", "", "Search query")
	searchListsCmd.Flags().StringP("query", "q", "", "Search query")

}
