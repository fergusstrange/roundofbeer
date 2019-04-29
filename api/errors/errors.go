package errors

import (
	"github.com/apex/log"
	"strings"
)

func LogFatal(err error, message ...string) {
	if err != nil {
		joinedMessage := strings.Join(message, "\n")
		log.WithError(err).Fatal(joinedMessage)
	}
}

func LogError(err error, message ...string) {
	if err != nil {
		joinedMessage := strings.Join(message, "\n")
		log.WithError(err).Error(joinedMessage)
	}
}
