package phash_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"k8s.io/klog/v2"
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
		assert(t, hash == tCase.imageHash, "phash do not match")
	}
}

func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		klog.Infof("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}
