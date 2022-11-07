package main

import (
	"convterImgToAscii/helpers"
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	value    string
	filename string
)

func main() {

	// helpers.GenerateFile()

	x := flag.Int("x", 10, "type value x")
	y := flag.Int("y", 8, "type value y")
	image := flag.String("image", "", "type image path")
	flag.Parse()

	img, err := helpers.ReadImage(*image)

	if err != nil {
		helpers.HandleError(err)
	}

	ramp := ".:-=+*#%@"
	max := img.Bounds().Max
	scaleX, scaleY := *x, *y

	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			c := helpers.AveregePixel(img, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp)*c/65538]))
			// value = string(ramp[len(ramp)*c/65538])
			// helpers.WriteFile(value)
		}
		fmt.Println()
	}
}
