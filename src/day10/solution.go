package main

import (
	"fmt"
	"helpers"
	"math"
	"os"
	"sort"
)

type vector struct {
	y float64
	x float64
}

func isBetween(src, target, other vector) bool {
	dxc := other.x - src.x
	dyc := other.y - src.y

	dxl := target.x - src.x
	dyl := target.y - src.y

	cross := dxc*dyl - dyc*dxl
	if cross != 0 {
		return false
	}
	if math.Abs(float64(dxl)) >= math.Abs(float64(dyl)) {
		if dxl > 0 {
			return src.x <= other.x && other.x <= target.x
		}
		return target.x <= other.x && other.x <= src.x
	}
	if dyl > 0 {
		return src.y <= other.y && other.y <= target.y
	}
	return target.y <= other.y && other.y <= src.y
}

func isBlocked(a1, a2 vector, all []vector) bool {
	for _, other := range all {
		isA1 := other.x == a1.x && other.y == a1.y
		isA2 := other.x == a2.x && other.y == a2.y
		if isA1 || isA2 {
			continue
		}
		if isBetween(a1, a2, other) {
			return true
		}
	}
	return false
}

func getDetected(asteroids []vector, asteroid vector) []vector {
	var detected []vector
	for _, asteroid2 := range asteroids {
		if !isBlocked(asteroid, asteroid2, asteroids) {
			detected = append(detected, asteroid2)
		}
	}
	return detected
}

func getDetectedAsteroids(asteroids []vector) ([]vector, vector) {
	var maxDetected []vector
	var center vector

	for _, asteroid := range asteroids {
		detected := getDetected(asteroids, asteroid)
		if len(detected) > len(maxDetected) {
			maxDetected = detected
			center = asteroid
		}
	}

	return maxDetected, center
}

func getAngle(p vector, center vector) float64 {
	angle := math.Atan2(p.y-center.y, p.x-center.x)
	angle += math.Pi / 2
	if angle < 0 {
		angle += math.Pi * 2
	}
	angle = math.Mod(angle, 2*math.Pi)
	return angle
}

func main() {
	grid := helpers.ReadStrings(os.Open("./input.txt"))
	var asteroids []vector

	for y, row := range grid {
		for x, val := range row {
			if val == '#' {
				asteroids = append(asteroids, vector{float64(y), float64(x)})
			}
		}
	}

	detected, center := getDetectedAsteroids(asteroids)
	fmt.Println("Part 1: ", len(detected))

	sort.Slice(detected, func(i, j int) bool { return getAngle(detected[i], center) < getAngle(detected[j], center) })
	part2Asteroid := detected[199]
	fmt.Println("Part 2: ", part2Asteroid.x*100+part2Asteroid.y)
}
