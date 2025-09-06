package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/luckylittle/mdblist-cli/internal/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiClient *client.Client
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mdblist-cli",
	Short: "A CLI for interacting with the MDBList API",
	Long:  `A command-line interface to perform various actions against the MDBList RESTful API.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// This function runs before any subcommand
		apiKey := viper.GetString("api_key")
		var err error
		apiClient, err = client.New(apiKey)
		if err != nil {
			return fmt.Errorf("failed to initialize API client: %w", err)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Read MDBLIST_API_KEY from environment variable
	viper.SetEnvPrefix("mdblist")
	viper.BindEnv("api_key")
}

// Helper to print API response as pretty JSON
func printJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(b))
}
