// Maintainer: Lucian Maly <lmaly@redhat.com>
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/luckylittle/mdblist-cli/internal/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	apiClient *client.Client
	output    string
)

var rootCmd = &cobra.Command{
	Use:   "mdblist-cli",
	Short: "A CLI for interacting with the MDBList API",
	Long:  `A command-line interface to perform various actions against the MDBList RESTful API.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("api_key")
		var err error
		apiClient, err = client.New(apiKey)
		if err != nil {
			return fmt.Errorf("failed to initialize API client: %w", err)
		}
		return nil
	},
}

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

	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "json", "Output format (json, yaml)")
}

func printJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(b))
}

func printData(data interface{}) {
	switch output {
	case "json":
		printJSON(data)
	case "yaml":
		printYAML(data)
	default:
		fmt.Printf("Error: unknown output format %q\n", output)
	}
}

func printYAML(data interface{}) {
	b, err := yaml.Marshal(data)
	if err != nil {
		fmt.Println("Error formatting YAML:", err)
		return
	}
	fmt.Println(string(b))
}
