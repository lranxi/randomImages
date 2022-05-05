package images

import (
	"api/configs"
	"api/internal/model"
	"encoding/json"
)

func (i *imageService) List() ([]*model.Image, error) {
	values, err := i.redis.SMembers(configs.ValidImageRedisKey)
	if err != nil {
		return nil, err
	}
	// 转为切片返回
	var images []*model.Image
	for _, v := range values {
		img := &model.Image{}
		json.Unmarshal([]byte(v), img)
		images = append(images, img)
	}
	return images, nil
}
