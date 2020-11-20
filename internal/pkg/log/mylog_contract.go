package log

type Logger interface {
	Info(message string, field Field)
	Warn(message string, field Field)
	Debug(message string, field Field)
	Error(message string, field Field)
	DPanic(message string, field Field)
	Fatal(message string, field Field)
	Panic(message string, field Field)
	With(field Field)
	ToSugar()
}
