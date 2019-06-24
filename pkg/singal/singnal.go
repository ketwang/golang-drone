package singal

import (
	"context"
	"fmt"
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

func WithSignalsContext(ctx context.Context) context.Context {
	newCtx, cancel := context.WithCancel(ctx)

	// define chan to receive signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, terminationSignalList...)

	// goroute waiting for upper ctx done event
	go func() {
		defer cancel()

		select {
		case <-ctx.Done():
		case v := <-sig:
			fmt.Println("got signal: ", v)
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

/*
annotations:
    migrate: "false"
0688df96-9ca9-4320-9119-685e1ac49354
d31d8188-8ff0-4d38-b24d-24814198f86a
840fa503-68fa-41aa-96db-d155dec0a9b0
306f49f5-5ad9-4362-ae04-8bccd8fb5600
4a4bfce4-554c-4540-96a3-a743da7e4b2a
*/
