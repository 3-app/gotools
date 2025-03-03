package dirtool

import (
	"log"
	"os"
	"strings"
)

// PathExist ， 判断文件是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MustCreateDir , 创建文件夹，不返回错误
func MustCreateDir(path string) {
	exist, err := PathExist(path)
	if err != nil {
		log.Fatalln(path, err)
	}
	if !exist {
		os.MkdirAll(path, 0777)
	}
}

// GetBasePath ，获取项目的根目录，带 "/"
func GetBasePath() string {
	pwd, _ := os.Getwd()
	return pwd + string(os.PathSeparator)
}

// GetParentDirectory 获取上层目录
func GetParentDirectory(dirctory string) string {
	return dirctory[0:strings.LastIndex(dirctory, string(os.PathSeparator))]
}

// GetConfigPath ，获取项目的配置目录
func GetConfigPath() string {
	return GetBasePath() + "config" + string(os.PathSeparator)
}
