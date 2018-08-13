package logfile

import (
	"fmt"
	"log"
	"os"
)

var GlobalLog *LogFileName

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

func LogFileInit()  {
	var l LogFileName
	GlobalLog = &l
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
		return  err
	}
	l.debug = log.New(l.openFile, "[DEBUG]", log.Ldate|log.Ltime)
	l.info = log.New(l.openFile, "[INFO]", log.Ldate|log.Ltime)
	l.warn = log.New(l.openFile, "[WARN]", log.Ldate|log.Ltime|log.Llongfile)
	l.error = log.New(l.openFile, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
	l.fatal = log.New(l.openFile, "[FATAL]", log.Ldate|log.Ltime|log.Llongfile)
	l.logLevel = L_DEBUG

	return  err
}

//close log file
func (l *LogFileName)LogFileClosed() {
	l.openFile.Close()
}

//DEBUG log like Println
func (l *LogFileName)Debugln( a ...interface{}) {
	if l.logLevel >= L_DEBUG {
		l.debug.Println(a)
	}
}
//DEBUG log like Printf
func (l *LogFileName)Debugf( format string, v ...interface{}) {
	if l.logLevel >= L_DEBUG {
		l.debug.Printf(format, v)
	}
}


//INFO log like Println
func (l *LogFileName)Infoln( a ...interface{}) {
	if l.logLevel >= L_INFO {
		l.info.Println(a)
	}
}
//INFO log like like Printf
func (l *LogFileName)Infof( format string, v ...interface{}) {
	if l.logLevel >= L_INFO {
		l.info.Printf(format, v)
	}
}


//WARN log like Println
func (l *LogFileName)Warnln( a ...interface{}) {
	if l.logLevel >= L_WARN {
		l.warn.Println(a)
	}
}
//WARN log like like Printf
func (l *LogFileName)Warnf( format string, v ...interface{}) {
	if l.logLevel >= L_WARN {
		l.info.Printf(format, v)
	}
}

//ERROR log like Println
func (l *LogFileName)Errorln( a ...interface{}) {
	if l.logLevel >= L_ERROR {
		l.error.Println(a)
	}
}
//ERROR log like like Printf
func (l *LogFileName)Errof( format string, v ...interface{}) {
	if l.logLevel >= L_ERROR {
		l.info.Printf(format, v)
	}
}

//FATAL log like Println
func (l *LogFileName)Fatalln( a ...interface{}) {
	if l.logLevel >= L_FATAL {
		l.fatal.Println(a)
	}
}
//FATAL log like like Printf
func (l *LogFileName)Fatalf( format string, v ...interface{}) {
	if l.logLevel >= L_FATAL {
		l.info.Printf(format, v)
	}
}










