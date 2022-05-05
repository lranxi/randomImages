package images

import (
	"api/configs"
	"api/internal/model"
	"api/pkg/file"
	"api/pkg/img"
	"encoding/json"
	"fmt"
)

func (i *imageService) Image(width, height int) (*model.Image, error) {
	value, err := i.redis.SRandMember(configs.ValidImageRedisKey)
	if err != nil {
		return nil, err
	}
	image := &model.Image{}
	err = json.Unmarshal([]byte(value), image)
	if err != nil {
		return nil, err
	}

	// 根据要求的尺寸来读取图片，如果图片不存在，则需要重新生成
	newFilePath := fmt.Sprintf("%s_%d_%d%s", file.NoSuffixFileName(image.Path), width, height, file.Filetype(image.Path))
	_, exist := file.IsExists(newFilePath)
	if exist {
		image.Path = newFilePath
		return image, nil
	}
	// 文件不存在，重新生成
	if width <= 0 || height <= 0 {
		return image, nil
	}

	if image.OriginalWidth < width || image.OriginalHeight < height {
		return image, nil
	}
	newPath, err := img.Compression(image.Path, uint(width), uint(height))
	if err != nil {
		return nil, err
	}
	image.Path = newPath

	return image, nil
}
