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
		// Get values from flags
		dbUser, _ := cmd.Flags().GetString("DB_USER")
		dbPassword, _ := cmd.Flags().GetString("DB_PASSWORD")
		dbURL, _ := cmd.Flags().GetString("DB_URL")
		dbPort, _ := cmd.Flags().GetString("DB_PORT")
		dbName, _ := cmd.Flags().GetString("DB_NAME")
		clientID, _ := cmd.Flags().GetString("CLIENTID")
		apiSecret, _ := cmd.Flags().GetString("APISECRET")
		bearerToken, _ := cmd.Flags().GetString("BEARERTOKEN")
		namespace, _ := cmd.Flags().GetString("NAMESPACE")
		repository, _ := cmd.Flags().GetString("REPOSITORY")

		// Print the fetched values (for debugging purposes)
		log.Printf("DB_USER: %s", dbUser)
		log.Printf("DB_PASSWORD: %s", dbPassword)
		log.Printf("DB_URL: %s", dbURL)
		log.Printf("DB_PORT: %s", dbPort)
		log.Printf("DB_NAME: %s", dbName)
		log.Printf("CLIENTID: %s", clientID)
		log.Printf("APISECRET: %s", apiSecret)
		log.Printf("BEARERTOKEN: %s", bearerToken)
		log.Printf("NAMESPACE: %s", namespace)
		log.Printf("REPOSITORY: %s", repository)

		// Call the function to fetch certsuite usage
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
