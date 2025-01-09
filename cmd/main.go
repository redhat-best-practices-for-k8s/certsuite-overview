package main

import (
	"log"

	"github.com/redhat-best-practices-for-k8s/certsuite-overview/config"
	"github.com/spf13/cobra"
)

// Command for 'fetch' action
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch certsuite usage from Quay and DCI",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		// Get values from flags
		config.AppConfig.DBUser, err = cmd.Flags().GetString("DB_USER")
		if err != nil {
			log.Fatalf("Error fetching DB_USER: %v", err)
		}
		config.AppConfig.DBPassword, err = cmd.Flags().GetString("DB_PASSWORD")
		if err != nil {
			log.Fatalf("Error fetching DB_PASSWORD: %v", err)
		}
		config.AppConfig.DBURL, err = cmd.Flags().GetString("DB_URL")
		if err != nil {
			log.Fatalf("Error fetching DB_URL: %v", err)
		}
		config.AppConfig.DBPort, err = cmd.Flags().GetString("DB_PORT")
		if err != nil {
			log.Fatalf("Error fetching DB_PORT: %v", err)
		}
		config.AppConfig.DBName, err = cmd.Flags().GetString("DB_NAME")
		if err != nil {
			log.Fatalf("Error fetching DB_NAME: %v", err)
		}
		config.AppConfig.ClientID, err = cmd.Flags().GetString("CLIENTID")
		if err != nil {
			log.Fatalf("Error fetching CLIENTID: %v", err)
		}
		config.AppConfig.APISecret, err = cmd.Flags().GetString("APISECRET")
		if err != nil {
			log.Fatalf("Error fetching APISECRET: %v", err)
		}
		config.AppConfig.BearerToken, err = cmd.Flags().GetString("BEARERTOKEN")
		if err != nil {
			log.Fatalf("Error fetching BEARERTOKEN: %v", err)
		}
		config.AppConfig.Namespace, err = cmd.Flags().GetString("NAMESPACE")
		if err != nil {
			log.Fatalf("Error fetching NAMESPACE: %v", err)
		}
		config.AppConfig.Repository, err = cmd.Flags().GetString("REPOSITORY")
		if err != nil {
			log.Fatalf("Error fetching REPOSITORY: %v", err)
		}
		// Call the function to fetch certsuite usage
		err = FetchCertsuiteUsage()
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
