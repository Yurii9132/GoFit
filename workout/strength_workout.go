package workout

import (
	"fmt"
	"time"
)

// StrengthWorkout struct
type StrengthWorkout struct {
	duration time.Duration
	weight   int
	reps     int
	sets     int
}

func (s StrengthWorkout) WorkoutDuration() time.Duration {
	return s.duration
}

func (s StrengthWorkout) CalloriesBurned() float64 {
	return s.duration.Minutes() * 5 * (float64(s.weight) / 10)
}

func (s StrengthWorkout) RecordStats() {
	fmt.Printf("Strength workout record stats: %d sets, %d reps, %d weight.\n", s.sets, s.reps, s.weight)
}

func (s StrengthWorkout) GetType() string {
	return "Strength"
}

func (s StrengthWorkout) GetDuration() time.Duration {
	return s.duration
}

func NewStrengthWorkout(duration time.Duration, weight int, reps int, sets int) StrengthWorkout {
	return StrengthWorkout{
		duration: duration,
		weight:   weight,
		reps:     reps,
		sets:     sets,
	}
}
