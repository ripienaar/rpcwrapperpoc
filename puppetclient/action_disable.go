package puppetclient

import (
	"context"
	"fmt"

	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// Requestor is a generic request handler
type Requestor struct {
	client *PuppetClient
	action string
	args   map[string]interface{}
}

// DisableRequestor performs a RPC request
type DisableRequestor struct {
	r *Requestor
}

// DisableResult is the result from a disable request
type DisableResult struct {
	Stats   *rpcclient.Stats
	outputs []*DisableOutput
}

// DisableOutput is the output from the disable action
type DisableOutput struct {
	reply map[string]interface{}
}

// Status returns the status value
func (d *DisableOutput) Status() string {
	v, ok := d.reply["status"]
	if !ok {
		return ""
	}

	return v.(string)
}

// Enabled indicates the agent status
func (d *DisableOutput) Enabled() bool {
	v, ok := d.reply["enabled"]
	if !ok {
		return false
	}

	return v.(bool)
}

// Disable disable the Puppet agent
func (p *PuppetClient) Disable() *DisableRequestor {
	d := &DisableRequestor{
		r: &Requestor{
			args:   make(map[string]interface{}),
			action: "disable",
			client: p,
		},
	}

	return d
}

// Message supply a reason for disabling the Puppet agent
func (d *DisableRequestor) Message(m string) *DisableRequestor {
	d.r.args["message"] = m

	return d
}

func (r *Requestor) Do(ctx context.Context) (*rpcclient.Stats, error) {
	r.client.Infof("Starting discovery")
	targets, err := r.client.ns.Discover(ctx, r.client.fw)
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

	r.client.Infof("Invoking %s action with %#v", r.action, r.args)

	res, err := agent.Do(ctx, r.action, r.args, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not perform disable request: %s", err)
	}

	return res.Stats(), nil
}

// Do performs the request
func (d *DisableRequestor) Do(ctx context.Context) (*DisableResult, error) {
	res, err := d.r.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &DisableResult{Stats: res}, nil
}
