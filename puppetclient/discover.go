package puppetclient

import (
	"context"
	"sync"
	"time"

	"github.com/choria-io/go-client/discovery/broadcast"
	"github.com/choria-io/go-protocol/protocol"
)

type BroadcastNS struct {
	nodeCache []string
	f         *protocol.Filter

	sync.Mutex
}

func (b *BroadcastNS) SetFilter(f *protocol.Filter) {
	b.Lock()
	defer b.Unlock()

	b.f = f
	b.nodeCache = []string{}
}

func (b *BroadcastNS) Discover(ctx context.Context, fw ChoriaFramework) ([]string, error) {
	b.Lock()
	defer b.Unlock()

	copier := func() []string {
		out := make([]string, len(b.nodeCache))
		for i, n := range b.nodeCache {
			out[i] = n
		}

		return out
	}

	if b.nodeCache != nil && len(b.nodeCache) > 0 {
		return copier(), nil
	}

	if b.nodeCache == nil {
		b.nodeCache = []string{}
	}

	if b.f == nil {
		b.f = protocol.NewFilter()
	}

	cfg := fw.Configuration()
	nodes, err := broadcast.New(fw).Discover(ctx, broadcast.Filter(b.f), broadcast.Timeout(time.Second*time.Duration(cfg.DiscoveryTimeout)))
	if err != nil {
		return []string{}, err
	}

	b.nodeCache = nodes

	return copier(), nil
}
