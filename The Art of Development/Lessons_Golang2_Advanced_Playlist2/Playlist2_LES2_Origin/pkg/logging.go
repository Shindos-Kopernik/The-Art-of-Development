package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct { // своя структура т.к. логгер может поменяться.
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err := os.MkdirAll("logs", 0744)
	if err != nil {
		panic(err)
	}
	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard) // не чего ни куда не пиши

	l.AddHook(&writeHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	// wrt := io.MultiWriter(os.Stdout, allFile) TODO стандартный метод, нагружает систему
	// l.SetOutput(wrt)

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}

// info.log -- info, warning, error, critical
//debug.log -- info, warning, error, critical, debug
// trace.log -- info, warning, error, critical, debug, trace
//error.log -- error, critical

// kafka -- info, debug
// file -- error, trace
//stdout -- warning, error
