package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "TraffikOne is a open-source Load Balancer based in Round Robin algorithms and routing rules",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
	fmt.Println("Root command runned.")
}