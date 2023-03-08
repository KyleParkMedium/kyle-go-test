package flogging

import (
	"go.uber.org/zap"
)

var Global *Logging

// MustGetLogger creates a logger with the specified name. If an invalid name
// is provided, the operation will panic.
func MustGetLogger(loggerName string) *FabricLogger {

	return Global.Logger(loggerName)

	return NewZapLogger(core).Named(name)

}

var log *zap.Logger

func init() {
	var err error

	log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
