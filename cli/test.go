package cli

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var times int

var httpTestCmd = &cobra.Command{
	Use:   "http",
	Short: "Test backends with Load Balancer",
	Run: func(cmd *cobra.Command, args []string) {
		loadBalancerEntryPoint := fmt.Sprintf("http://%s:%s/", os.Getenv("HTTP_LOAD_BALANCER_HOST"), os.Getenv("HTTP_LOAD_BALANCER_PORT"))

		log.Printf("Sending %d HTTP requests to load balancer at %s", times, loadBalancerEntryPoint)

		client := &http.Client{}
		if times > 100 {
			log.Print("Max of 100 test HTTP requests to load balancer")
			return
		}

		for i := 0; i < times; i++ {
			req, err := http.NewRequest("GET", loadBalancerEntryPoint, nil)
			if err != nil {
				log.Printf("[Request %d] Failed to create request: %v", i+1, err)
				continue
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Printf("[Request %d] Request failed: %v", i+1, err)
				continue
			}
			log.Printf("[Request %d] Status: %s", i+1, resp.Status)
			resp.Body.Close()
		}
	},
}

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test load balancer connection with backends",
}

func init() {
	httpTestCmd.Flags().IntVar(&times, "times", 0, "Number of test requests to send")
	_ = httpTestCmd.MarkFlagRequired("times")

	TestCmd.AddCommand(httpTestCmd)
	RootCmd.AddCommand(TestCmd)
}

