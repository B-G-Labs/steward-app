package logging

import (
	"go.uber.org/zap"
)

func InitLogger() {
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))

}
