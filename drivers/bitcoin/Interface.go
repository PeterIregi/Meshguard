
package bitcoin

// Purpose: Defines the Bitcoin Core driver interface.
// Used by:
//   - api/handlers.go
//   - services/
//   - tests and mocks

// BitcoinDriver describes all supported Bitcoin Core operations.
type BitcoinDriver interface {

	// Health
	Ping() error
	Health() map[string]interface{}

	// Blockchain
	GetBlockchainInfo() (*BlockchainInfo, error)
	GetBlockCount() (int64, error)
	GetBlockHash(height int64) (string, error)
	GetChainTips() ([]ChainTip, error)

	// Blocks
	GetBlock(hash string) (*Block, error)

	// Mempool
	GetMempoolInfo() (*MempoolInfo, error)
	GetRawMempool() ([]string, error)

	// Network
	GetNetworkInfo() (*NetworkInfo, error)

	// Peers
	GetPeerInfo() ([]PeerInfo, error)

	// Mining
	GetMiningInfo() (*MiningInfo, error)

	
	
