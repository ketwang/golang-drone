package singal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var (
	// hup used for reload configuration
	reloadingSignalList   = []os.Signal{}
	terminationSignalList = []os.Signal{syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT}
)

func init() {
	//default handling: ignore sig hup signal
	signal.Ignore(reloadingSignalList...)
}

func WitchSingalsContext(ctx context.Context) context.Context {
	newCtx, cancel := context.WithCancel(ctx)

	// define chan to receive signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, terminationSignalList...)

	// goroute waiting for upper ctx done event
	go func() {
		defer cancel()

		select {
		case <-ctx.Done():
		case <-sig:
		}
	}()

	return newCtx
}

func WithReloadingContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	// reset sig hub handling
	sig := make(chan os.Signal, 1)
	signal.Reset(reloadingSignalList...)
	signal.Notify(sig, reloadingSignalList...)

	go func() {
		defer cancel()

		select {
		case <-ctx.Done():
		case <-sig:
		}
	}()

	return ctx
}
