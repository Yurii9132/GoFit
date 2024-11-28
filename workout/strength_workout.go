package workout

import (
	"fmt"
	"time"
)

// StrengthWorkout struct
type StrengthWorkout struct {
	Duration time.Duration
	Weight   int
	Reps     int
	Sets     int
}

func (s StrengthWorkout) WorkoutDuration() time.Duration {
	return s.Duration
}

func (s StrengthWorkout) CalloriesBurned() float64 {
	return s.Duration.Minutes() * 5 * (float64(s.Weight) / 10)
}

func (s StrengthWorkout) RecordStats() {
	fmt.Printf("Strength workout record stats: %d sets, %d reps, %d weight.\n", s.Sets, s.Reps, s.Weight)
}

func (s StrengthWorkout) GetType() string {
	return "Strength"
}