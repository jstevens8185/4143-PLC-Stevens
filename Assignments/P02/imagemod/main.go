package main

import (
	"./imageManipulator/imageManipulator.go"
)

func main() {
	// Create imageManipulator instance
	im := imageManipulator.NewImageManipulator(800, 600)

	// Draw rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save image to a file
	im.SaveToFile("rectangle.png")
}
