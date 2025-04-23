package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeSoft/traffik-one/internal/core/dto"
	"github.com/spf13/cobra"
)

var ipv4 string
var port int
var protocol string
var weight int
var hostname string
var poolID string

var backendCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new backend",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating backend with:")
		fmt.Printf("- IPv4: %s\n", ipv4)
		fmt.Printf("- Port: %d\n", port)
		fmt.Printf("- Protocol: %s\n", protocol)
		fmt.Printf("- Weight: %d\n", weight)
		fmt.Printf("- Hostname: %s\n", hostname)
		fmt.Printf("- PoolID: %s\n", poolID)

		if ipv4 == "" || port == 0 || protocol == "" || poolID == "" {
			log.Print("Missing required parameters")
			return
		}

		if hostname == "" {
			hostname = "none"
		}

		backend := dto.AddBackendInput{
			IPv4:     ipv4,
			Port:     port,
			Protocol: protocol,
			Weight:   weight,
			Hostname: hostname,
			PoolID:   poolID,
		}

		jsonData, err := json.Marshal(backend)
		if err != nil {
			log.Printf("Failed to marshal JSON: %v", err)
			return
		}

		createBackendURL := fmt.Sprintf("http://%s:%s/backends/add", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))
		req, err := http.NewRequest("POST", createBackendURL, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Request failed: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Request failed with status: %v", resp.Status)
			return
		}

		log.Println("Backend successfully created.")
	},
}

var BackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Manage backends for TraffikOne",
}

func init() {
	backendCreateCmd.Flags().StringVar(&ipv4, "ipv4", "", "Backend IPv4 address")
	backendCreateCmd.Flags().IntVar(&port, "port", 0, "Backend port")
	backendCreateCmd.Flags().StringVar(&protocol, "protocol", "", "Protocol (http, https, etc)")
	backendCreateCmd.Flags().IntVar(&weight, "weight", 1, "Load balancing weight")
	backendCreateCmd.Flags().StringVar(&hostname, "hostname", "", "Hostname of the backend")
	backendCreateCmd.Flags().StringVar(&poolID, "pool-id", "", "ID of the pool to assign this backend")

	_ = backendCreateCmd.MarkFlagRequired("ipv4")
	_ = backendCreateCmd.MarkFlagRequired("port")
	_ = backendCreateCmd.MarkFlagRequired("protocol")
	_ = backendCreateCmd.MarkFlagRequired("pool-id")

	BackendCmd.AddCommand(backendCreateCmd)
	RootCmd.AddCommand(BackendCmd)
}
