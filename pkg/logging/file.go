package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 文件目录
// 日志拓展名
//

var (
	LogSavePath   = "runtime/logs/"
	LogFileName   = "log"
	LogFileExt    = "log"
	LogTimeFormat = "20240517"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogFileName, time.Now().Format(LogTimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filepath string) *os.File {
	_, err := os.Stat(filepath)

	switch {
	case os.IsNotExist(err):
		mkdir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Fail to OpenFile : %v", err)
	}
	return file
}

func mkdir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
