package main

import (
	"fmt"

	// [module name]/[package name from subdirectory]
	"myimageapp/imageManipulator"
)

func main() {
	// Create an ImageManipulator instance with an existing image
	im, err := imageManipulator.NewImageManipulatorWithImage("mustangs.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Draw a rectangle over the image
	im.DrawRectangle(150, 50, 560, 411)

	// Save the modified image to a file
	im.SaveToFile("mustangs.png")
}
