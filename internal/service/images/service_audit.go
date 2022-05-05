package images

import (
	"api/configs"
	"api/internal/model"
	"api/pkg/file"
	"encoding/json"
	"os"
	"path"
	"strconv"
)

func (i *imageService) Audit(images []model.Image) error {
	for _, auditResult := range images {
		value, _ := i.redis.HGet(configs.UncheckImageRedisKey, strconv.FormatInt(auditResult.Phash, 10))
		if value != "" {
			img := &model.Image{}
			json.Unmarshal([]byte(value), img)
			oldPath := img.Path
			// 存到set中
			if auditResult.Check {
				newFilePath := path.Join(configs.StaticFileDir, strconv.Itoa(auditResult.Category)) + "/" + strconv.FormatInt(img.Phash, 10) + file.Filetype(img.Path)

				img.Check = true
				img.Category = auditResult.Category
				img.Path = newFilePath
				// 存到redis
				i.redis.SAdd(configs.ValidImageRedisKey, img)

				// 移动文件到新的分类
				file.DirMustExist(path.Join(configs.StaticFileDir, strconv.Itoa(auditResult.Category)))
				os.Rename(oldPath, newFilePath)
				// 删除hash中的数据
				i.redis.HDel(configs.UncheckImageRedisKey, strconv.FormatInt(img.Phash, 10))
			}
			// 删除文件
			os.Remove(oldPath)
		}
	}
	return nil
}
