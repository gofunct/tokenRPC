package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	contracts "github.com/gofunct/hack/contracts/token_service"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&clientAddr, "client-addr", "c", "http://192.168.99.100:8545", "the ethereum client address")
	rootCmd.PersistentFlags().StringVarP(&privateKey, "key", "k", "", "deployer's private key")
	rootCmd.PersistentFlags().StringVarP(&contractAddr, "contract-addr", "t", "",  "contract address")
	rootCmd.PersistentFlags().StringVarP(&serverAddr, "proxy-addr", "s", "127.0.0.1:5555", "server address to dial")
}

// proxyCmd represents the serve command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// connect to ethereum client
		conn, err := ethclient.Dial(clientAddr)
		if err != nil {
			fmt.Printf("Failed to connect ethereum: %v\n", err)
			os.Exit(-1)
		}
		defer conn.Close()

		// Deploy contracts
		var addr common.Address
		if privateKey != "" {
			//account := account.New(conn, privateKey)

			// addr, _, _, err = contracts.DeployTokenService(account.TransactOpts(), conn)
			// if err != nil {
			//	fmt.Printf("Failed to deploy contract: %v\n", err)
			//	os.Exit(-1)
			//}

			fmt.Printf("Deployed contract: %v\n", addr.Hex())
		} else {
			addr = common.HexToAddress(contractAddr)
		}

		s := grpc.NewServer()
		contracts.RegisterTokenServer(s, contracts.N(addr, conn))

		lis, err := net.Listen("tcp", serverAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s.Serve(lis)
	},
}
