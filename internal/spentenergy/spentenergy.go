package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be negative")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be a positive number")
	}
	if height <= 0 {
		return 0, errors.New("growth must be a positive number")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be negative")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be a positive number")
	}
	if height <= 0 {
		return 0, errors.New("growth must be a positive number")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	hours := duration.Hours()
	speed := dist / hours
	return speed
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	resultDist := ((height * stepLengthCoefficient) * float64(steps)) / mInKm
	return resultDist
}
