package model

import "encoding/json"

// Image 图片
type Image struct {
	Phash          int64  `json:"phash"`          // 图片的感知哈希
	OriginalWidth  int    `json:"originalWidth"`  // 原始宽度
	OriginalHeight int    `json:"OriginalHeight"` // 原始高度
	Category       int    `json:"category"`       // 分类
	Path           string `json:"path"`           // 路径
	Check          bool   `json:"check"`          // 是否检查
}

func (i *Image) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Image) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}
