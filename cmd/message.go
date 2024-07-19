package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"demo-starter/connection"
	messageBox "demo-starter/contracts/message-box"
)

var getMessageCmd = &cobra.Command{
	Use:   "message [flags] contract_address",
	Short: "Get the message stored inside the MessageBox contract",
	Args:  cobra.ExactArgs(1),
	Run:   GetMessage,
}

func init() {
	RootCmd.AddCommand(getMessageCmd)
}

func GetMessage(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	// Connect to the network.
	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	mb, err := messageBox.NewMessageBox(contractAddr, conn.Sapphire)
	if err != nil {
		ExitWithError("Unable to get instance of contract", err)
	}

	// Retrieve message from box.
	fmt.Fprintf(os.Stderr, "Retrieving message from MessageBox...\n")

	retrievedMsg, err := mb.Message(&bind.CallOpts{From: conn.Address})
	if err != nil {
		ExitWithError("Failed to retrieve message", err)
	}

	// Output retrieved message to stdout.
	fmt.Printf("%s\n", retrievedMsg)
}
