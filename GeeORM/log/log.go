package log

import (
	"io"
	"log"
	"os"
	"sync"
)

//logger instance
var (
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{infoLog, errorLog}
	mux      sync.Mutex
)

//导出函数，提供log服务
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

//log level
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

func SetLogLevel(level int) {
	mux.Lock()
	defer func() { mux.Unlock() }()

	if level >= Disabled {
		errorLog.SetOutput(io.Discard)
		infoLog.SetOutput(io.Discard)
	} else if level >= ErrorLevel {
		infoLog.SetOutput(io.Discard)
	}
}
