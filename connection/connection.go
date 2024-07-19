package connection

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	sapphire "github.com/oasisprotocol/sapphire-paratime/clients/go"
)

// Connection represents a connection to the network via our account.
type Connection struct {
	// Client is an instance of the Go Ethereum client.
	Client *ethclient.Client

	// Sapphire is a wrapped backend that enables confidentiality.
	Sapphire sapphire.WrappedBackend

	// ChainID is the network ID of the network that we're connected to.
	ChainID *big.Int

	// PrivateKey is our private key (loaded from the env var PRIVATE_KEY).
	PrivateKey *ecdsa.PrivateKey

	// PublicKey is our public key (derived from PrivateKey).
	PublicKey *ecdsa.PublicKey

	// Address is our wallet's address (derived from PublicKey).
	Address common.Address
}

// NewConnection creates a new Connection to the given address.
func NewConnection(ctx context.Context, where string) (*Connection, error) {
	var (
		c   Connection
		err error
		ok  bool
	)

	c.Client, err = ethclient.Dial(where)
	if err != nil {
		return nil, err
	}

	c.ChainID, err = c.Client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get chain ID: %v", err)
	}

	pk := os.Getenv("PRIVATE_KEY")
	if strings.HasPrefix(pk, "0x") {
		pk = strings.Replace(pk, "0x", "", 1)
	}

	c.PrivateKey, err = crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("unable to get private key from environment: %v", err)
	}
	c.PublicKey, ok = c.PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not in ECDSA format")
	}
	c.Address = crypto.PubkeyToAddress(*c.PublicKey)

	wc, err := sapphire.WrapClient(c.Client, func(digest [32]byte) ([]byte, error) {
		return crypto.Sign(digest[:], c.PrivateKey)
	})
	if err != nil {
		return nil, fmt.Errorf("unable to wrap backend: %v", err)
	}
	c.Sapphire = *wc

	return &c, nil
}

// PrepareNextTx returns a transactor for the next transaction.
func (c *Connection) PrepareNextTx(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.Sapphire.PendingNonceAt(ctx, c.Address)
	if err != nil {
		return nil, fmt.Errorf("unable to get pending nonce: %v", err)
	}

	gasPrice, err := c.Sapphire.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get suggested gas price: %v", err)
	}

	auth := c.Sapphire.Transactor(c.Address)
	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.GasPrice = gasPrice

	return auth, nil
}
