package puppetclient

import (
	"context"
	"fmt"

	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// Requestor is a generic request handler
type Requestor struct {
	client *PuppetClient
	action string
	args   map[string]interface{}
}

// Do performs the request
func (r *Requestor) Do(ctx context.Context, handler func(pr protocol.Reply, r *rpcclient.RPCReply)) (*rpcclient.Stats, error) {
	r.client.Infof("Starting discovery")
	targets, err := r.client.ns.Discover(ctx, r.client.fw, r.client.filters)
	if err != nil {
		return nil, err
	}

	if len(targets) == 0 {
		return nil, fmt.Errorf("no nodes were discovered")
	}
	r.client.Infof("Discovered %d nodes", len(targets))

	agent, err := rpcclient.New(r.client.fw, r.client.ddl.Metadata.Name, rpcclient.DDL(r.client.ddl))
	if err != nil {
		return nil, fmt.Errorf("could not create client: %s", err)
	}

	opts := []rpcclient.RequestOption{rpcclient.Targets(targets)}
	for _, opt := range r.client.clientRPCOpts {
		opts = append(opts, opt)
	}
	opts = append(opts, rpcclient.ReplyHandler(handler))

	r.client.Infof("Invoking %s#%s action with %#v", r.client.ddl.Metadata.Name, r.action, r.args)

	res, err := agent.Do(ctx, r.action, r.args, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not perform disable request: %s", err)
	}

	return res.Stats(), nil
}