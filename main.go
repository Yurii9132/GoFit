package main

import (
	"Gofit/workout"
	"fmt"
	"time"
)

func summarizeWorkouts(workouts []workout.Workout) {
	for _, workout := range workouts {
		workout.RecordStats()
	}
}

func summarizeDuration(workouts []workout.Workout) {
	var length time.Duration
	for _, workout := range workouts {
		length += workout.GetDuration()
	}
	fmt.Printf("Workout duration: %.2f", length.Minutes())
}

func main() {
	workouts := make([]workout.Workout, 0, 8)
	fmt.Println("Calories burned:")

	workouts = append(workouts, workout.NewCardioWorkout(30*time.Minute, 6, 120))
	workouts = append(workouts, workout.NewStrengthWorkout(10*time.Minute, 80, 8, 3))
	workouts = append(workouts, workout.NewCardioWorkout(30*time.Minute, 5, 90))
	workouts = append(workouts, workout.NewStrengthWorkout(2*time.Minute, 90, 5, 1))
	workouts = append(workouts, workout.NewCardioWorkout(45*time.Minute, 7, 100))

	for i, workout := range workouts {
		fmt.Printf("Workout #%d calories burned: %.2f\n", i+1, workout.CalloriesBurned())
	}

	fmt.Println("Workout Summary:")
	summarizeWorkouts(workouts)
	summarizeDuration(workouts)

}
