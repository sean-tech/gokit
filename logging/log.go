package logging

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/sean-tech/gokit/validate"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

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


type LogConfig struct {
	RunMode			string	`validate:"required,oneof=debug release"`
	RuntimeRootPath string	`validate:"required,gt=1"`
	LogSavePath 	string	`validate:"required,gt=1"`
	LogPrefix		string	`validate:"required,gt=1"`
}
var _config LogConfig

/**
 * setup
 */
func Setup(config LogConfig) {
	if err := validate.ValidateParameter(config); err != nil {
		log.Fatal(err)
	}
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

func Logger() *log.Logger {
	return _logger
}



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

func Warn(v ...interface{})  {
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



type FmtLogger struct {
}
var _formatlogger = &FmtLogger{}

func FormatLogger() *FmtLogger {
	return _formatlogger
}

func (this *FmtLogger) Debug(v ...interface{}) {
	Debug(v)
}

func (this *FmtLogger) Debugf(format string, v ...interface{}) {
	Debug(fmt.Sprintf(format, v))
}

func (this *FmtLogger) Info(v ...interface{}) {
	Info(v)
}

func (this *FmtLogger) Infof(format string, v ...interface{}) {
	Info(fmt.Sprintf(format, v))
}

func (this *FmtLogger) Warn(v ...interface{}) {
	Warn(v)
}

func (this *FmtLogger) Warnf(format string, v ...interface{}) {
	Warn(fmt.Sprintf(format, v))
}

func (this *FmtLogger) Error(v ...interface{}) {
	Error(v)
}

func (this *FmtLogger) Errorf(format string, v ...interface{}) {
	Error(fmt.Sprintf(format, v))
}

func (this *FmtLogger) Fatal(v ...interface{}) {
	Fatal(v)
}

func (this *FmtLogger) Fatalf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v))
}

func (this *FmtLogger) Panic(v ...interface{}) {
	Fatal(v)
}

func (this *FmtLogger) Panicf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v))
}



