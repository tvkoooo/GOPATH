package ljlog

import (
	"fmt"
	"log"
	"os"
)

// Log data structure
type LogFileName struct {
	openFile *os.File
	logLevel int
	debug    *log.Logger
	info     *log.Logger
	warn     *log.Logger
	error    *log.Logger
	fatal    *log.Logger
}

// log variable
const (
	L_OFF   = 0
	L_FATAL = 1
	L_ERROR = 2
	L_WARN  = 3
	L_INFO  = 4
	L_DEBUG = 5
)

// Setting the log level
func SetLoglevel(logFile *LogFileName, level int) {
	(*logFile).logLevel = level
}

// Getting the log level
func GetLoglevel(logFile *LogFileName) (level int) {
	return (*logFile).logLevel
}

// Initialization(Open) log . You need to provide a log path
func LogFileOpen(patch string) (logFile *LogFileName, err error) {
	var logstr LogFileName
	logstr.openFile, err = os.OpenFile(patch, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("日志文件操作失败", err)
		return nil, err
	}
	logstr.debug = log.New(logstr.openFile, "[DEBUG]", log.Ldate|log.Ltime|log.Llongfile)
	logstr.info = log.New(logstr.openFile, "[INFO]", log.Ldate|log.Ltime|log.Llongfile)
	logstr.warn = log.New(logstr.openFile, "[WARN]", log.Ldate|log.Ltime|log.Llongfile)
	logstr.error = log.New(logstr.openFile, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
	logstr.fatal = log.New(logstr.openFile, "[FATAL]", log.Ldate|log.Ltime|log.Llongfile)
	logstr.logLevel = L_DEBUG
	logFile = &logstr
	return logFile, err
}

//close log file
func LogFileClosed(logFile *LogFileName) {
	(*logFile).openFile.Close()
}

//DEBUG log like Println
func Debugln(logFile *LogFileName, a ...interface{}) {
	if (*logFile).logLevel >= L_DEBUG {
		(*logFile).debug.Println(a)
	}
}
//DEBUG log like Printf
func Debugf(logFile *LogFileName, format string, v ...interface{}) {
	if (*logFile).logLevel >= L_DEBUG {
		(*logFile).debug.Printf(format, v)
	}
}


//INFO log like Println
func Infoln(logFile *LogFileName, a ...interface{}) {
	if (*logFile).logLevel >= L_INFO {
		(*logFile).info.Println(a)
	}
}
//INFO log like like Printf
func Infof(logFile *LogFileName, format string, v ...interface{}) {
	if (*logFile).logLevel >= L_INFO {
		(*logFile).info.Printf(format, v)
	}
}


//WARN log like Println
func Warnln(logFile *LogFileName, a ...interface{}) {
	if (*logFile).logLevel >= L_WARN {
		(*logFile).warn.Println(a)
	}
}
//WARN log like like Printf
func Warnf(logFile *LogFileName, format string, v ...interface{}) {
	if (*logFile).logLevel >= L_WARN {
		(*logFile).info.Printf(format, v)
	}
}

//ERROR log like Println
func Errorln(logFile *LogFileName, a ...interface{}) {
	if (*logFile).logLevel >= L_ERROR {
		(*logFile).error.Println(a)
	}
}
//ERROR log like like Printf
func Errof(logFile *LogFileName, format string, v ...interface{}) {
	if (*logFile).logLevel >= L_ERROR {
		(*logFile).info.Printf(format, v)
	}
}

//FATAL log like Println
func Fatalln(logFile *LogFileName, a ...interface{}) {
	if (*logFile).logLevel >= L_FATAL {
		(*logFile).fatal.Println(a)
	}
}
//FATAL log like like Printf
func Fatalf(logFile *LogFileName, format string, v ...interface{}) {
	if (*logFile).logLevel >= L_FATAL {
		(*logFile).info.Printf(format, v)
	}
}










