package errors

import (
	"github.com/apex/log"
	"runtime/debug"
)

func LogFatal(err error) {
	if err != nil {
		logWithError(err).Fatal("Unexpected Fatal Error")
	}
}

func LogError(err error) {
	if err != nil {
		logWithError(err).Error("Unexpected Error")
	}
}

func logWithError(err error) *log.Entry {
	return log.WithError(err).
		WithField("stacktrace", string(debug.Stack()))
}
