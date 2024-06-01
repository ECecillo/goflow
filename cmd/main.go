package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type FlowmodorTimer struct {
	startTime   time.Time
	stopTime    time.Time
	elapsedTime time.Duration
}

func NewFlowmodorTimer() *FlowmodorTimer {
	return &FlowmodorTimer{
		startTime:   time.Time{},
		stopTime:    time.Time{},
		elapsedTime: 0,
	}
}

func (fwt *FlowmodorTimer) StartFlowmodorTimer() {
	fmt.Println("Starting Flowmodor Timer")
	fwt.startTime = time.Now()
}

func (fwt *FlowmodorTimer) StopFlowmodorTimer() {
	fmt.Println("Stopping Flowmodor Timer")
	fwt.stopTime = time.Now()
}

func (fwt *FlowmodorTimer) GetElapsedTime() {
	fwt.elapsedTime = fwt.stopTime.Sub(fwt.startTime)
	fmt.Println("⏱️ Elapsed time: ", fwt.elapsedTime)
	fmt.Println("⏳ Amount of break: ", fwt.elapsedTime/5)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	flowmodorTime := NewFlowmodorTimer()

	for {
		fmt.Print("goflow>")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if line == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		switch line {
		case "start":
			flowmodorTime.StartFlowmodorTimer()
		case "stop":
			flowmodorTime.StopFlowmodorTimer()
		case "break":
			flowmodorTime.GetElapsedTime()
		default:
			fmt.Println("Unknown command")
		}

		fmt.Printf("So you said : %s \n", line)
	}
}
