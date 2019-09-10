package puppetclient

import (
	"time"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// Reset resets the client options to use across requests to an empty list
func (p *PuppetClient) Reset() *PuppetClient {
	p.clientRPCOpts = []rpcclient.RequestOption{}
	p.ns = &BroadcastNS{}

	return p
}

// Filter sets the filter to use
func (p *PuppetClient) Filter(f *protocol.Filter) *PuppetClient {
	p.ns.SetFilter(f)
	return p
}

// Collective sets the collective to target
func (p *PuppetClient) Collective(c string) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.Collective(c))
	return p
}

// InBatches performs requests in batches
func (p *PuppetClient) InBatches(size int, sleep int) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.InBatches(size, sleep))
	return p
}

// Timeout configures the request timeout
func (p *PuppetClient) Timeout(t time.Duration) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.Timeout(t))
	return p
}

// DiscoveryTimeout configures the request discovery timeout, defaults to configured discovery timeout
func (p *PuppetClient) DiscoveryTimeout(t time.Duration) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.DiscoveryTimeout(t))
	return p
}

// LimitMethod configures the method to use when limiting targets - "random" or "first"
func (p *PuppetClient) LimitMethod(m string) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitMethod(m))
	return p
}

// LimitSize sets limits on the targets, either a number of a percentage like "10%"
func (p *PuppetClient) LimitSize(s string) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSize(s))
	return p
}

// LimitSeed sets the random seed used to select targets when limiting and limit method is "random"
func (p *PuppetClient) LimitSeed(s int64) *PuppetClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSeed(s))
	return p
}
