// This codebase is inspired by jmcvetta/randutil's WeightedChoice.

package main

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

type Choice struct {
	Item   string
	Weight float64
}

// weightedRandom uses weighted random selection to return one of the supplied
// choices. Weights of 0 are never selected. All other weight values are
// relative. E.g., if you have two choices both weighted 0.3, they will be
// returned equally often; and each will be returned 3 times as often as a
// choice weighted 0.1. The sum of all weights must be 1.

func weightedRandom(choices []Choice) (string, error) {
	var ret Choice
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Choices validation
	if len(choices) == 0 {
		return ret.Item, errors.New("no choices provided")
	}

	sum := 0.0
	for _, c := range choices {
		sum += c.Weight
	}
	// deal w/ floating point number
	sum = math.Round(sum)
	if sum != 1.0 {
		return ret.Item, errors.New("sum of weights must equal 1")
	}

	// Generate a random float in the range [0.0, 1.0)
	randNum := rng.Float64()

	for _, c := range choices {
		randNum = randNum - c.Weight
		if randNum < 0 {
			return c.Item, nil
		}
	}

	// The code should never reach here because the loop logic should always return a choice
	return ret.Item, errors.New("something has gone wrong seriously")
}
