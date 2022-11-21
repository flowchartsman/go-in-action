package maybe

import (
	"math/rand"
	"time"
)

// pseudorandomSource generates psudo-random results from a seed taken from the
// environment's current time represented in nanoseconds since the Unix Epoch
// (one of the more common techniques)
var pseudorandomSource = rand.New(rand.NewSource(time.Now().UnixNano()))

// deterministicSource provides a constant random seed so any examples that need
// to demonstrate "randomness" with output will match those in the book
var deterministicSource = rand.New(rand.NewSource(42))

// PercentD simulates randomness to illustrate the different paths
func Percent(percent int) bool {
	return pseudorandomSource.Intn(100) <= percent
}

// PercentD simulates randomness to illustrate the different paths, deterministically
func PercentD(percent int) bool {
	return deterministicSource.Intn(100) <= percent
}

// Random returns a random int [0,rng)
func Random(rng int) int {
	return pseudorandomSource.Intn(rng)
}

// Random returns a random int [0,rng), deterministically
func RandomD(rng int) int {
	return deterministicSource.Intn(rng)
}
