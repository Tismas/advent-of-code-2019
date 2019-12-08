package main

import (
	"fmt"
	"helpers"
	"os"
)

var printWrappers = map[string]string{
	"0": "\033[1;34m%s\033[0m",
	"1": "\033[1;31m%s\033[0m",
}

func countPixels(layer []string, targetPixel string) int {
	count := 0
	for _, pixel := range layer {
		if pixel == targetPixel {
			count++
		}
	}
	return count
}

func main() {
	file, e := os.Open("./input.txt")
	imageData := helpers.ReadAndSplitStrings(file, e, "")[0]
	const width = 25
	const height = 6
	const layerSize = width * height
	var layers [][]string
	for i := 0; i < len(imageData); i += layerSize {
		layers = append(layers, imageData[i:i+layerSize])
	}
	minZeros := layerSize + 1
	var minZerosLayer []string
	for _, layer := range layers {
		zerosCount := countPixels(layer, "0")
		if zerosCount < minZeros {
			minZeros = zerosCount
			minZerosLayer = layer
		}
	}
	result := countPixels(minZerosLayer, "1") * countPixels(minZerosLayer, "2")
	fmt.Println("Part 1: ", result)

	var finalImage []string = make([]string, layerSize)
	for i := 0; i < layerSize; i++ {
		finalImage[i] = "2"
	}
	for _, layer := range layers {
		for i, pixel := range layer {
			if finalImage[i] == "2" {
				finalImage[i] = pixel
			}
		}
	}

	fmt.Println("Part 2:")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf(printWrappers[finalImage[y*width+x]], "*")
		}
		fmt.Println()
	}
}
