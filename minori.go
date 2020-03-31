package minori

import (
	"fmt"
	"io"
	"os"
	"time"
	"strings"

	"github.com/mattn/go-colorable"
)

const (
	OFF int = iota
	FATAL
	PANIC
	ERROR
	WARN
	INFO
	DEBUG
)

var Level = DEBUG
var Out = colorable.NewColorableStdout()

type Logger struct {
	Name  string
	Out   io.Writer
	Level int
}

func GetLogger(name string) *Logger {
	return &Logger{Name: name, Out: nil, Level: -1}
}

func (l *Logger) GetLogger(name string) *Logger {
	return &Logger{
		Name:  l.Name + "/" + name,
		Out:   l.Out,
		Level: l.Level,
	}
}

func (l *Logger) log(lvl int, from, message string) {
	level := l.Level
	if level == -1 {
		level = Level
	}

	if lvl > level {
		return
	}

	out := l.Out
	if out == nil {
		out = Out
	}

	for _, msg := range strings.Split(message, "\n") {
		if strings.Trim(msg, " ") == "" {
			continue
		}
		fmt.Fprintf(out, "%s | \x1b[%dm%s\x1b[0m | \x1b[35m%s\x1b[0m | \x1b[94m%s\x1b[0m | %s\n",
			time.Now().Format("2006 01 02 | 15 04 05"),
			getColorByLevel(lvl), getMessageByLevel(lvl),
			l.Name, from, msg,
		)
	}
}

func (l *Logger) Panic(from string, v interface{}) {
	l.log(PANIC, from, fmt.Sprint(v))
	panic(v)
}

func (l *Logger) Panicf(from, format string, v ...interface{}) {
	p := fmt.Sprintf(format, v...)
	l.log(PANIC, from, p)
	panic(p)
}

func (l *Logger) Fatal(from string, v ...interface{}) {
	l.log(FATAL, from, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(from, format string, v ...interface{}) {
	l.log(FATAL, from, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Info(from string, v ...interface{}) {
	l.log(INFO, from, fmt.Sprint(v...))
}

func (l *Logger) Infof(from string, format string, v ...interface{}) {
	l.log(INFO, from, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(from string, v ...interface{}) {
	l.log(ERROR, from, fmt.Sprint(v...))
}

func (l *Logger) Errorf(from string, format string, v ...interface{}) {
	l.log(ERROR, from, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(from string, v ...interface{}) {
	l.log(DEBUG, from, fmt.Sprint(v...))
}

func (l *Logger) Debugf(from string, format string, v ...interface{}) {
	l.log(DEBUG, from, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(from string, v ...interface{}) {
	l.log(WARN, from, fmt.Sprint(v...))
}

func (l *Logger) Warnf(from string, format string, v ...interface{}) {
	l.log(WARN, from, fmt.Sprintf(format, v...))
}

func getMessageByLevel(level int) string {
	switch level {
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	default:
		return ""
	}
}

func getColorByLevel(level int) int {
	switch level {
	case DEBUG:
		// cyan
		return 36
	case WARN:
		// yellow
		return 33
	case ERROR, FATAL, PANIC:
		// red
		return 31
	default:
		// green
		return 32
	}
}
