package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var(
	//各个等级日志输出信息
	debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
	warnLog = log.New(os.Stdout, "[WARN] ", log.LstdFlags|log.Lshortfile)
	errorLog = log.New(os.Stdout, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{debugLog, infoLog, warnLog, errorLog}
	mu sync.Mutex
)

var(
	//日志打印函数
	DEBUG = debugLog.Println
	DEBUGF = debugLog.Printf
	INFO = infoLog.Println
	INFOF = infoLog.Printf
	WARN = warnLog.Println
	WARNF = warnLog.Printf
	ERROR = errorLog.Println
	ERRORF = errorLog.Printf
)


//日志等级
type logLevel int
const(
	DebugLevel logLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	Disabled
)


func SetLevel(level logLevel){
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers{
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel > level{
		errorLog.SetOutput(ioutil.Discard)
	}
	if WarnLevel > level{
		warnLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel > level{
		infoLog.SetOutput(ioutil.Discard)
	}
	if DebugLevel > level{
		debugLog.SetOutput(ioutil.Discard)
	}
}