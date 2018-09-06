package logfile

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"path"
)

//系统日志
var (
	SysLogPath *string
    FilePath string
    FileName string
	FileInstance string
)



//默认 最后一次初始化的 日志系统，便于跨函数（无参数）使用
var GlobalLog *LogFileName
// Log data struct
type LogFileName struct {
	openFile *os.File
	logLevel int
	debug    *log.Logger
	info     *log.Logger
	notice   *log.Logger
	warn     *log.Logger
	err      *log.Logger
	alert    *log.Logger
	fail     *log.Logger
}

// log variable
//#define	LOG_EMERG	0	/* system log is unusable */
//#define	LOG_FAIL	1	/* system is fail */
//#define	LOG_ALERT	2	/* action must be taken immediately */
//#define	LOG_ERR		3	/* error conditions */
//#define	LOG_WARNING	4	/* warning conditions */
//#define	LOG_NOTICE	5	/* normal but significant condition */
//#define	LOG_INFO	6	/* informational */
//#define	LOG_DEBUG	7	/* debug-level messages */
const (
	LOG_EMERG    = 0
	LOG_FAIL     = 1
	LOG_ALERT    = 2
	LOG_ERR      = 3
	LOG_WARNING  = 4
	LOG_NOTICE   = 5
	LOG_INFO     = 6
	LOG_DEBUG    = 7
)
//初始化一个日志系统，如果使用者不接收函数输出，可以使用默认最后一次初始化的 GlobalLog，并且后续禁用再次初始化，否则使用默认初始化（ GlobalLog ）会变化
func LogFileNew() (*LogFileName) {
	var l LogFileName
	//给个默认日志等级
	l.logLevel = LOG_DEBUG
	GlobalLog = &l
	return &l
}

// Setting the log level
func (l *LogFileName)SetLogLevel( level int) {
	l.logLevel = level
}

// Getting the log level
func (l *LogFileName)GetLogLevel() (level int) {
	return l.logLevel
}

// Initialization(Open) log . You need to provide a log path
func (l *LogFileName)LogFileOpen(patch string) ( err error) {
	l.openFile, err = os.OpenFile(patch, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		SystemLogPrintln("FAIL","日志文件操作失败", err)
		os.Exit(1)
		return  err
	}
	l.debug = log.New(l.openFile, "[DEBUG]", log.Ldate|log.Ltime)
	l.info = log.New(l.openFile, "[INFO]", log.Ldate|log.Ltime)
	l.notice = log.New(l.openFile, "[NOTICE]", log.Ldate|log.Ltime)
	l.warn = log.New(l.openFile, "[WARN]", log.Ldate|log.Ltime|log.Llongfile)
	l.err = log.New(l.openFile, "[ERROR]", log.Ldate|log.Ltime|log.Llongfile)
	l.alert = log.New(l.openFile, "[ALERT]", log.Ldate|log.Ltime|log.Llongfile)
	l.fail = log.New(l.openFile, "[FAIL]", log.Ldate|log.Ltime|log.Llongfile)
	return  err
}

//close log file
func (l *LogFileName)LogFileClosed() {
	l.openFile.Close()
}

//DEBUG log like Println
func (l *LogFileName)Debugln( a ...interface{}) {
	if l.logLevel >= LOG_DEBUG {
		l.debug.Output(2, fmt.Sprintln(a...))
	}
}
//DEBUG log like Printf
func (l *LogFileName)Debugf( format string, v ...interface{}) {
	if l.logLevel >= LOG_DEBUG {
		l.debug.Output(2, fmt.Sprintf(format, v...))
	}
}


//INFO log like Println
func (l *LogFileName)Infoln( a ...interface{}) {
	if l.logLevel >= LOG_INFO {
		l.info.Output(2, fmt.Sprintln(a...))
	}
}
//INFO log like like Printf
func (l *LogFileName)Infof( format string, v ...interface{}) {
	if l.logLevel >= LOG_INFO {
		l.info.Output(2, fmt.Sprintf(format, v...))
	}
}
//NOTICE log like Println
func (l *LogFileName)Noticeln( a ...interface{}) {
	if l.logLevel >= LOG_NOTICE {
		l.notice.Output(2, fmt.Sprintln(a...))
	}
}
//NOTICE log like like Printf
func (l *LogFileName)Noticef( format string, v ...interface{}) {
	if l.logLevel >= LOG_NOTICE {
		l.notice.Output(2, fmt.Sprintf(format, v...))
	}
}

//WARN log like Println
func (l *LogFileName)Warnln( a ...interface{}) {
	if l.logLevel >= LOG_WARNING {
		l.warn.Output(2, fmt.Sprintln(a...))
	}
}
//WARN log like like Printf
func (l *LogFileName)Warnf( format string, v ...interface{}) {
	if l.logLevel >= LOG_WARNING {
		l.warn.Output(2, fmt.Sprintf(format, v...))
	}
}

//ERROR log like Println
func (l *LogFileName)Errorln( a ...interface{}) {
	if l.logLevel >= LOG_ERR {
		l.err.Output(2, fmt.Sprintln(a...))
	}
}
//ERROR log like like Printf
func (l *LogFileName)Errorf( format string, v ...interface{}) {
	if l.logLevel >= LOG_ERR {
		l.err.Output(2, fmt.Sprintf(format, v...))
	}
}
//ERROR log like Println
func (l *LogFileName)Alterln( a ...interface{}) {
	if l.logLevel >= LOG_ALERT {
		l.alert.Output(2, fmt.Sprintln(a...))
	}
}
//ERROR log like like Printf
func (l *LogFileName)Alterf( format string, v ...interface{}) {
	if l.logLevel >= LOG_ALERT {
		l.alert.Output(2, fmt.Sprintf(format, v...))
	}
}


//FATAL log like Println
func (l *LogFileName)Fatalln( a ...interface{}) {
	if l.logLevel >= LOG_FAIL {
		l.fail.Output(2, fmt.Sprintln(a...))
	}
}
//FATAL log like like Printf
func (l *LogFileName)Fatalf( format string, v ...interface{}) {
	if l.logLevel >= LOG_FAIL {
		l.fail.Output(2, fmt.Sprintf(format, v...))
	}
}

func SystemLogSetDefaultPath()  {
	_,FilePath,_,_ = runtime.Caller(1)
	FileName = path.Base(FilePath)
	FileInstance = ""
	SysLogPath = &FileName
}
func SystemLogSetPath(str *string)  {
	_,FilePath,_,_ = runtime.Caller(1)
	FileName = path.Base(FilePath)
	FileInstance = ""
	SysLogPath = str
}

func SystemLogSetInstance(str string)  {
	FileInstance = str
}

func SystemLogPrintln(logLevel string,a ...interface{})() {
	openFile, err := os.OpenFile(*SysLogPath + ".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("日志文件操作失败", err)
		os.Exit(1)
	}
	var logText string
	if FileInstance == ""{
		logText = "[" + logLevel + "]"
	}else {
		logText ="[" + FileName + "_" + FileInstance + "_" + logLevel + "]"
	}
	if logLevel == "info" || logLevel == "INFO" || logLevel == "Info"{
		outLog := log.New(openFile, logText, log.Ldate|log.Ltime)
		outLog.Output(2, fmt.Sprintln(a...))
	}else {
		outLog := log.New(openFile, logText, log.Ldate|log.Ltime|log.Llongfile)
		outLog.Output(2, fmt.Sprintln(a...))
	}
	openFile.Close()
}
func SystemLogPrintf(logLevel string,format string, v ...interface{})() {
	openFile, err := os.OpenFile(*SysLogPath + ".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Println("日志文件操作失败", err)
		os.Exit(1)
	}
	var logText string
	if FileInstance == ""{
		logText = "[" + logLevel + "]"
	}else {
		logText ="[" + FileName + "_" + FileInstance + "_" + logLevel + "]"
	}
	if logLevel == "info" || logLevel == "INFO" || logLevel == "Info"{
		outLog := log.New(openFile, logText, log.Ldate|log.Ltime)
		outLog.Output(2, fmt.Sprintf(format, v...))
	}else {
		outLog := log.New(openFile, logText, log.Ldate|log.Ltime|log.Llongfile)
		outLog.Output(2, fmt.Sprintf(format, v...))
	}
	openFile.Close()
}

