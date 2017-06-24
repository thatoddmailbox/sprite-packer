package main

import (
	"image"
)

func Utils_Rotate90(input image.Image) image.Image {
	bounds := input.Bounds()
	output := image.NewRGBA(bounds)
	width := (bounds.Max.X - bounds.Min.X)
	height := (bounds.Max.Y - bounds.Min.Y)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			output.Set((width - 1) - j, i, input.At(i, j))
		}
	}
	return output
}