package logging

import (
	"fmt"
	"github.com/sean-tech/gokit/foundation"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"sync"
	"github.com/robfig/cron"
)

type Level int
const (
	level_debug Level = iota
	level_info
	level_warning
	level_error
	level_fatal
)

const (
	__defaultPrefix      = ""
	__defaultCallerDepth = 2
)

var (
	_levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	_lock       sync.Mutex
	_logger     *log.Logger
	//_ginWriter  io.Writer
)

func Setup(config LogConfig) {
	_config = config
	initLogger()
	logFileSliceTiming()
}

func initLogger() {
	var err error
	file, err := openLogFile(getLogFileName(), getLogFilePath())
	if err != nil {
		log.Fatalln(err)
	}
	_lock.Lock()
	defer _lock.Unlock()
	_logger = log.New(file, __defaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{})  {
	if _config.RunMode == foundation.RUN_MODE_DEBUG {
		fmt.Println(v)
		setPrefix(level_debug)
		_logger.Print(v)
	}
}

func Info(v ...interface{})  {
	setPrefix(level_info)
	_logger.Print(v)
}

func Warning(v ...interface{})  {
	setPrefix(level_warning)
	_logger.Print(v)
}

func Error(v ...interface{})  {
	setPrefix(level_error)
	_logger.Print(v)
}

func Fatal(v ...interface{})  {
	setPrefix(level_fatal)
	_logger.Print(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(__defaultCallerDepth)
	var logPrefix string = ""
	if ok {
		logPrefix = fmt.Sprintf("[%s]:[%s:%d]", _levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", _levelFlags[level])
	}
	_logger.SetPrefix(logPrefix)
}

func logFileSliceTiming()  {
	c := cron.New()
	spec := "0 0 0 * * *"
	err := c.AddFunc(spec, func() {
		if fileTimePassDaySlice() {
			initLogger()
			if ginWriterCallback != nil {
				GinWriterGet(ginWriterCallback)
			}
		}

	})
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}


type GinWriterCallback func(writer io.Writer)
var ginWriterCallback GinWriterCallback = nil
/**
 * 提供gin日志文件writer回调
 */
func GinWriterGet(callback GinWriterCallback)  {
	if callback == nil {
		return
	}
	if &ginWriterCallback != &callback {
		ginWriterCallback = callback
	}

	//var err error
	//_ginWriter, err = openLogFile(getLogFileName(_levelFlags[level_gin]), getLogFilePath())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	ginWriterCallback(_logger.Writer())
}