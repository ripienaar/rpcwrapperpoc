package puppetclient

import (
	"context"
	"encoding/json"
	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// DisableRequestor performs a RPC request
type DisableRequestor struct {
	r *Requestor
}

// DisableOutput is the output from the disable action
type DisableOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// DisableResult is the result from a disable request
type DisableResult struct {
	stats   *rpcclient.Stats
	outputs []*DisableOutput
}

// Stats is the rpc request stats
func (d *DisableResult) Stats() *rpcclient.Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *DisableOutput) ResultDetails() *ResultDetails {
	return d.details
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

// Message supply a reason for disabling the Puppet agent
func (d *DisableRequestor) Message(m string) *DisableRequestor {
	d.r.args["message"] = m

	return d
}

// Do performs the request
func (d *DisableRequestor) Do(ctx context.Context) (*DisableResult, error) {
	dres := &DisableResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &DisableOutput{
			reply: make(map[string]interface{}),
			details: &ResultDetails{
				sender:  pr.SenderID(),
				code:    int(r.Statuscode),
				message: r.Statusmsg,
				ts:      pr.Time(),
			},
		}

		err := json.Unmarshal(r.Data, &output.reply)
		if err != nil {
			d.r.client.Errorf("Could not decode reply from %s: %s", pr.SenderID(), err)
		}

		dres.outputs = append(dres.outputs, output)
	}

	res, err := d.r.Do(ctx, handler)
	if err != nil {
		return nil, err
	}

	dres.stats = res

	return dres, nil
}

// EachOutput iterates over all results received
func (d *DisableResult) EachOutput(h func(r *DisableOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}
