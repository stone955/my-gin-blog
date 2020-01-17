package logging

import (
	"fmt"
	"github.com/stone955/my-gin-blog/pkg/file"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err := openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func getLogFilePath() string {
	return filepath.Join(setting.AppCfg.RuntimeRootPath, setting.AppCfg.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s_%s.%s",
		setting.AppCfg.LogSaveName,
		time.Now().Format(setting.AppCfg.TimeFormat),
		setting.AppCfg.LogFileExt)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := filepath.Join(dir, filePath)
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	if err := file.MkDirIfNotExist(src); err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(filepath.Join(src, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to open file :%v", err)
	}

	return f, nil
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, f, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(f), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
