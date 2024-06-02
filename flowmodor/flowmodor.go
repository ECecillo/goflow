// Package flowmodor provides functionality to start, stop,
// and calculate break time for a timer.
package flowmodor

import (
	"fmt"
	"time"
)

type FlowmodorTimer struct {
	StartTime   time.Time
	StopTime    time.Time
	ElapsedTime time.Duration
	BreakTime   time.Duration
}

func NewFlowmodorTimer() *FlowmodorTimer {
	return &FlowmodorTimer{
		StartTime:   time.Time{},
		StopTime:    time.Time{},
		ElapsedTime: 0,
		BreakTime:   0,
	}
}

// Start the timer of the FlowmodorTimer and return the current Start time.
func (fwt *FlowmodorTimer) Start() time.Time {
	fwt.StartTime = time.Now()
	// Ensuire that when we start the timer, stop timer is empty because it is used as a break point for our flowmodorTimer.
	if !fwt.StopTime.IsZero() {
		fwt.StopTime = time.Time{}
	}
	return fwt.StartTime
}

// Stop the timer of the FlowmodorTimer and return the current Stop time.
func (fwt *FlowmodorTimer) Stop() time.Time {
	fwt.StopTime = time.Now()
	// Ensure that we have started the timer before we can calculate the elapsed time.
	if !fwt.StartTime.IsZero() {
		fwt.ElapsedTime = fwt.StopTime.Sub(fwt.StartTime).Round(time.Second)
	}
	return fwt.StopTime
}

// Calculate the break time of the FlowmodorTimer and return the current Break time.
// Consider that when the break time is run, the start and stop time are reset.
func (fwt *FlowmodorTimer) Break() time.Duration {
	fwt.BreakTime = fwt.CalculateBreakTime(fwt.ElapsedTime)
	fwt.StartTime = time.Time{}
	fwt.StopTime = time.Time{}
	return fwt.BreakTime
}

// Rest all FlowmodorTimer timer to their initialial value.
func (fwt *FlowmodorTimer) Reset() {
	fwt.StartTime = time.Time{}
	fwt.StopTime = time.Time{}
	fwt.ElapsedTime = 0
	fwt.BreakTime = 0
}

func (fwt *FlowmodorTimer) FormatElapsedTime() string {
	stringifiedElapsedTime := fwt.ElapsedTime.String()
	return fmt.Sprintf("\n  ⏳ %s  \n\n", stringifiedElapsedTime)
}

func (fwt *FlowmodorTimer) FormatBreakTime() string {
	stringifiedBreakTime := fwt.BreakTime.String()
	return fmt.Sprintf("\n  Break Time ☕️ %s   \n\n", stringifiedBreakTime)
}

// Return calculated amount of break time we can take based on an elapsedTime value.
func (fwt *FlowmodorTimer) CalculateBreakTime(elapsedTime time.Duration) time.Duration {
	return elapsedTime / 5
}
