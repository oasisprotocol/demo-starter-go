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

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the MessageBox contract",
	Args:  cobra.NoArgs,
	Run:   Deploy,
}

func init() {
	RootCmd.AddCommand(deployCmd)
}

func Deploy(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	// Connect to the network.
	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	// Deploy MessageBox contract.
	auth, err := conn.PrepareNextTx(ctx)
	if err != nil {
		ExitWithError("Failed to prepare next tx", err)
	}

	fmt.Fprintf(os.Stderr, "Deploying MessageBox contract...\n")

	mbAddr, deployTx, _, err := messageBox.DeployMessageBox(auth, conn.Sapphire)
	if err != nil {
		ExitWithError("Failed to create deploy tx", err)
	}

	_, err = bind.WaitDeployed(ctx, conn.Sapphire, deployTx)
	if err != nil {
		ExitWithError("Failed to deploy contract", err)
	}

	// Output deployed contract's address to stdout.
	fmt.Printf("%s\n", mbAddr.Hex())
}
