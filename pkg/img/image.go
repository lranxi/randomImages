package img

import (
	"api/pkg/file"
	"errors"
	"fmt"
	"github.com/azr/phash"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
	"strings"
)

const format string = `^(\.(png|PNG|jpg|JPG|jpeg|JPEG))$`

// ComputePHash 计算图片的感知hash
func ComputePHash(path string) (int64, error) {
	_, ok := file.IsExists(path)
	if !ok {
		return 0, errors.New(fmt.Sprintf("file does not exist: %s", path))
	}
	image, err := openImage(path)
	if err != nil {
		return 0, nil
	}
	phash := phash.DTC(image)
	return int64(phash), nil
}

// IsImageFormat 是否为图片格式
func IsImageFormat(filename string) bool {
	fileDot := strings.Index(filename, ".")
	fileType := filename[fileDot:]
	ok, _ := regexp.MatchString(format, fileType)
	return ok
}

// CalculateWidthHeight 计算图片宽高
func CalculateWidthHeight(filepath string) (width int, height int, err error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return 0, 0, err
	}

	c, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	width = c.Width
	height = c.Height

	return width, height, err
}

// Compression 图片压缩和resize
func Compression(source string, width, height uint) (string, error) {
	var err error
	var file *os.File
	// 检查图片格式
	reg, _ := regexp.Compile(`^.*\.((png)|(PNG)|(jpg)|(JPG)|(jpeg)|(JPEG))$`)
	if !reg.MatchString(source) {
		err = errors.New("%s is not a .png or .jpg file")
		return "", err
	}
	if file, err = os.Open(source); err != nil {
		return "", err
	}
	defer file.Close()
	name := file.Name()
	var img image.Image
	switch {
	case strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".PNG"):
		if img, err = png.Decode(file); err != nil {
			return "", err
		}
	case strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".JPG"):
		if img, err = jpeg.Decode(file); err != nil {
			return "", err
		}
	default:
		err = fmt.Errorf("images %s name not right", name)
		return "", err
	}
	resizeImg := resize.Resize(width, height, img, resize.Lanczos3)
	newName := newName(source, width, height)
	if outFile, err := os.Create(newName); err != nil {
		return "", err
	} else {
		defer outFile.Close()
		err = jpeg.Encode(outFile, resizeImg, nil)
		if err != nil {
			return "", err
		}
	}
	return newName, nil
}

// 重命名
func newName(name string, width, height uint) string {
	return fmt.Sprintf("%s_%d_%d%s", file.NoSuffixFileName(name), width, height, file.Filetype(name))
}

// OpenImage 打开图片
func openImage(path string) (image.Image, error) {
	isImage := IsImageFormat(path)
	if !isImage {
		return nil, errors.New(fmt.Sprintf("the specified file is not in image format: %s", path))
	}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(file)
	if err != nil {
		return nil, nil
	}
	return image, nil
}
