package puppetclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"context"

	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-config"
	"github.com/choria-io/go-protocol/protocol"
	"github.com/choria-io/go-srvcache"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
	"github.com/choria-io/mcorpc-agent-provider/mcorpc/ddl/agent"
	"github.com/sirupsen/logrus"
)

// ChoriaFramework is the choria framework to use
type ChoriaFramework interface {
	Configuration() *config.Config
	Logger(string) *logrus.Entry
	NewRequestID() (string, error)
	Certname() string
	MiddlewareServers() (srvcache.Servers, error)
	NewConnector(ctx context.Context, servers func() (srvcache.Servers, error), name string, logger *logrus.Entry) (conn choria.Connector, err error)
	NewMessage(payload string, agent string, collective string, msgType string, request *choria.Message) (msg *choria.Message, err error)
	NewTransportFromJSON(data string) (message protocol.TransportMessage, err error)
	NewReplyFromTransportJSON(payload []byte, skipvalidate bool) (msg protocol.Reply, err error)
}

// NodeSource discovers nodes
type NodeSource interface {
	Discover(ctx context.Context, fw ChoriaFramework) ([]string, error)
	SetFilter(*protocol.Filter)
}

// PuppetClient to the puppet agent
type PuppetClient struct {
	fw            ChoriaFramework
	cfg           *config.Config
	ddl           *agent.DDL
	ns            NodeSource
	clientOpts    *initOptions
	clientRPCOpts []rpcclient.RequestOption
}

// Metadata is the agent metadata
type Metadata struct {
	License     string `json:"license"`
	Author      string `json:"author"`
	Timeout     int    `json:"timeout"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// New creates a new client to the puppet agent
func New(opts ...InitializationOption) (client *PuppetClient, err error) {
	c := &PuppetClient{
		ddl:           &agent.DDL{},
		ns:            &BroadcastNS{},
		clientRPCOpts: []rpcclient.RequestOption{},
		clientOpts: &initOptions{
			cfgFile: choria.UserConfig(),
		},
	}

	for _, opt := range opts {
		opt(c.clientOpts)
	}

	c.fw, err = choria.New(c.clientOpts.cfgFile)
	if err != nil {
		return nil, fmt.Errorf("could not initialize choria: %s", err)
	}

	c.cfg = c.fw.Configuration()

	if c.clientOpts.logger == nil {
		c.clientOpts.logger = c.fw.Logger("puppet")
	}

	ddlj, err := base64.StdEncoding.DecodeString(rawDDL)
	if err != nil {
		return nil, fmt.Errorf("could not parse embedded DDL: %s", err)
	}

	err = json.Unmarshal(ddlj, c.ddl)
	if err != nil {
		return nil, fmt.Errorf("could not parse embedded DDL: %s", err)
	}

	return c, nil
}

// Metadata is the agent metadata this client supports
func (p *PuppetClient) Metadata() *Metadata {
	return &Metadata{
		License:     p.ddl.Metadata.License,
		Author:      p.ddl.Metadata.Author,
		Timeout:     p.ddl.Metadata.Timeout,
		Name:        p.ddl.Metadata.Name,
		Version:     p.ddl.Metadata.Version,
		URL:         p.ddl.Metadata.URL,
		Description: p.ddl.Metadata.Description,
	}
}

// SetNodeSource sets the discovery source
func (p *PuppetClient) SetNodeSource(ns NodeSource) *PuppetClient {
	p.ns = ns
	return p
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
