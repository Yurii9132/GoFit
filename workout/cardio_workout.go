package workout

import (
	"fmt"
	"time"
)

type CardioWorkout struct {
	duration     time.Duration
	distance     float64
	avgHeartRate float64
}

func (c CardioWorkout) WorkoutDuration() time.Duration {
	return c.duration
}

func (c CardioWorkout) CalloriesBurned() float64 {
	return c.duration.Minutes() * 10 * (c.avgHeartRate / 100)
}

func (c CardioWorkout) RecordStats() {
	fmt.Printf("Cardio workout record stats: %.2f miles in %.2f minutes\n", c.distance, c.duration.Minutes())
}

func (c CardioWorkout) GetType() string {
	return "Cardio"
}

func (s CardioWorkout) GetDuration() time.Duration {
	return s.duration
}

func NewCardioWorkout(duration time.Duration, distance float64, avgHeartRate float64) CardioWorkout {
	return CardioWorkout{
		duration:     duration,
		distance:     distance,
		avgHeartRate: avgHeartRate,
	}
}
