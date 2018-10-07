package image_processor

import (
	"os"
	"strings"
	"image"
	"fmt"
	"image/color"
	"image/jpeg"
	"github.com/rwcarlsen/goexif/exif"
)

const (
	dist = "image_processor/dist"
)

func Start() {
	ip := NewImageProcessor("./assets/flower_1.jpeg")
	ip.Init()
	defer ip.CloseFile()

	ip.GrayScale()
	fmt.Println(ip.Exif)
}

type ImageProcessor struct {
	Ext string
	Exif *exif.Exif
	FileName string
	Image *os.File
	Src string
}

func NewImageProcessor(src string) *ImageProcessor {
	return &ImageProcessor{Src: src}
}

/*
 * Set -> Image, Src, FileName, Ext, Exif
 */
func (ip *ImageProcessor) Init() error {
	file, err := os.Open(ip.Src)
	if err != nil {
		return err
	}

	ext := ip.getExt()
	fileName := ip.getFileName()

	exif, err := ip.getExif()
	if err != nil {
		return err
	}

	ip.Image = file
	ip.FileName = fileName
	ip.Ext = ext
	ip.Exif = exif

	return nil
}

func (ip *ImageProcessor) GrayScale() error {
	img, _, err := image.Decode(ip.Image)
	if err != nil {
		return err
	}

	dstFileName := fmt.Sprintf("%s/%s", dist, ip.FileName)
	dstfile, err := os.Create(dstFileName)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	bounds := img.Bounds()
	proccessedImage := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			proccessedImage.Set(x, y, gray)
		}
	}

	return jpeg.Encode(dstfile, proccessedImage, nil)
}

func (ip *ImageProcessor) CloseFile() {
	ip.Image.Close()
}

/*-----------------
  private methods
------------------*/
func (ip *ImageProcessor) getExt() string {
	pos := strings.LastIndex(ip.Src, ".")
	ext := ip.Src[:pos]
	return ext[1:]
}

func (ip *ImageProcessor) getFileName() string {
	pos := strings.LastIndex(ip.Src, "/")
	fn := ip.Src[pos:]
	return fn[1:]
}

func (ip *ImageProcessor) getExif() (*exif.Exif, error) {
	exif, err := exif.Decode(ip.Image)
	if err != nil {
		return nil, err
	}

	return exif, nil
}