package puppetclient

import (
	"github.com/sirupsen/logrus"
)

type initOptions struct {
	cfgFile string
	logger  *logrus.Entry
}

// InitializationOption is an optional setting used to initialize the client
type InitializationOption func(opts *initOptions)

// ConfigFile sets the config file to use, when not set will use the user default
func ConfigFile(f string) InitializationOption {
	return func(o *initOptions) {
		o.cfgFile = f
	}
}

// Logger sets the logger to use else one is made via the choria framework
func Logger(l *logrus.Entry) InitializationOption {
	return func(o *initOptions) {
		o.logger = l
	}
}
