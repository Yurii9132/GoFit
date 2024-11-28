package workout

import (
	"fmt"
	"time"
)

type CardioWorkout struct {
	Duration     time.Duration
	Distance     float64
	AvgHeartRate float64
}


func (c CardioWorkout) WorkoutDuration() time.Duration {
	return c.Duration
}

func (c CardioWorkout) CalloriesBurned() float64 {
	return c.Duration.Minutes() * 10 * (c.AvgHeartRate / 100)
}

func (c CardioWorkout) RecordStats() {
	fmt.Printf("Cardio workout record stats: %.2f miles in %.2f minutes\n", c.Distance, c.Duration.Minutes())
}

func (c CardioWorkout) GetType() string {
	return "Cardio"
}