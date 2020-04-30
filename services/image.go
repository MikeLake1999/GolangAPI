package services

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/gift"
)

type ImageSize struct {
	Width int
	Path  string
}

var Widths = [...]int{1920, 1600, 1280, 1024, 800, 256}

// load image file to image.Image struct
func LoadImage(filename string) (img image.Image, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	img, _, err = image.Decode(f)
	if err != nil {
		return
	}
	return
}
func (img ImageSize) Resize(w int) (dst string, err error) {
	g := gift.New(
		gift.Resize(w, 64, gift.LanczosResampling),
	)
	imageSrc, err := LoadImage(img.Path)
	if err != nil {
		return
	}

	imageDst := image.NewRGBA(g.Bounds(imageSrc.Bounds()))
	g.Draw(imageDst, imageSrc)

	dst = fmt.Sprintf(
		"%s/%s_%dw.jpg",
		filepath.Dir(img.Path),
		FilenameWithoutExt(filepath.Base(img.Path)),
		w,
	)

	err = SaveImage(dst, imageDst)

	return dst, err
}

// save image.Image struct to jpeg file
func SaveImage(filename string, img image.Image) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	err = jpeg.Encode(f, img, &jpeg.Options{
		Quality: 80,
	})
	if err != nil {
		return
	}

	return
}

// get file name without extension
func FilenameWithoutExt(fn string) string {
	return strings.TrimSuffix(fn, filepath.Ext(fn))
}

// resize image to given width
func Resize(src string, w int) (dst string, err error) {
	g := gift.New(
		gift.Resize(w, 0, gift.LanczosResampling),
	)
	imageSrc, err := LoadImage(src)
	if err != nil {
		return
	}

	imageDst := image.NewRGBA(g.Bounds(imageSrc.Bounds()))
	g.Draw(imageDst, imageSrc)

	dst = fmt.Sprintf(
		"%s/%s_%dw.jpg",
		filepath.Dir(src),
		FilenameWithoutExt(filepath.Base(src)),
		w,
	)

	err = SaveImage(dst, imageDst)

	return
}

// get image dimension
func GetDimension(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, err
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}

func ResizeAll(src string) (sizes map[int]string, err error) {
	sizes = map[int]string{}

	w, _, err := GetDimension(src)
	if err != nil {
		return
	}

	var cimg = make(chan ImageSize)
	counter := 0

	for _, width := range Widths {
		if w > width {
			counter++
			go func(w int, cimg chan ImageSize) {
				dst, err := Resize(src, w)
				if err != nil {
					return
				}
				cimg <- ImageSize{
					Width: w,
					Path:  dst,
				}
			}(width, cimg)
		}
	}

	for ; counter > 0; counter-- {
		is := <-cimg
		sizes[is.Width] = is.Path
	}

	return
}
