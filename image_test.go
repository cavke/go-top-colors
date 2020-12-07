package main

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"image"
	"os"
	"testing"
)

const imgWith1ColorPng = "sample_images/img_with_1_color.png"
const imgWith1ColorJpg = "sample_images/img_with_1_color.jpg"
const imgWith4ColorPng = "sample_images/img_with_4_color.png"
const imgWith6ColorJpg = "sample_images/img_with_6_color.jpg"
const imgUsbSticks = "sample_images/img_usb_sticks.jpg"

type ImageTestSuite struct {
	suite.Suite
}

func TestImage(t *testing.T) {
	suite.Run(t, new(ImageTestSuite))
}

func (s *ImageTestSuite) TestColorsCount_imgWith1ColorPng() {
	imgPath := imgWith1ColorPng

	inputFile, err := os.Open(imgPath)
	s.NoError(err)
	defer inputFile.Close()
	image, typeStr, err := image.Decode(inputFile)
	s.NoError(err)

	colorCountMap := getColorMapFromImageWithPix(image)
	topThreeColors := getTopThreeColors(colorCountMap)

	fmt.Printf("Decoded image: %s, type: %s, colors: %s\n", imgPath, typeStr, topThreeColors)

	expected := []string{"#660099", "", ""}
	s.Equal(expected, topThreeColors, "Detected colors doesn't match")
}

func (s *ImageTestSuite) TestColorsCount_imgWith1ColorJpg() {
	imgPath := imgWith1ColorJpg

	inputFile, err := os.Open(imgPath)
	s.NoError(err)
	defer inputFile.Close()
	image, typeStr, err := image.Decode(inputFile)
	s.NoError(err)

	colorCountMap := getColorMapFromImageWithPix(image)
	topThreeColors := getTopThreeColors(colorCountMap)

	fmt.Printf("Decoded image: %s, type: %s, colors: %s\n", imgPath, typeStr, topThreeColors)

	expected := []string{"#7B67EC", "#7C68ED", "#7B67ED"}
	s.Equal(expected, topThreeColors, "Detected colors doesn't match")
}

func (s *ImageTestSuite) TestColorsCount_imgWith4ColorPng() {
	imgPath := imgWith4ColorPng

	inputFile, err := os.Open(imgPath)
	s.NoError(err)
	defer inputFile.Close()
	image, typeStr, err := image.Decode(inputFile)
	s.NoError(err)

	colorCountMap := getColorMapFromImageWithPix(image)
	topThreeColors := getTopThreeColors(colorCountMap)

	fmt.Printf("Decoded image: %s, type: %s, colors: %s\n", imgPath, typeStr, topThreeColors)

	expected := []string{"#F7DF74", "#68BCDD", "#298EB5"}
	s.Equal(expected, topThreeColors, "Detected colors doesn't match")
}

func (s *ImageTestSuite) TestColorsCount_imgWith6ColorJpg() {
	imgPath := imgWith6ColorJpg

	inputFile, err := os.Open(imgPath)
	s.NoError(err)
	defer inputFile.Close()
	image, typeStr, err := image.Decode(inputFile)
	s.NoError(err)

	colorCountMap := getColorMapFromImageWithPix(image)
	topThreeColors := getTopThreeColors(colorCountMap)

	fmt.Printf("Decoded image: %s, type: %s, colors: %s\n", imgPath, typeStr, topThreeColors)

	expected := []string{"#FFFFFF", "#000000", "#F3C300"}
	s.Equal(expected, topThreeColors, "Detected colors doesn't match")
}

func (s *ImageTestSuite) TestColorsCount_imgUsbSticks() {
	imgPath := imgUsbSticks

	inputFile, err := os.Open(imgPath)
	s.NoError(err)

	defer inputFile.Close()
	image, typeStr, err := image.Decode(inputFile)
	s.NoError(err)

	colorCountMap := getColorMapFromImageWithPix(image)
	topThreeColors := getTopThreeColors(colorCountMap)

	fmt.Printf("Decoded image: %s, type: %s, colors: %s\n", imgPath, typeStr, topThreeColors)

	expected := []string{"#6E81CD", "#6E81CE", "#FEFEFE"}
	s.Equal(expected, topThreeColors, "Detected colors doesn't match")
}
