package log

import "go.uber.org/zap"

var globalLog *zap.Logger

type (
	myLog struct {
		logger *zap.Logger
	}
)

func newMyLog() *myLog {
	return &myLog{
		logger: globalLog,
	}
}

func newMyLogWithField(k string, v interface{}) *myLog {
	return &myLog{
		logger: globalLog.With(zap.Any(k, v)),
	}
}

func newMyLogWithFields(field Field) *myLog {
	return &myLog{
		logger: globalLog.With(convertToZapField(field)...),
	}
}

func (m myLog) Info(message string, field Field)  {
	m.logger.Info(message, convertToZapField(field)...)
}

func (m myLog) Warn(message string, field Field)  {
	m.logger.Warn(message, convertToZapField(field)...)
}

func (m myLog) Debug(message string, field Field)  {
	m.logger.Debug(message, convertToZapField(field)...)
}

func (m myLog) Error(message string, field Field)  {
	m.logger.Error(message, convertToZapField(field)...)
}

func (m myLog) DPanic(message string, field Field)  {
	m.logger.DPanic(message, convertToZapField(field)...)
}

func (m myLog) Fatal(message string, field Field)  {
	m.logger.Fatal(message, convertToZapField(field)...)
}

func (m myLog) Panic(message string, field Field)  {
	m.logger.Panic(message, convertToZapField(field)...)
}

func (m myLog) With(field Field)  {
	m.logger.With(convertToZapField(field)...)
}

func (m myLog) ToSugar()  {
	m.logger.Sugar()
}