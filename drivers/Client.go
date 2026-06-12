Here is a production-quality `client.go` for your `drivers/bitcoin` package. It provides:

* Configurable timeout
* Reusable HTTP client
* Dependency injection
* Thread-safe defaults
* Future support for metrics and retries
* Clean constructor pattern

```go
package bitcoin

// Purpose: Bitcoin Core client configuration and initialization.
// Connects to:
//   - rpc.go
//   - blockchain.go
//   - mempool.go
//   - network.go
//   - peers.go
// Used by:
//   - api/handlers.go

import (
	"net/http"
	"time"
)

const (
	DefaultTimeout = 10 * time.Second
)

// RPCClient communicates with Bitcoin Core JSON-RPC.
type RPCClient struct {
	host string
	user string
	pass string

	httpClient *http.Client
}

// Config contains client configuration.
type Config struct {
	Host    string
	User    string
	Pass    string
	Timeout time.Duration
}

// NewRPCClient creates a Bitcoin Core RPC client.
//
// Example:
//
//	client := bitcoin.NewRPCClient(
//		bitcoin.Config{
//			Host: "localhost:18443",
//			User: "bitcoin
```
