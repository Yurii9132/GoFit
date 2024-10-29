package main

import (
	"time"
)

type Workout interface {
	Duration() time.Duration
	CalloriesBurned() float64
	RecordStats()
	GetType() string
}