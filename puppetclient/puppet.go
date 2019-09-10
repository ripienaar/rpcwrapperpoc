package puppetclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"context"

	"github.com/choria-io/go-choria/choria"
	coreclient "github.com/choria-io/go-client/client"
	"github.com/choria-io/go-config"
	rpcclient "github.com/choria-io/mcorpc-agent-provider/mcorpc/client"
	"github.com/choria-io/mcorpc-agent-provider/mcorpc/ddl/agent"
)

// NodeSource discovers nodes
type NodeSource interface {
	Reset()
	Discover(ctx context.Context, fw *choria.Framework, filters []coreclient.Filter) ([]string, error)
}

// PuppetClient to the puppet agent
type PuppetClient struct {
	fw            *choria.Framework
	cfg           *config.Config
	ddl           *agent.DDL
	ns            NodeSource
	clientOpts    *initOptions
	clientRPCOpts []rpcclient.RequestOption
	filters       []coreclient.Filter
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
		filters:       []coreclient.Filter{},
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
		outc: nil,
		r: &Requestor{
			args:   make(map[string]interface{}),
			action: "disable",
			client: p,
		},
	}

	return d
}

// Enable enables the Puppet agent
func (p *PuppetClient) Enable() *EnableRequestor {
	d := &EnableRequestor{
		outc: nil,
		r: &Requestor{
			args:   make(map[string]interface{}),
			action: "enable",
			client: p,
		},
	}

	return d
}
