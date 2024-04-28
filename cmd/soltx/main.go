package main

import (
	"context"
	"fmt"
	"github.com/sol-tx/soltx"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	go shutdown(cancel, quit)

	if len(os.Args) < 2 {
		panic("args is invalid")
	}
	workSpace := os.Args[1]

	s := soltx.New(ctx, workSpace)
	s.Service()
}

func shutdown(cancel context.CancelFunc, quit <-chan os.Signal) {
	osCall := <-quit
	fmt.Printf("System call: %v,  trader shutting down......\n", osCall)
	cancel()
}
