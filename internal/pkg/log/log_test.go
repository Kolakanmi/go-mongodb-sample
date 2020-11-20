package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	//file1
	f, err := os.Create("l.log")
	if err != nil {
		t.Error(err)
	}
	writeSync := zapcore.AddSync(f)
	//file2
	multipleWrite := zapcore.NewMultiWriteSyncer(writeSync, os.Stdout)
	c := zap.NewProductionEncoderConfig()
	c.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	c.CallerKey = "caller"
	core := zapcore.NewCore(zapcore.NewJSONEncoder(c), multipleWrite, zap.InfoLevel)
	logger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel), zap.WithCaller(true))
	logger.Sugar().Info("success")
	fmt.Println("Success")
}

func TestNewLogger2(t *testing.T) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"l.log",
	}
	l, err := cfg.Build()
	if err != nil {
		t.Error(err)
	}
	l.Error("Success")
}