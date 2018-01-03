package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		width, height = 1400, 2018
	)
	var (
		Color00 = color.Gray{255}
		Color10 = color.Gray{10}
		Color01 = color.Gray{220}
		Color11 = color.RGBA{250,0,0,250}
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var c0 [height]int
	var c1 [height]int

	c0[0] = 1
	c0[1] = 1

	for i := width - 1; i >= 0; i-- {
		if i == width - 1 {
			img.Set(i, 0, Color10)
			img.Set(i, 1, Color10)
		} else {
			img.Set(i, 0, Color00)
			img.Set(i, 1, Color00)
		}
		for j := 2; j < height; j++ {
			val := c0[j-1] + c0[j-2] + c1[j]
			switch val {
			case 0:
				img.Set(i, j, Color00)
				c0[j] = 0
				c1[j] = 0
			case 1:
				img.Set(i, j, Color10)
				c0[j] = 1
				c1[j] = 0
			case 2:
				img.Set(i, j, Color01)
				c0[j] = 0
				c1[j] = 1
			case 3:
				img.Set(i, j, Color11)
				c0[j] = 1
				c1[j] = 1
			}
		}
		for j := 0; j < height; j++ {
			c0[j] = 0
		}
	}

	png.Encode(os.Stdout, img)

}
