package nut

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Run(prefix string, name string) {

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGUSR1,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	exit_chan := make(chan int)

	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGUSR1:
				PollUPS(prefix, name)
			case syscall.SIGHUP:
				PollUPS(prefix, name)
			case syscall.SIGTERM:
				fmt.Println("Exiting on SIGTERM")
				exit_chan <- 0
			case syscall.SIGINT:
				fmt.Println("Exiting on SIGINT")
				exit_chan <- 0
			default:
				fmt.Println("Exiting on SIGINT")
				exit_chan <- 0
			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)
}
