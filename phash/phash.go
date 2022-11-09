package phash

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/azr/phash"
	"github.com/corona10/goimagehash"
)

func GenerateImagePhash(imagePath string) string {
	imageFile, _ := os.Open(imagePath)
	defer imageFile.Close()

	decoded, _, _ := image.Decode(imageFile)
	width, height := 8, 8
	hash, _ := goimagehash.ExtAverageHash(decoded, width, height)

	return hash.ToString()
}

func GenerateAzrPhash(imagePath string) string {
	f, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	hash1 := phash.DTC(img)
	// hash2 := phash.DTC(img)
	return fmt.Sprintf("%x", hash1)
}
