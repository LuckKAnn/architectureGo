package logging

import (
	"fmt"
	"ginDemo/pkg/file"
	"ginDemo/pkg/setting"
	"log"
	"os"
	"time"
)

// 文件目录
// 日志拓展名
//

//var (
//	LogSavePath   = "runtime/logs/"
//	LogFileName   = "log"
//	LogFileExt    = "log"
//	LogTimeFormat = "20240517"
//)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filepath string) (*os.File, error) {
	// 这里这种是否算是一种通用的模式
	// 通过stat获取文件的描述符，之后判断文件是否创建，权限是否具备
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	src := dir + "/" + filepath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	// 文件打开的方式和对应的权限
	// 这里是以追加写的方式，打开或者创建一个文件
	file, err := file.Open(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Fail to OpenFile : %v", err)
	}
	return file, nil
}

func mkdir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
