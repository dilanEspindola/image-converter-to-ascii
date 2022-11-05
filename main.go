package main

import (
	"convterImgToAscii/helpers"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func readImage(pathImg string) (image.Image, error) {

	file, err := os.Open(pathImg)
	helpers.HandleError(err)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err3 := image.Decode(file)

	return img, err3
}

func grayScale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func averegePixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayScale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}

func main() {
	const image string = "./images/102184902_p0_master1200.jpg"

	img, err := readImage(image)

	if err != nil {
		helpers.HandleError(err)
	}

	ramp := ".:-=+*#%@"
	// ramp := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,"^`'. "
	max := img.Bounds().Max
	scaleX, scaleY := 9, 5

	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			c := averegePixel(img, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp)*c/65538]))
		}
		fmt.Println()
	}
}
