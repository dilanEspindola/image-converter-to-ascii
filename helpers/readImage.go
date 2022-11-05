package helpers

import (
	"image"
	"image/color"
	"os"
)

func ReadImage(pathImg string) (image.Image, error) {

	file, err := os.Open(pathImg)
	HandleError(err)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err3 := image.Decode(file)

	return img, err3
}

func GrayScale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func AveregePixel(img image.Image, w, h, x, y int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < j+h && j < max.Y; j++ {
			sum += GrayScale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}
