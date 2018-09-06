package logfile

import (
	"fmt"
	"log"
	"os"
)

//默认 最后一次初始化的 日志系统，便于跨函数（无参数）使用
var GlobalLog *LogFileName


// Log data struct
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
//初始化一个日志系统，如果使用者不接收函数输出，可以使用默认最后一次初始化的 GlobalLog，并且后续禁用再次初始化，否则使用默认初始化（ GlobalLog ）会变化
func LogFileNew() (*LogFileName) {
	var l LogFileName
	//给个默认日志等级
	l.logLevel = L_DEBUG
	GlobalLog = &l
	return &l
}

// Setting the log level
func (l *LogFileName)SetLoglevel( level int) {
	l.logLevel = level
}

// Getting the log level
func (l *LogFileName)GetLoglevel() (level int) {
	return l.logLevel
}

// Initialization(Open) log . You need to provide a log path
func (l *LogFileName)LogFileOpen(patch string) ( err error) {
	l.openFile, err = os.OpenFile(patch, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("日志文件操作失败", err)
		os.Exit(1)
		return  err
	}
	l.debug = log.New(l.openFile, "[DEBUG]", log.Ldate|log.Ltime)
	l.info = log.New(l.openFile, "[INFO]", log.Ldate|log.Ltime)
	l.warn = log.New(l.openFile, "[WARN]", log.Ldate|log.Ltime|log.Llongfile)
	l.error = log.New(l.openFile, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
	l.fatal = log.New(l.openFile, "[FATAL]", log.Ldate|log.Ltime|log.Llongfile)

	return  err
}

//close log file
func (l *LogFileName)LogFileClosed() {
	l.openFile.Close()
}

//DEBUG log like Println
func (l *LogFileName)Debugln( a ...interface{}) {
	if l.logLevel >= L_DEBUG {
		l.debug.Output(2, fmt.Sprintln(a...))
	}
}
//DEBUG log like Printf
func (l *LogFileName)Debugf( format string, v ...interface{}) {
	if l.logLevel >= L_DEBUG {
		l.debug.Output(2, fmt.Sprintf(format, v...))
	}
}


//INFO log like Println
func (l *LogFileName)Infoln( a ...interface{}) {
	if l.logLevel >= L_INFO {
		l.info.Output(2, fmt.Sprintln(a...))
	}
}
//INFO log like like Printf
func (l *LogFileName)Infof( format string, v ...interface{}) {
	if l.logLevel >= L_INFO {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}


//WARN log like Println
func (l *LogFileName)Warnln( a ...interface{}) {
	if l.logLevel >= L_WARN {
		l.warn.Output(2, fmt.Sprintln(a...))
	}
}
//WARN log like like Printf
func (l *LogFileName)Warnf( format string, v ...interface{}) {
	if l.logLevel >= L_WARN {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}

//ERROR log like Println
func (l *LogFileName)Errorln( a ...interface{}) {
	if l.logLevel >= L_ERROR {
		l.error.Output(2, fmt.Sprintln(a...))
	}
}
//ERROR log like like Printf
func (l *LogFileName)Errof( format string, v ...interface{}) {
	if l.logLevel >= L_ERROR {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}

//FATAL log like Println
func (l *LogFileName)Fatalln( a ...interface{}) {
	if l.logLevel >= L_FATAL {
		l.fatal.Output(2, fmt.Sprintln(a...))
	}
}
//FATAL log like like Printf
func (l *LogFileName)Fatalf( format string, v ...interface{}) {
	if l.logLevel >= L_FATAL {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}










