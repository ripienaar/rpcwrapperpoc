package puppetclient

import (
	"time"

	"github.com/choria-io/mcorpc-agent-provider/mcorpc"
)

// ResultDetails is the details about a result
type ResultDetails struct {
	sender  string
	code    int
	message string
	ts      time.Time
}

func (d *ResultDetails) Sender() string {
	return d.sender
}

// OK determines if the request was successfull
func (d *ResultDetails) OK() bool {
	return mcorpc.StatusCode(d.code) == mcorpc.OK
}
