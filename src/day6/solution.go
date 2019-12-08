package main

import (
	"fmt"
	"helpers"
	"os"
	"strings"
)

type spaceObject struct {
	parent   *spaceObject
	children []*spaceObject
	name     string
}

func distanceBetween(obj1, obj2 *spaceObject) int {
	distance := 0
	var obj1Parents []*spaceObject
	for current := obj1; current.parent != nil; current = current.parent {
		obj1Parents = append(obj1Parents, current)
	}
	for current := obj2; current.parent != nil; current = current.parent {
		for i, obj := range obj1Parents {
			if obj.name == current.name {
				distance += i
				return distance - 2
			}
		}
		distance++
	}
	return -1
}

func main() {
	orbitsMap := helpers.ReadStrings(os.Open("./input.txt"))
	COM := &spaceObject{name: "COM"}
	var YOU *spaceObject
	var SAN *spaceObject
	spaceObjects := []*spaceObject{COM}
	for _, orbit := range orbitsMap {
		arr := strings.Split(orbit, ")")
		parentName := arr[0]
		childName := arr[1]
		var parent *spaceObject
		var child *spaceObject
		for _, spaceObject := range spaceObjects {
			if spaceObject.name == parentName {
				parent = spaceObject
			}
			if spaceObject.name == childName {
				child = spaceObject
			}
		}
		if parent == nil {
			parent = &spaceObject{name: parentName}
			spaceObjects = append(spaceObjects, parent)
			if parentName == "YOU" {
				YOU = parent
			}
			if parentName == "SAN" {
				SAN = parent
			}
		}
		if child == nil {
			child = &spaceObject{name: childName}
			spaceObjects = append(spaceObjects, child)
			if childName == "YOU" {
				YOU = child
			}
			if childName == "SAN" {
				SAN = child
			}
		}
		child.parent = parent
		parent.children = append(parent.children, child)
	}

	totalOrbits := 0
	for _, spaceObject := range spaceObjects {
		current := spaceObject
		for current.parent != nil {
			totalOrbits++
			current = current.parent
		}
	}
	fmt.Println("Part 1: ", totalOrbits)
	fmt.Println("Part 2: ", distanceBetween(SAN, YOU))
}
