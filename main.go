package main

import (
	"log"
	"time"
)

func main() {

	device, err := FindDevice()
	if err != nil {
		log.Fatalf("Failed to find device: %v", err)
	}

	for {
		res, err := device.Read()
		if err != nil {
			log.Fatalf("Failed to read temp: %v", err)
		}

		log.Printf("Stamp: %d, Temp: %.2f\n", res.TimeStamp, res.Temperature)
		time.Sleep(5 * time.Second)
	}
}
