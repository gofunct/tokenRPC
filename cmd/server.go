package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/gofunct/hack/contracts/token_service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

// proxyCmd represents the serve command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A grpc based ethereum server",
	Run: func(cmd *cobra.Command, args []string) {
		// connect to ethereum client
		conn, err := ethclient.Dial(ethConfig.EthAddr)
		if err != nil {
			fmt.Printf("Failed to connect ethereum: %v\n", err)
			os.Exit(-1)
		}
		defer conn.Close()

		// Deploy contracts
		var addr common.Address

		auth, err := bind.NewTransactor(strings.NewReader(ethConfig.PrivateKey), "admin")
		if err != nil {
			log.Fatalf("Failed to create authorized transactor: %v", err)
		}

		// Deploy a new awesome contract for the binding demo
		address, tx, token, err := contracts.DeployToken(auth, conn, new(big.Int), "Colecoin", "C$")
		if err != nil {
			log.Fatalf("Failed to deploy new token contract: %v", err)
		}

		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

		// Don't even wait, check its presence in the local pending state
		time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

		name, err := token.Name(&bind.CallOpts{Pending: true})
		if err != nil {
			log.Fatalf("Failed to retrieve pending name: %v", err)
		}
		fmt.Println("Pending name:", name)

		s := grpc.NewServer()
		contracts.RegisterTokenServer(s, contracts.NewService(addr, conn))

		lis, err := net.Listen("tcp", ethConfig.ServerAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s.Serve(lis)
	},
}
