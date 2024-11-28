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

func main() {
	workouts := make([]workout.Workout, 0, 8)
	fmt.Println("Calories burned:")

	workouts = append(workouts, workout.CardioWorkout{Duration: 30 * time.Minute, Distance: 6, AvgHeartRate: 120})
	workouts = append(workouts, workout.StrengthWorkout{Duration: 10 * time.Minute, Weight: 80, Reps: 8, Sets: 3})
	workouts = append(workouts, workout.CardioWorkout{Duration: 30 * time.Minute, Distance: 5, AvgHeartRate: 90})
	workouts = append(workouts, workout.StrengthWorkout{Duration: 2 * time.Minute, Weight: 90, Reps: 5, Sets: 1})
	workouts = append(workouts, workout.CardioWorkout{Duration: 45 * time.Minute, Distance: 7, AvgHeartRate: 100})

	for i, workout := range workouts {
		fmt.Printf("Workout #%d calories burned: %.2f\n", i+1, workout.CalloriesBurned())
	}

	fmt.Println("Workout Summary:")
	summarizeWorkouts(workouts)

}
