package image_processor

import (
	"os"
	"strings"
	"image"
	"image/color"
	"image/jpeg"
	"fmt"
)

const (
	dist = "image_processor/dist"
)

type ImageProcessor struct {
	Image *os.File
	Src string
	FileName string
	Ext string
}

func NewImageProcessor(src string) *ImageProcessor {
	return &ImageProcessor{Src: src}
}

func (ip *ImageProcessor) Init() error {
	file, err := os.Open(ip.Src)
	if err != nil {
		return err
	}

	ext := ip.getExt()
	fileName := ip.getFileName()

	ip.Image = file
	ip.FileName = fileName
	ip.Ext = ext

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