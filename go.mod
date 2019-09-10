module github.com/ripienaar/puppet

go 1.13

replace golang.org/x/sys v0.0.0-20190726091711-fde4db37ae7a => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a

require (
	github.com/apex/log v1.1.1
	github.com/choria-io/go-choria v0.12.0
	github.com/choria-io/go-client v0.5.0
	github.com/choria-io/go-config v0.0.3
	github.com/choria-io/go-protocol v1.3.2
	github.com/choria-io/go-srvcache v0.0.6
	github.com/choria-io/mcorpc-agent-provider v0.7.1
	github.com/sirupsen/logrus v1.4.2
)
