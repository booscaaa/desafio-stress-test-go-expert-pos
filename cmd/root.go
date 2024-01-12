/*
Copyright © 2024 VINÍCIUS BOSCARDIN boscardinvinicius@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var totalRequests int = 0
var report = map[string]map[string]any{}
var reportMutex = &sync.Mutex{}

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Stress test application",
	Long:  `Performs load testing on a web service. The user provides the URL of the service, the total number of requests, and the number of simultaneous calls.`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime := time.Now()
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt32("requests")
		concurrency, _ := cmd.Flags().GetInt32("concurrency")

		if url == "" {
			log.Fatal("The URL is empty")
		}

		var wg sync.WaitGroup

		for i := int32(0); i < concurrency; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := int32(0); j < requests/concurrency; j++ {

					makeRequest(url)
				}
			}()
		}

		wg.Wait()

		makeReport(startTime)
	},
}

func makeReport(startTime time.Time) {
	fmt.Printf("Test completed\n")
	fmt.Printf("----------------------------------------------------------------------\n")
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("----------------------------------------------------------------------\n")

	if _, ok := report["200"]; ok {
		fmt.Printf("Successful requests: status code 200; total %d\n", report["200"]["total"])
		fmt.Printf("----------------------------------------------------------------------\n")
	}

	for key, value := range report {
		if key != "200" {
			fmt.Printf("Requests with error: status code %s; total %d\n", key, value["total"])
			fmt.Printf("----------------------------------------------------------------------\n")
		}
	}

	durationTime := time.Since(startTime)

	fmt.Printf("Total execution time: %s\n", durationTime)
}

func makeRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		addToReport("0", 1)
		return
	}
	defer resp.Body.Close()

	addToReport(strconv.Itoa(resp.StatusCode), 1)
}

func addToReport(statusCode string, count int) {
	reportMutex.Lock()
	defer reportMutex.Unlock()
	totalRequests++
	if _, ok := report[statusCode]; ok {
		report[statusCode]["total"] = report[statusCode]["total"].(int) + count
	} else {
		report[statusCode] = map[string]any{"total": count}
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL of the service to be tested.")
	rootCmd.Flags().Int32P("requests", "r", 10, "Total number of requests.")
	rootCmd.Flags().Int32P("concurrency", "c", 2, "Number of simultaneous calls.")
}
