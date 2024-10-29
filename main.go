package main

import (
	"fmt"
	"time"
)

func summarizeWorkouts(workouts []Workout) {
	for _, workout := range workouts {
		workout.RecordStats()
	}
}

func main() {
	workouts := make([]Workout, 0, 8)
	fmt.Println("Callories burned:")

	workouts = append(workouts, CardioWorkout{30 * time.Minute, 6, 120})
	workouts = append(workouts, StrengthWorkout{10 * time.Minute, 80, 8, 3})
	workouts = append(workouts, CardioWorkout{30 * time.Minute, 20, 90})
	workouts = append(workouts, StrengthWorkout{2 * time.Minute, 90, 5, 1})

	for i, workout := range workouts {
		fmt.Printf("Workout #%d callories burned %.2f\n", i, workout.CalloriesBurned())
	}

	fmt.Println("Workout Summary:")
	summarizeWorkouts(workouts)

}
