package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

/**********************************************************************/
type loggingManager struct {
	directory string
	fileName  string
	fullPath  string
	file      *os.File

	mutex *sync.Mutex

	trace *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

var loggingMgr *loggingManager

/**********************************************************************/
func (l *loggingManager) attachFp() {
	if l.file != nil {
		l.file.Close()
	}

	fp, err := os.OpenFile(l.fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	l.file = fp
	l.trace = log.New(loggingMgr.file, "[TRACE]   ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	l.info = log.New(loggingMgr.file, "[INFO]    ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	l.warn = log.New(loggingMgr.file, "[WARNING] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	l.error = log.New(loggingMgr.file, "[ERROR]   ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func (l *loggingManager) getToday() (string, bool) {
	date := time.Now().Format("20060102")

	if l.directory != date {
		return date, true
	}

	return date, false
}

func (l *loggingManager) prep() {
	if date, status := l.getToday(); status {
		defer l.mutex.Unlock()
		l.mutex.Lock()
		l.directory = date
		l.fullPath = date + "/" + l.file
		l.attachFp()
	}
}

/**********************************************************************/
/*
traceF
traceLn
infoF
infoLn
warnF
warnLn
errorF
errorLn
*/

func traceF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.trace.Printf(format, v...)
}

func traceLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.trace.Println(v...)
}

func infoF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.info.Printf(format, v...)
}

func infoLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.info.Println(v...)
}

func warnF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.warn.Printf(format, v...)
}

func warnLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.warn.Println(v...)
}

func errorF(format string, v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.error.Printf(format, v...)
}

func errorLn(v ...interface{}) {
	loggingMgr.prep()
	loggingMgr.error.Println(v...)
}

func Init() {
	loggingMgr = new(loggingManager)
	loggingMgr.mutex = &sync.Mutex{}

	directory := time.Now().Format("20060102")
	file := "loggingTest"

	if err := os.MkdirAll(directory, 0755); err != nil {
		panic(err)
	}

	loggingMgr.directory = directory
	loggingMgr.fileName = file
	loggingMgr.fullPath = directory + "/" + file
	fmt.Println(loggingMgr.fullPath)
	fmt.Printf("%s", loggingMgr.fullPath)

	loggingMgr.attachFp()
}

func main() {
	Init()

	traceLn("HI")
	infoLn("HI")
	warnLn("HI")
	errorLn("HI")

	test := 1
	traceF("%s:%d\n", "TRACE TEST", test)
	infoF("%s:%d\n", "INFO TEST", test)
	warnF("%s:%d\n", "WARN TEST", test)
	errorF("%s:%d\n", "ERROR TEST", test)
}
