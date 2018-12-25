package cmd

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"github.com/spf13/cobra"
	contracts "github.com/gofunct/hack/contracts/token_service"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// Set up a connection to the server.
		conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Did not connect: %v\n", err)
			os.Exit(-1)
		}
		defer conn.Close()
		c := contracts.NewTokenClient(conn)

		// Contact the server and print out its response. setnamereq
		// Contact the server and print out its response.
		tx, err := c.Transfer(context.Background(), &contracts.TransferReq{
			Opts: &contracts.TransactOpts{
				PrivateKey: privateKey,
			},
		})

		if err != nil {
			fmt.Printf("Failed to transfer funds: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Set name in tx: %v\n", tx.TxHash)
	},
}
