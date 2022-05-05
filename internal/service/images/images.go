package images

import (
	"api/internal/model"
	"api/internal/repository/cache"
)

type imageService struct {
	redis cache.Cache
}

type ImageService interface {
	UploadPreCheck(path string, category int) error
	UncheckList() ([]*model.Image, error)
	Audit(images []model.Image) error
	List() ([]*model.Image, error)
	Image(width, height int) (*model.Image, error)
}

func New(cache cache.Cache) ImageService {
	return &imageService{redis: cache}
}
