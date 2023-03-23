package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type (
	// appLogger implements interface in log package
	appLogger struct {
		prefix string
	}

	// IAppLogger is an interface represent methods available in app logger
	IAppLogger interface {
		Info(args ...any)
		Infof(format string, args ...any)
		Warn(args ...any)
		Warnf(format string, args ...any)
		Error(args ...any)
		Errorf(format string, args ...any)
	}
)

var (
	once           = &sync.Once{}
	logger         = &appLogger{}
	logrusInstance = &logrus.Logger{}
)

// Info is a logrus log message at level info on the standard logger
func (l *appLogger) Info(args ...any) {
	l.formatLog().Info(args...)
}

// Infof is a logrus log message at level info with formatted string logger
func (l *appLogger) Infof(format string, args ...any) {
	l.formatLog().Infof(format, args...)

}

// Warn is a logrus log message at warn info on the standard logger
func (l *appLogger) Warn(args ...any) {
	l.formatLog().Warn(args...)
}

// Warnf is a logrus log message at level warn with formatted string logger
func (l *appLogger) Warnf(format string, args ...any) {
	l.formatLog().Warnf(format, args...)
}

// Error is a logrus log message at level error on the standard logger
func (l *appLogger) Error(args ...any) {
	l.formatLog().Error(args...)
}

// Errorf is a logrus log message at level error with formatted string logger
func (l *appLogger) Errorf(format string, args ...any) {
	l.formatLog().Errorf(format, args...)
}

func (l *appLogger) formatLog() *logrus.Entry {
	var source string

	if pc, file, line, ok := runtime.Caller(2); ok {
		var funcName string

		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = fn.Name()
			if i := strings.LastIndex(funcName, "."); i != -1 {
				funcName = funcName[i+1:]
			}
		}

		source = fmt.Sprintf("%s:%v:%s()", path.Base(file), line, path.Base(funcName))
	}

	return logrusInstance.WithFields(logrus.Fields{
		"prefix": l.prefix,
		"source": source,
	})
}

func NewAppLogger(prefix string) IAppLogger {
	once.Do(func() {
		logrusInstance = logrus.New()
		logrusInstance.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})

		logger = &appLogger{prefix}
	})

	return logger
}
