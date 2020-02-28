package minori

import (
	"fmt"
	"io"
	"os"
	"time"

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
	return &Logger{Name: name, Out: Out, Level: -1}
}

func (l *Logger) GetLogger(name string) *Logger {
	return &Logger{
		Name:  l.Name + "/" + name,
		Out:   l.Out,
		Level: l.Level,
	}
}

func (l *Logger) log(lvl int, msg string) {
	level := l.Level
	if level == -1 {
		level = Level
	}

	if lvl > level {
		return
	}

	ws := ""
	if lvl == INFO || lvl == WARN {
		ws = " "
	}

	fmt.Fprintf(l.Out, "[%s] \x1b[%dm[%s]%s\x1b[0m \x1b[35m[%s]\x1b[0m %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		getColorByLevel(lvl), getMessageByLevel(lvl), ws,
		l.Name, msg,
	)
}

func (l *Logger) Panic(v interface{}) {
	l.log(PANIC, fmt.Sprint(v))
	panic(v)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	p := fmt.Sprintf(format, v...)
	l.log(PANIC, p)
	panic(p)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.log(FATAL, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log(FATAL, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Info(v ...interface{}) {
	l.log(INFO, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(INFO, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.log(ERROR, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(ERROR, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.log(DEBUG, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(DEBUG, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.log(WARN, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log(WARN, fmt.Sprintf(format, v...))
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
