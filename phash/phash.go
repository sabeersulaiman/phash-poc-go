package phash

import (
	"image/jpeg"
	"os"

	"github.com/corona10/goimagehash"
)

func GenerateImagePhash(imagePath string) string {
	image, _ := os.Open(imagePath)
	defer image.Close()

	decoded, _ := jpeg.Decode(image)
	width, height := 8, 8
	hash, _ := goimagehash.ExtAverageHash(decoded, width, height)

	return hash.ToString()
}
