package workout

import (
	"time"
)

// Workout interface
type Workout interface {
	WorkoutDuration() time.Duration
	CalloriesBurned() float64
	RecordStats()
	GetType() string
	GetDuration() time.Duration
}
