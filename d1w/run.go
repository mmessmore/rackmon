package d1w

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(prefix string) {

	device, err := FindW1Device()
	if err != nil {
		log.Fatalf("Failed to find device: %v", err)
	}

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
				Emit(device, prefix)
			case syscall.SIGHUP:
				Emit(device, prefix)
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
func Emit(device W1Device, prefix string) {
	res, err := device.Read()
	if err != nil {
		log.Fatalf("Failed to read temp: %v", err)
	}
	fmt.Println(res.Graphite(prefix))
}
