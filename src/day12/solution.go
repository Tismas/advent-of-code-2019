package main

import (
	"fmt"
	"helpers"
	"math"
	"os"
)

func comparator(x1, x2 int) int {
	if x1 < x2 {
		return 1
	}
	if x1 == x2 {
		return 0
	}
	return -1
}

func updateVelocity(currentMoon *helpers.Movable, moons []helpers.Movable) {
	for _, moon := range moons {
		currentMoon.Vx += comparator(currentMoon.X, moon.X)
		currentMoon.Vy += comparator(currentMoon.Y, moon.Y)
		currentMoon.Vz += comparator(currentMoon.Z, moon.Z)
	}
}
func updatePosition(moon *helpers.Movable) {
	moon.X += moon.Vx
	moon.Y += moon.Vy
	moon.Z += moon.Vz
}

func main() {
	moons := helpers.ReadPositions(os.Open("./input.txt"))

	for i := 0; i < 1000; i++ {
		for i := range moons {
			updateVelocity(&moons[i], moons)
		}
		for i := range moons {
			updatePosition(&moons[i])
		}
	}

	total := 0.0

	for _, moon := range moons {
		pot := math.Abs(float64(moon.X)) + math.Abs(float64(moon.Y)) + math.Abs(float64(moon.Z))
		kin := math.Abs(float64(moon.Vx)) + math.Abs(float64(moon.Vy)) + math.Abs(float64(moon.Vz))
		total += pot * kin
	}

	fmt.Println("Part 1: ", total)
}
