package logger

import "go.uber.org/zap"

func NewZapProduction() (AppLogger, error) {
	return zap.NewProduction()
}

func NewZapDevelopment() (AppLogger, error) {
	return zap.NewDevelopment()
}
