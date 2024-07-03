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

var setMessageCmd = &cobra.Command{
	Use:   "setMessage [flags] contract_address [message]",
	Short: "Set the message inside the MessageBox contract",
	Args:  cobra.RangeArgs(1, 2),
	Run:   SetMessage,
}

func init() {
	RootCmd.AddCommand(setMessageCmd)
}

func SetMessage(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	message := "Hello, world!"
	if len(args) > 1 {
		message = args[1]
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

	// Store message in box.
	auth, err := conn.PrepareNextTx(ctx)
	if err != nil {
		ExitWithError("Failed to prepare next tx", err)
	}

	fmt.Fprintf(os.Stderr, "Storing message in MessageBox...\n")

	setMsgTx, err := mb.SetMessage(auth, message)
	if err != nil {
		ExitWithError("Failed to create set tx", err)
	}

	_, err = bind.WaitMined(ctx, conn.Sapphire, setMsgTx)
	if err != nil {
		ExitWithError("Failed to store message", err)
	}

	fmt.Fprintf(os.Stderr, "Message stored successfully.\n")
}
