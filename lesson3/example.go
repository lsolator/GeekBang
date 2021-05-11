package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(ctx)

	s1 := http.Server{Addr: ":8080"}
	s2 := http.Server{Addr: ":8090"}
	g.Go(func() error {
		fmt.Printf("start s1\n")
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		fmt.Printf("start s2\n")
		return s2.ListenAndServe()
	})

	// signal
	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signalChannel:
			fmt.Printf("Received signal: %s\n", sig)
			done()
		}
		return nil
	})

	// 监听ctx cancel
	g.Go(func() error {
		select {
		case <-ctx.Done():
			_ = s1.Shutdown(ctx)
			fmt.Printf("shut down s1\n")
			_ = s2.Shutdown(ctx)
			fmt.Printf("shut down s2\n")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("exit process")
}
