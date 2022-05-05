package file

import (
	"os"
	"strings"
)

// DirMustExist 目录不存在则创建
func DirMustExist(path string) error {
	_, newPathExist := IsExists(path)
	if !newPathExist {
		err := os.Mkdir(path, 0766)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsExists 文件或目录是否存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

// Filetype 取文件后缀名
func Filetype(filename string) string {
	fileDot := strings.Index(filename, ".")
	fileType := filename[fileDot:]
	return fileType
}

// NoSuffixFileName 无后缀的文件名称
func NoSuffixFileName(filename string) string {
	fileDot := strings.Index(filename, ".")
	name := filename[:fileDot]
	return name
}
