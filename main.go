package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	rh := initRequestHandler()

	ctx, cancelFunc := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)

	// subscribe chan to signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// handle sig, cancel execution after receiving sigterm, sigint
	go func() {
		<-sigs
		cancelFunc()
	}()

	// parses cryptocompare every minute and saves it to pg
	go func() {
		rh.ParseCryptoCompare(ctx)
	}()

	rh.serveHTTP(ctx)
}
