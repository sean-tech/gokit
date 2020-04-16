package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"github.com/robfig/cron"
	"github.com/sean-tech/gokit/foundation"
)

type LogConfig struct {
	RunMode			foundation.RUN_MODE
	RuntimeRootPath string
	LogSavePath 	string
	LogPrefix		string
}
var _config LogConfig

/**
 * setup
 */
func Setup(config LogConfig) {
	_config = config

	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	_file, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}
	_logger = log.New(_file, _defaultPrefix, log.LstdFlags)

	rotateTimingStart()
}

func Writer() io.Writer {
	return _logger.Writer()
}

var (
	_defaultPrefix      = ""
	_defaultCallerDepth = 2
	_levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	_logPrefix  = ""

	_file *os.File
	_logger     *log.Logger
	_lock       sync.RWMutex
	_rotating sync.WaitGroup
)

type Level int
const (
	LEVEL_DEBUG Level = iota
	LEVEL_INFO
	LEVEL_WARNING
	LEVEL_ERROR
	LEVEL_FATAL
)

func Debug(v ...interface{})  {
	setPrefix(LEVEL_DEBUG)
	_rotating.Wait()
	_logger.Print(v)
}

func Info(v ...interface{})  {
	setPrefix(LEVEL_INFO)
	_rotating.Wait()
	_logger.Print(v)
}

func Warning(v ...interface{})  {
	setPrefix(LEVEL_WARNING)
	_rotating.Wait()
	_logger.Print(v)
}

func Error(v ...interface{})  {
	setPrefix(LEVEL_ERROR)
	_rotating.Wait()
	_logger.Print(v)
}

func Fatal(v ...interface{})  {
	setPrefix(LEVEL_FATAL)
	_rotating.Wait()
	_logger.Print(v)
}

func setPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(_defaultCallerDepth)
	if ok {
		_logPrefix = fmt.Sprintf("[%s]:[%s:%d]", _levelFlags[level], filepath.Base(file), line)
	} else {
		_logPrefix = fmt.Sprintf("[%s]", _levelFlags[level])
	}
	_logger.SetPrefix(_logPrefix)
}

func rotateTimingStart() {
	c := cron.New()
	spec := "0 0 0 * * *"
	if err := c.AddFunc(spec, func() {
		if logFileShouldRotate() == true {
			_rotating.Add(1)
			logFileRotate()
			_rotating.Done()
		}

	}); err != nil {
		log.Fatal(err)
	}
	c.Start()
}