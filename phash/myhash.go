package phash

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"

	_ "golang.org/x/image/bmp"

	_ "golang.org/x/image/tiff"

	"phash-poc.mmt.com/transforms"

	"github.com/nfnt/resize"
)

var pixelPool32 = sync.Pool{
	New: func() interface{} {
		p := make([]float64, 1024)
		return &p
	},
}

var pixelPool64 = sync.Pool{
	New: func() interface{} {
		p := make([]float64, 4096)
		return &p
	},
}

func ImagePerceptualHash(file string) string {
	imageFile, _ := os.Open(file)
	im, _, _ := image.Decode(imageFile)
	imageSize := 64

	pixels := pixelPool64.Get().(*[]float64)
	resized := resize.Resize(uint(imageSize), uint(imageSize), im, resize.Lanczos3)
	transforms.Rgb2GrayFast(resized, pixels)
	flattens := transforms.DCT2DFast64(pixels)

	pixelPool32.Put(&flattens)
	median := MedianOfPixelsFast64(flattens[:])
	hash := uint64(0)

	for idx, p := range flattens {
		if p > median {
			hash |= 1 << uint(64-idx-1) // leftShiftSet
		}
	}

	return fmt.Sprintf("%016x", hash)
}

func MedianOfPixelsFast64(pixels []float64) float64 {
	tmp := [64]float64{}
	copy(tmp[:], pixels)
	l := len(tmp)
	pos := l / 2
	return quickSelectMedian(tmp[:], 0, l-1, pos)
}

// MedianOfPixelsFast64 function returns a median value of pixels.
// It uses quick selection algorithm.
func MedianOfPixelsFast32(pixels []float64) float64 {
	tmp := [32]float64{}
	copy(tmp[:], pixels)
	l := len(tmp)
	pos := l / 2
	return quickSelectMedian(tmp[:], 0, l-1, pos)
}

func quickSelectMedian(sequence []float64, low int, hi int, k int) float64 {
	if low == hi {
		return sequence[k]
	}

	for low < hi {
		pivot := low/2 + hi/2
		pivotValue := sequence[pivot]
		storeIdx := low
		sequence[pivot], sequence[hi] = sequence[hi], sequence[pivot]
		for i := low; i < hi; i++ {
			if sequence[i] < pivotValue {
				sequence[storeIdx], sequence[i] = sequence[i], sequence[storeIdx]
				storeIdx++
			}
		}
		sequence[hi], sequence[storeIdx] = sequence[storeIdx], sequence[hi]
		if k <= storeIdx {
			hi = storeIdx
		} else {
			low = storeIdx + 1
		}
	}

	if len(sequence)%2 == 0 {
		return sequence[k-1]/2 + sequence[k]/2
	}
	return sequence[k]
}
