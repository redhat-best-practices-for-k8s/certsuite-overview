package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redhat-best-practices-for-k8s/certsuite-overview/config"
	"github.com/redhat-best-practices-for-k8s/certsuite-overview/pkg"
	"github.com/spf13/cobra"
)

var (
	dciComponentsMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "dci_components_metrics",
			Help: "Metrics from dci_components table",
		},
		[]string{"job_id", "commit_hash"},
	)
	aggregatedLogsMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "aggregated_logs_metrics",
			Help: "Metrics from aggregated_logs table",
		},
		[]string{"kind"},
	)
	mutex = sync.Mutex{}
)

func fetchDataFromDB(db *sql.DB) error {
	// Select the database first
	_, err := db.Exec("USE certsuite_usage_db;")
	if err != nil {
		return fmt.Errorf("failed to select database certsuite_usage_db: %w", err)
	}

	// Query the dci_components table
	rows, err := db.Query("SELECT job_id, commit_hash, totalSuccess, totalFailures, totalErrors, totalSkips FROM dci_components")
	if err != nil {
		return fmt.Errorf("failed to query dci_components: %w", err)
	}
	defer rows.Close()

	mutex.Lock()
	for rows.Next() {
		var jobID, commitHash string
		var totalSuccess, totalFailures, totalErrors, totalSkips int
		if err := rows.Scan(&jobID, &commitHash, &totalSuccess, &totalFailures, &totalErrors, &totalSkips); err != nil {
			mutex.Unlock()
			return fmt.Errorf("failed to scan row: %w", err)
		}
		dciComponentsMetrics.WithLabelValues(jobID, commitHash).Set(float64(totalSuccess))
		dciComponentsMetrics.WithLabelValues(jobID, commitHash).Set(float64(totalFailures))
		dciComponentsMetrics.WithLabelValues(jobID, commitHash).Set(float64(totalErrors))
		dciComponentsMetrics.WithLabelValues(jobID, commitHash).Set(float64(totalSkips))
	}
	mutex.Unlock()

	// Query the aggregated_logs table
	rows, err = db.Query("SELECT kind, count FROM aggregated_logs")
	if err != nil {
		return fmt.Errorf("failed to query aggregated_logs: %w", err)
	}
	defer rows.Close()

	mutex.Lock()
	for rows.Next() {
		var kind string
		var count int
		if err := rows.Scan(&kind, &count); err != nil {
			mutex.Unlock()
			return fmt.Errorf("failed to scan row: %w", err)
		}
		aggregatedLogsMetrics.WithLabelValues(kind).Set(float64(count))
	}
	mutex.Unlock()

	return nil
}

func updateMetrics(db *sql.DB) {
	for {
		if err := fetchDataFromDB(db); err != nil {
			log.Printf("Error fetching data from DB: %v", err)
		}
		log.Println("Metrics updated successfully")
		time.Sleep(10 * time.Second)
	}
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch certsuite usage from Quay and DCI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := FetchCertsuiteUsage(); err != nil {
			log.Fatalf("Failed to fetch certsuite usage: %v", err)
		}
		log.Println("Certsuite usage fetched successfully")
	},
}

var rootCmd = &cobra.Command{
	Use:   "certsuite-overview",
	Short: "A CLI to interact with certsuite data",
}

func init() {
	config.LoadConfig()
	rootCmd.AddCommand(fetchCmd)
	prometheus.MustRegister(dciComponentsMetrics)
	prometheus.MustRegister(aggregatedLogsMetrics)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	dbChoice := os.Getenv("DB_CHOICE")
	var db *sql.DB
	var err error

	if dbChoice == "aws" {
		db, _, err = pkg.ConnectToAWSDB()
	} else {
		db, err = pkg.ConnectToLocalDB()
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	go updateMetrics(db)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Serving metrics at :8080/metrics")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	select {}
}