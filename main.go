package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	args := HandleArgs()

	device, err := FindDevice()
	if err != nil {
		log.Fatalf("Failed to find device: %v", err)
	}

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGUSR1,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	exit_chan := make(chan int)

	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGUSR1:
				Emit(device, args)
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

func usage() {
	prog := os.Args[0]
	fmt.Printf("%s\n\n", prog)
	fmt.Println("This monitors the temperature of a Dallas 1-Wire temperature")
	fmt.Println("sensor and is intended to work with the telegraf execd plugin.")
	fmt.Println()
	fmt.Println("This runs as a daemon outputting Graphite formatted metrics")
	fmt.Println("to STDOUT when it recieves a SIGUSR1.  It will exit on")
	fmt.Println("SIGTERM or SIGINT (^C).")
	fmt.Println()
	fmt.Printf("USAGE\n\n")
	fmt.Printf("  %s [metric_prefix...]\n", prog)
	fmt.Printf("  %s -h\n\n", prog)
	fmt.Println("ARGUMENTS")
	fmt.Println("  metric_prefix	all arguments are joined by a '.' as the metric name")
	fmt.Println()
	fmt.Println("OPTIONS")
	fmt.Println("  -h	This handy usage message")
	fmt.Println()
}

func HandleArgs() []string {
	var args []string
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		switch arg {
		case "-h":
			usage()
			os.Exit(0)
		case "--help":
			usage()
			os.Exit(0)
		default:
			if arg[0] == '-' {
				fmt.Printf("ERROR: Unknown option: %s\n\n", arg)
				usage()
				os.Exit(22)
			}
			args = append(args, arg)
		}
	}
	return args
}

func Emit(device Device, args []string) {
	res, err := device.Read()
	if err != nil {
		log.Fatalf("Failed to read temp: %v", err)
	}
	fmt.Println(res.Graphite(args...))
}
