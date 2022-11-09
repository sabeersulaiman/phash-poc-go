package phash_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phash-poc.mmt.com/phash"
)

type ImagePhashCase struct {
	image     string
	imageHash string
}

func TestHashGeneration(t *testing.T) {
	images := []ImagePhashCase{
		{
			imageHash: "83daac7396aa4d92",
			image:     "images/3.jpg",
		},
		{
			imageHash: "f3980da7ec62e328",
			image:     "images/7.jpg",
		},
	}

	for _, tCase := range images {
		hash := phash.GenerateImagePhash(tCase.image)
		assert.Equal(t, tCase.imageHash, hash)
	}
}

func TestCorrectedHashGeneration(t *testing.T) {
	images := []ImagePhashCase{
		{
			image:     "../images/3.jpg",
			imageHash: "83daac7396aa4d92",
		},
		{
			image:     "../images/7.jpg",
			imageHash: "f3980da7ec62e328",
		},
		{
			image:     "../images/1.png",
			imageHash: "c285be6bd0375139",
		},
		{
			image:     "../images/2.png",
			imageHash: "c689f28028deb95f",
		},
		{
			image:     "../images/1.webp",
			imageHash: "c29f3ef0c80f650e",
		},
		{
			image:     "../images/2.webp",
			imageHash: "e619c33e09f424f3",
		},
		{
			image:     "../images/1.bmp",
			imageHash: "c5f82b093ec47707",
		},
		{
			image:     "../images/2.bmp",
			imageHash: "818e17f1ea215e5b",
		},
		{
			image:     "../images/1.tiff",
			imageHash: "dada03a55c1cb60f",
		},
	}

	for _, tCase := range images {
		hash := phash.ImagePerceptualHash(tCase.image)
		t.Logf("expected: %s; actual: %s, image: %s\n", tCase.imageHash, hash, tCase.image)
		assert.Equal(t, tCase.imageHash, hash)
	}
}
