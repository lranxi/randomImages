package images

import (
	"api/configs"
	"api/internal/model"
	"encoding/json"
)

func (i *imageService) UncheckList() ([]*model.Image, error) {
	kvPair, err := i.redis.HGetAll(configs.UncheckImageRedisKey)
	if err != nil {
		return nil, err
	}
	var images []*model.Image
	for _, v := range kvPair {
		img := &model.Image{}
		json.Unmarshal([]byte(v), img)
		images = append(images, img)
	}
	return images, nil
}
