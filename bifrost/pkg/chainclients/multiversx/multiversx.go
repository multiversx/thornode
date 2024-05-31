package multiversx

import (
	"fmt"
	"math/big"

	"gitlab.com/thorchain/thornode/bifrost/thorclient/types"
	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/config"
)

// Client is the glue-code struct able to interface the MultiversX chain
type Client struct {
}

// NewClient is the constructor function that creates a new instance of type MultiversX Client
func NewClient() (*Client, error) {
	return &Client{}, nil
}

// Start starts the chain client with the given queues.
func (c *Client) Start(globalTxsQueue chan types.TxIn, globalErrataQueue chan types.ErrataBlock, globalSolvencyQueue chan types.Solvency) {
	// TODO: add implementation
}

// Stop stops the chain client.
func (c *Client) Stop() {
	// TODO: add implementation
}

// IsBlockScannerHealthy returns true if the block scanner is healthy.
func (c *Client) IsBlockScannerHealthy() bool {
	// TODO: add implementation

	return true
}

// SignTx returns the signed transaction.
func (c *Client) SignTx(tx types.TxOutItem, height int64) ([]byte, []byte, *types.TxInItem, error) {
	// TODO: add implementation

	return nil, nil, nil, fmt.Errorf("not implemented")
}

// BroadcastTx broadcasts the transaction and returns the transaction hash.
func (c *Client) BroadcastTx(txOutItem types.TxOutItem, hexTx []byte) (string, error) {
	// TODO: add implementation

	return "", fmt.Errorf("not implemented")
}

// GetHeight returns the current height of the chain.
func (c *Client) GetHeight() (int64, error) {
	// TODO: add implementation

	return 0, fmt.Errorf("not implemented")
}

// GetAddress returns the address for the given public key.
func (c *Client) GetAddress(poolPubKey common.PubKey) string {
	// TODO: add implementation

	return ""
}

// GetAccount returns the account for the given public key.
func (c *Client) GetAccount(poolPubKey common.PubKey, height *big.Int) (common.Account, error) {
	// TODO: add implementation

	return common.Account{}, fmt.Errorf("not implemented")
}

// GetAccountByAddress returns the account for the given address.
func (c *Client) GetAccountByAddress(address string, height *big.Int) (common.Account, error) {
	// TODO: add implementation

	return common.Account{}, fmt.Errorf("not implemented")
}

// GetChain returns the chain.
func (c *Client) GetChain() common.Chain {
	// TODO: add implementation

	return ""
}

// GetConfig returns the chain configuration.
func (c *Client) GetConfig() config.BifrostChainConfiguration {
	// TODO: add implementation

	return config.BifrostChainConfiguration{}
}

// OnObservedTxIn is called when a new observed tx is received.
func (c *Client) OnObservedTxIn(txIn types.TxInItem, blockHeight int64) {
	// TODO: add implementation
}

// GetConfirmationCount returns the confirmation count for the given tx.
func (c *Client) GetConfirmationCount(txIn types.TxIn) int64 {
	// TODO: add implementation

	return 0
}

// ConfirmationCountReady returns true if the confirmation count is ready.
func (c *Client) ConfirmationCountReady(txIn types.TxIn) bool {
	// TODO: add implementation

	return false
}

// GetBlockScannerHeight returns block scanner height for chain
func (c *Client) GetBlockScannerHeight() (int64, error) {
	// TODO: add implementation

	return 0, fmt.Errorf("not implemented")
}

// GetLatestTxForVault returns last observed and broadcasted tx for a particular vault and chain
func (c *Client) GetLatestTxForVault(vault string) (string, string, error) {
	// TODO: add implementation

	return "", "", fmt.Errorf("not implemented")
}
