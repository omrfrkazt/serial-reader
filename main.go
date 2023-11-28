package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "COM4", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	file, err := os.OpenFile("can_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error occured while opening .txt folder: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(s)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatalf("Error occured while reading data: %v", err)
		}

		logLine := fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), string(line))
		fmt.Print(logLine)
		_, err = file.WriteString(logLine)
		if err != nil {
			log.Fatalf("Error occured while writing data: %v", err)
		}
	}
}
