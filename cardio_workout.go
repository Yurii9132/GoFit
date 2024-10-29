package main

import (
	"fmt"
	"time"
)

// CardioWorkout struct
type CardioWorkout struct {
	duration     time.Duration
	distance     float64
	avgHeartRate float64
}


func (c CardioWorkout) Duration() time.Duration {
	return c.duration
}

func (c CardioWorkout) CalloriesBurned() float64 {
	return c.duration.Minutes() * 10 * (c.avgHeartRate / 100)
}

func (c CardioWorkout) RecordStats() {
	fmt.Printf("Cardio workout: %.2f miles in %.2f minutes\n", c.distance, c.duration.Minutes())
}

func (c CardioWorkout) GetType() string {
	return "Cardio"
}