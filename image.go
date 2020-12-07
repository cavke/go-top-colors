package main

import (
	"fmt"
	"image"
	"image/draw"
)

// Function that counts the colors in the given image with "At" approach
// Returns a map with key that represents color and value that is the count of that color
func getColorMapFromImageWithAt(img image.Image) map[string]uint {
	colorCounts := make(map[string]uint)

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Color's RGBA values are in the range [0, 65535].
			r, g, b, _ := img.At(x, y).RGBA()
			// Shifting by 8 reduces this to the range [0, 255].
			c := string([]byte{byte(r >> 8), byte(g >> 8), byte(b >> 8)})
			colorCounts[c]++
		}
	}

	return colorCounts
}

// Function that counts the colors in the given image with "Pix" approach
// Returns a map with key that represents color and value that is the count of that color
func getColorMapFromImageWithPix(img image.Image) map[string]uint {
	colorCounts := make(map[string]uint)

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			i := rgba.PixOffset(x, y)
			s := rgba.Pix[i : i+4 : i+4]
			c := string([]byte{s[0], s[1], s[2]})
			colorCounts[c]++
		}
	}

	return colorCounts
}

// Function that returns three most common colors from the color map
func getTopThreeColors(colorMap map[string]uint) []string {
	// contains top three colors, from highest (at index=0) to lowest color count
	topThree := make([]string, 3)

	for color, count := range colorMap {
		if count > colorMap[topThree[2]] { // check if the color's count is larger than lowest color count
			if count > colorMap[topThree[0]] { // check if the color's count is larger then the highest one, shift all others to one place right and put it in the first place
				topThree[2] = topThree[1]
				topThree[1] = topThree[0]
				topThree[0] = color
			} else {
				if count > colorMap[topThree[1]] { // check if the color's count is larger then the middle one
					topThree[2] = topThree[1]
					topThree[1] = color
				} else { // color's count is larger then the last one, just replace the last element
					topThree[2] = color
				}
			}
		}
	}

	// convert to user readable format
	for i, c := range topThree {
		if c != "" {
			topThree[i] = fmt.Sprintf("#%02X%02X%02X", c[0], c[1], c[2])
		}
	}

	return topThree
}
