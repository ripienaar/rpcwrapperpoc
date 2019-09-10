package main

import (
	"context"
	"fmt"

	"github.com/ripienaar/puppet/puppetclient"
	"github.com/sirupsen/logrus"
)

func disable(ctx context.Context, puppet *puppetclient.PuppetClient) {
	res, err := puppet.Disable().Message("testing 1,2,3").Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Responses: \n")
	res.EachOutput(func(o *puppetclient.DisableOutput) {
		if o.ResultDetails().OK() {
			fmt.Printf("%s: enabled: %v: %s\n", o.ResultDetails().Sender(), o.Enabled(), o.Status())
		} else {
			fmt.Printf("%s: %s\n", o.ResultDetails().Sender(), o.ResultDetails().StatusMessage())
		}
	})

	fmt.Printf("success: %d fail: %d\n", res.Stats().OKCount(), res.Stats().FailCount())
}

func enable(ctx context.Context, puppet *puppetclient.PuppetClient) {
	res, err := puppet.Enable().Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Responses: \n")
	res.EachOutput(func(o *puppetclient.EnableOutput) {
		if o.ResultDetails().OK() {
			fmt.Printf("%s: enabled: %v: %s\n", o.ResultDetails().Sender(), o.Enabled(), o.Status())
		} else {
			fmt.Printf("%s: %s\n", o.ResultDetails().Sender(), o.ResultDetails().StatusMessage())
		}
	})

	fmt.Printf("success: %d fail: %d\n", res.Stats().OKCount(), res.Stats().FailCount())

}

func main() {
	logger := logrus.NewEntry(logrus.New())
	logger.Logger.SetLevel(logrus.DebugLevel)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	puppet, err := puppetclient.New(puppetclient.Logger(logger))
	if err != nil {
		panic(err)
	}

	puppet.FactFilter("country=mt")

	disable(ctx, puppet)
	enable(ctx, puppet)
}
