package main

import (
	"log"
	"github.com/spf13/cobra"
)

// Command for 'fetch' action
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch certsuite usage from Quay and DCI",
	Run: func(cmd *cobra.Command, args []string) {
		err := FetchCertsuiteUsage()
		if err != nil {
			log.Fatalf("Failed to fetch certsuite usage: %v", err)
		}
	},
}

// Root command
var rootCmd = &cobra.Command{
	Use:   "certsuite-overview",
	Short: "A CLI to interact with certsuite data",
}

func init() {
	// Define flags for the certsuite-overview command
	rootCmd.PersistentFlags().String("DB_USER", "", "Database user")
	rootCmd.PersistentFlags().String("DB_PASSWORD", "", "Database password")
	rootCmd.PersistentFlags().String("DB_URL", "", "Database URL")
	rootCmd.PersistentFlags().String("DB_PORT", "", "Database port")
	rootCmd.PersistentFlags().String("DB_NAME", "", "Database name")
	rootCmd.PersistentFlags().String("CLIENTID", "", "Client ID")
	rootCmd.PersistentFlags().String("APISECRET", "", "API Secret")
	rootCmd.PersistentFlags().String("BEARERTOKEN", "", "Bearer Token")
	rootCmd.PersistentFlags().String("NAMESPACE", "", "Namespace")
	rootCmd.PersistentFlags().String("REPOSITORY", "", "Repository")
}

func main() {
	// Add fetchCmd to root command
	rootCmd.AddCommand(fetchCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
