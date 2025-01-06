package main

import (
	"fmt"
	"log"
	"os"

	"github.com/redhat-best-practices-for-k8s/certsuite-overview/pkg"
)

// FetchCertsuiteUsage integrates data from Quay and DCI.
func FetchCertsuiteUsage() error {
	bearerToken := os.Getenv("BEARERTOKEN")
	log.Printf("BEARERTOKEN is set to: %s", bearerToken)
	if err := pkg.FetchQuayData(); err != nil {
		return fmt.Errorf("error fetching Quay data: %w", err)
	}
	if err := pkg.FetchDciData(); err != nil {
		return fmt.Errorf("error fetching DCI data: %w", err)
	}
	return nil
}
