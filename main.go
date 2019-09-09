package main

import (
	"context"
	"fmt"

	"github.com/choria-io/go-protocol/protocol"
	"github.com/ripienaar/puppet/puppetclient"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.NewEntry(logrus.New())
	logger.Logger.SetLevel(logrus.DebugLevel)

	puppet, err := puppetclient.New(puppetclient.Logger(logger))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f := protocol.NewFilter()
	err = f.AddFactFilter("country", "==", "mt")
	if err != nil {
		panic(err)
	}

	res, err := puppet.Filter(f).Disable().Message("testing 1,2,3").Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("success: %d fail: %d\n", res.Stats.OKCount(), res.Stats.FailCount())

	res, err = puppet.Filter(f).Disable().Message("testing 1,2,3").Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("success: %d fail: %d\n", res.Stats.OKCount(), res.Stats.FailCount())
}
