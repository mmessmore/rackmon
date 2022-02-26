package nut

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func PollUPS(prefix string, name string) {
	log.Printf("INFO polling ups: %s", name)
	log.Printf("INFO using prefix: %s", prefix)

	out, err := exec.Command("upsc", name).Output()
	if err != nil {
		log.Printf("ERROR failed polling ups %s: %v", name, err)
	}

	output := strings.Split(string(out), "\n")
	for _, line := range output {
		log.Printf("INFO %s", line)
		key, val, err := parseLine(line)
		if err != nil {
			continue
		}
		key = fmt.Sprintf("%s.%s", prefix, key)

		fmt.Printf("%s: %f:.2d", key, val)
	}
}

func parseLine(line string) (string, float64, error) {
	pieces := strings.Split(line, ": ")
	key := pieces[0]
	val, err := strconv.ParseFloat(pieces[1], 64)

	if err != nil {
		return key, 0, err
	}

	return key, val, nil
}
