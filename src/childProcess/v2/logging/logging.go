package logging

import (
	//"fmt"
	"log"
	"os"
	"sync"
	"time"
)

/**********************************************************************/
type loggingManager struct {
	directory string
	date      string
	fileName  string
	fullPath  string
	file      *os.File

	mutex *sync.Mutex

	trace *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	debug *log.Logger
}

var loggingMgr *loggingManager

/**********************************************************************/
func (l *loggingManager) attachFp() {
	if l.file != nil {
		l.file.Close()
	}

	//fmt.Printf("[attach/fullPath] %s\n", l.fullPath)
	fp, err := os.OpenFile(l.fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	l.file = fp
	/* log.Llongfile || log.Lshortfile */
	l.debug = log.New(l.file, "[DEBUG] ", log.Ldate|log.Ltime|log.Lmicroseconds)
	l.trace = log.New(l.file, "[TRACE] ", log.Ldate|log.Ltime|log.Lmicroseconds)
	l.info = log.New(l.file, "[INFO] ", log.Ldate|log.Ltime|log.Lmicroseconds)
	l.warn = log.New(l.file, "[WARN]  ", log.Ldate|log.Ltime|log.Lmicroseconds)
	l.error = log.New(l.file, "[ERROR] ", log.Ldate|log.Ltime|log.Lmicroseconds)

}

func (l *loggingManager) getToday() (string, bool) {
	date := time.Now().Format("20060102")

	//fmt.Printf("[%s] [%s]\n", l.date, date)
	if l.date != date {
		return date, true
	}

	return date, false
}

func (l *loggingManager) prep() {
	if date, status := l.getToday(); status {
		defer l.mutex.Unlock()
		l.mutex.Lock()
		l.date = date

		l.fullPath = l.directory + "/" + date
		if err := os.MkdirAll(l.fullPath, 0755); err != nil {
			panic(err)
		}
		l.fullPath += "/" + l.fileName
		l.attachFp()
	}
}

/**********************************************************************/
//	traceF
//	traceLn
//	infoF
//	infoLn
//	warnF
//	warnLn
//	errorF
//	errorLn

func DebugF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.debug.Printf(format, v...)
}

func DebugLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.debug.Println(v...)
}

func TraceF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.trace.Printf(format, v...)
}

func TraceLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.trace.Println(v...)
}

func InfoF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.info.Printf(format, v...)
}

func InfoLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.info.Println(v...)
}

func WarnF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.warn.Printf(format, v...)
}

func WarnLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.warn.Println(v...)
}

func ErrorF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.error.Printf(format, v...)
}

func ErrorLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.error.Println(v...)
}

func Init(base string, fileName string) bool {
	if 0 >= len(base) || 0 >= len(fileName) {
		return false
	}

	loggingMgr = new(loggingManager)
	loggingMgr.mutex = &sync.Mutex{}

	directory := base
	date := time.Now().Format("20060102")
	fullPath := base + "/" + date

	//fmt.Println(fullPath)
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return false
	}

	loggingMgr.directory = directory
	loggingMgr.date = date
	loggingMgr.fileName = fileName
	loggingMgr.fullPath = fullPath + "/" + fileName
	//fmt.Println(loggingMgr.fullPath)
	//fmt.Printf("[full path] %s\n", loggingMgr.fullPath)

	loggingMgr.attachFp()

	return true
}

//func main() {
//	baseDir := os.Getenv("LOG_DIRECTORY")
//	fmt.Println(baseDir, "logfile")
//
//	if !Init(baseDir) {
//		return
//	}
//
//	traceLn("HI")
//	infoLn("HI")
//	warnLn("HI")
//	errorLn("HI")
//
//	test := 1
//	traceF("%s:%d\n", "TRACE TEST", test)
//	infoF("%s:%d\n", "INFO TEST", test)
//	warnF("%s:%d\n", "WARN TEST", test)
//	errorF("%s:%d\n", "ERROR TEST", test)
//}
