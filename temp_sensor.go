package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const w1_device_root = "/sys/bus/w1/devices/"

type Device struct {
	Name string
}

func FindDevice() (Device, error) {
	device := Device{}

	entries, err := ioutil.ReadDir(w1_device_root)
	if os.IsNotExist(err) {
		log.Printf("ERROR: %s does not exist.  Driver not loaded?", w1_device_root)
		return device, err
	}
	for e := range entries {
		name := entries[e].Name()
		if name != "w1_bus_master" {
			device.Name = name
			return device, nil
		}
	}

	log.Printf("ERROR: Could not find device")
	return device, err
}

func (device Device) Read() (TempHumidity, error) {

	th := TempHumidity{
		TimeStamp:   time.Now().Unix(),
		Device:      device.Name,
		Temperature: 0,
	}

	data, err := os.ReadFile(fmt.Sprintf("%s/%s/w1_slave", w1_device_root, device.Name))
	if err != nil {
		log.Printf("Could not read sensor data: %v", err)
		return th, err
	}

	foundTemp := false

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, "t=") {
			temp, err := convertTemp(line)
			if err != nil {
				log.Printf("ERROR: Could not convert temp %v\n%s", err, line)
				return th, err
			}
			th.Temperature = temp
			foundTemp = true
		}
	}

	if !foundTemp {
		log.Printf("EROR: Could not find a temperature reading")
		return th, errors.New("Could not find a temperature reading")
	}
	return th, nil
}

func convertTemp(text string) (float64, error) {
	parts := strings.Split(text, "=")
	if len(parts) != 2 {
		return 0, errors.New("Could not parse temp line")
	}
	temp, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	return float64(temp) / float64(1000), nil
}

type TempHumidity struct {
	TimeStamp   int64
	Device      string
	Temperature float64
}

func (t TempHumidity) Graphite(prefix ...string) string {
	var metric string

	if len(prefix) < 1 {
		metric = fmt.Sprintf("sensor.temp.%s", t.Device)
	} else {
		metric = strings.Join(prefix, ".")
	}

	return fmt.Sprintf("%s %.2f %d", metric, t.Temperature, t.TimeStamp)
}
