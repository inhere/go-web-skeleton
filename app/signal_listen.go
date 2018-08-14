package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// listenSignals Graceful start/stop server
func listenSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go handleSignals(sigChan)
}

// handleSignals handle process signal
func handleSignals(c chan os.Signal) {
	log.Print("Notice: System signal monitoring is enabled(watch: SIGINT,SIGTERM,SIGQUIT)\n")

	switch <-c {
	case syscall.SIGINT:
		fmt.Println("\nShutdown by Ctrl+C")
	case syscall.SIGTERM: // by kill
		fmt.Println("\nShutdown quickly")
	case syscall.SIGQUIT:
		fmt.Println("\nShutdown gracefully")
		// do graceful shutdown
	}

	// sync logs
	Logger.Sync()

	// unregister from eureka
	erkServer.Unregister()

	// 等待一秒
	time.Sleep(1e9 / 2)
	fmt.Println("\nBye...")

	os.Exit(0)
}
