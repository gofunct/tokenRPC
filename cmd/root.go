package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	clientAddr      string //  "ethereum"
	serverAddr 		string
	gatewayPort 	int
	privateKey      string // "private_key"
	contractAddr  	string // "contract_address"
	logLevel 		string
	logHeaders 		bool
	swaggerFile 	string
	corsAllowOrigin string
	corsAllowCredentials string
	corsAllowMethods string
	corsAllowHeaders string
	apiPrefix string
)

func init() {
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(proxyCmd)
	rootCmd.AddCommand(gatewayCmd)
}

var rootCmd = &cobra.Command{
	Use:   "token",
	Short: "A brief description of your application",

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}