package raytracer

import (
	"image"
	"image/png"
	"image/color"
	"os"
	"math"
)


type Image struct {
	Img *image.RGBA
	samplesPerPixel int
}


func NewImage(width int, height int, samplesPerPixel int) Image {
	img := Image{}
	img.Img = image.NewRGBA(image.Rect(0, 0, width, height))
	img.samplesPerPixel = samplesPerPixel

	return img
}


func (img *Image) WriteColour(x int, y int, pixel Colour) {
	samples := 1.0 / float64(img.samplesPerPixel)

	// sqrt -> gamma correction
	rPixel := math.Sqrt(pixel.X * samples)
	gPixel := math.Sqrt(pixel.Y * samples)
	bPixel := math.Sqrt(pixel.Z * samples)

	r := uint8(256 * Clamp(rPixel, 0.0, 0.999))
	g := uint8(256 * Clamp(gPixel, 0.0, 0.999))
	b := uint8(256 * Clamp(bPixel, 0.0, 0.999))

	c := color.RGBA{r, g, b,  255}

	img.Img.Set(x, y, c)
}


func (img *Image) SaveAsPng(filename string) {
	f, _ := os.Create(filename)
	png.Encode(f, img.Img)
}
