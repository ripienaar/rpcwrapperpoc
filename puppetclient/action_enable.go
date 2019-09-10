package puppetclient

import (
	"context"
	"encoding/json"
	"github.com/choria-io/go-protocol/protocol"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
)

// EnableRequestor performs a RPC request
type EnableRequestor struct {
	r    *Requestor
	outc chan *EnableOutput
}

// EnableOutput is the output from the disable action
type EnableOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// EnableResult is the result from a disable request
type EnableResult struct {
	stats   *rpcclient.Stats
	outputs []*EnableOutput
}

// Stats is the rpc request stats
func (d *EnableResult) Stats() *rpcclient.Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *EnableOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *EnableOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *EnableOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Status returns the status value
func (d *EnableOutput) Status() string {
	v, ok := d.reply["status"]
	if !ok {
		return ""
	}

	return v.(string)
}

// Enabled indicates the agent status
func (d *EnableOutput) Enabled() bool {
	v, ok := d.reply["enabled"]
	if !ok {
		return false
	}

	return v.(bool)
}

// Do performs the request
func (d *EnableRequestor) Do(ctx context.Context) (*EnableResult, error) {
	eres := &EnableResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &EnableOutput{
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

		if d.outc != nil {
			d.outc <- output
			return
		}

		eres.outputs = append(eres.outputs, output)
	}

	res, err := d.r.Do(ctx, handler)
	if err != nil {
		return nil, err
	}

	eres.stats = res

	return eres, nil
}

// EachOutput iterates over all results received
func (d *EnableResult) EachOutput(h func(r *EnableOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}
