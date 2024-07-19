package slice

import (
	"image"
	"math"
)

func CreateSlices(img image.Image, slices int) []*image.RGBA {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	imgList := []*image.RGBA{}

	sliceWidth := int(math.Floor(float64(imgWidth) / float64(slices)))
	sliceHeight := int(math.Floor(float64(imgHeight) / float64(slices)))

	var slicedImg *image.RGBA
	//  row loop
	for i := 0; i < slices; i++ {
		// column loop
		for j := 0; j < slices; j++ {
			slicedImg = image.NewRGBA(image.Rect(0, 0, sliceWidth, sliceHeight))

			// Calculate the starting position for the current slice
			startX := j * sliceWidth
			startY := i * sliceHeight

			// Copy pixels from the original image to the current slice
			for y := 0; y < sliceHeight; y++ {
				for x := 0; x < sliceWidth; x++ {
					slicedImg.Set(x, y, img.At(startX+x, startY+y))
				}
			}
			imgList = append(imgList, slicedImg)
		}
	}
	return imgList
}
