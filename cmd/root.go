package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"
)

const NetworkFlag = "network"

var RootCmd = &cobra.Command{
	Use:   "demo-starter",
	Short: "Oasis Sapphire confidential dApp example",
}

func init() {
	RootCmd.PersistentFlags().String(NetworkFlag, "sapphire-localnet", "name of network to connect to")
}

// ExitWithError terminates the program after writing the error to stderr.
func ExitWithError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", msg, err)
	os.Exit(1)
}

// GetNetworkAddress returns the dial address of the network passed via the
// network flag or exits with an error if the network isn't known.
func GetNetworkAddress() string {
	networks := map[string]string{
		"sapphire":          "https://sapphire.oasis.io",
		"sapphire-testnet":  "https://testnet.sapphire.oasis.io",
		"sapphire-localnet": "http://localhost:8545",
	}

	net, err := RootCmd.PersistentFlags().GetString(NetworkFlag)
	if err != nil {
		ExitWithError("GetNetworkAddress", fmt.Errorf("Please specify the network to connect to using --%s.", NetworkFlag))
	}
	net = strings.ToLower(net)

	addr, found := networks[net]
	if !found {
		validNets := []string{}
		for n := range networks {
			validNets = append(validNets, n)
		}

		ExitWithError("GetNetworkAddress", fmt.Errorf("Unknown network specified, please use one of the following: %s.", strings.Join(validNets, ", ")))
	}

	return addr
}

// ParseAddress converts the hex representation of an Ethereum address into
// common.Address or returns an error if the address is malformed.
func ParseAddress(addrHex string) (common.Address, error) {
	if strings.HasPrefix(addrHex, "0x") {
		addrHex = strings.Replace(addrHex, "0x", "", 1)
	}

	if len(addrHex) != 40 {
		return common.Address{}, fmt.Errorf("address is malformed")
	}

	return common.HexToAddress(addrHex), nil
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		ExitWithError("demo-starter", err)
	}
}
