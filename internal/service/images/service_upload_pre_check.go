package images

import (
	"api/configs"
	"api/internal/consts"
	"api/internal/model"
	"api/pkg/file"
	"api/pkg/img"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// UploadPreCheck 上传图片预检查
func (i *imageService) UploadPreCheck(path string, category int) error {
	_, categoryExist := consts.ImageCategory[category]
	if !categoryExist {
		return errors.New("分类不正确")
	}

	_, ok := file.IsExists(path)
	if !ok {
		return errors.New("temp images not found")
	}
	isImage := img.IsImageFormat(path)
	if !isImage {
		return errors.New("file not is image format")
	}
	// 检查宽高
	width, height, err := img.CalculateWidthHeight(path)
	if err != nil {
		return err
	}
	if width < configs.ImageMinWidth || height < configs.ImageMinHeight || width > configs.ImageMaxHeight || height > configs.ImageMaxHeight {
		return errors.New(fmt.Sprintf("图片最小尺寸：%d * %d,最大尺寸：%d * %d", configs.ImageMinWidth, configs.ImageMinHeight, configs.ImageMaxWidth, configs.ImageMaxHeight))
	}

	pHash, err := img.ComputePHash(path)
	img := &model.Image{
		OriginalWidth:  width,
		OriginalHeight: height,
		Path:           path,
		Check:          false,
		Phash:          pHash,
		Category:       category,
	}

	// 检查phash，如果存在，则删除临时文件
	exist, _ := i.redis.HExist(configs.UncheckImageRedisKey, strconv.FormatInt(pHash, 10))
	if exist {
		os.Remove(path)
		return nil
	}

	// 图片名称存入redis
	err = i.redis.HSetNX(configs.UncheckImageRedisKey, strconv.FormatInt(pHash, 10), img)
	if err != nil {
		return errors.New("failed to save images")
	}

	return nil
}
