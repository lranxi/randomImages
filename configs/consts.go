package configs

const (
	// ProjectName 项目名称
	ProjectName = "random-images"

	// ProjectAccessLogFile 访问日志存放路径
	ProjectAccessLogFile = "./logs/access.log"

	// UncheckStaticFileDir 静态文件目录
	UncheckStaticFileDir = "static/images/uncheck/"

	// StaticFileDir 有效静态图片目录
	StaticFileDir = "static/images/valid/"

	// FileNameLength 文件名称长度
	FileNameLength = 16

	// ImageMaxSize 单张图片最大容量
	ImageMaxSize = 5 << 20

	// ImageMinWidth 图片最小宽度
	ImageMinWidth = 1024

	// ImageMaxWidth 图片大宽度
	ImageMaxWidth = 4096

	// ImageMinHeight 图片最小高度
	ImageMinHeight = 768

	// ImageMaxHeight 图片最大高度
	ImageMaxHeight = 2160

	// UncheckImageRedisKey 未检查图片的redis key
	UncheckImageRedisKey = "images:uncheck"

	// ValidImageRedisKey 安全图片redis缓存
	ValidImageRedisKey = "images:valid"
)
