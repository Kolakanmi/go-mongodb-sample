package log

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
)

type (
	Config struct {
		Level string `envconfig:"LOG_LEVEL"`
		FilePath string `envconfig:"LOG_FILE_PATH"`
	}
	Field map[string]interface{}

	contextKey string
)

const (
	loggerKey = contextKey("logger_key")
	filePrefix = "file://"
)

var root Logger

func NewLogger(conf Config){
	out := outputFromEnv(conf.FilePath)
	writeSync := zapcore.AddSync(out)
	//multipleWrite := zapcore.NewMultiWriteSyncer(writeSync)
	c := zap.NewProductionEncoderConfig()
	c.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	c.CallerKey = "caller"
	core := zapcore.NewCore(zapcore.NewJSONEncoder(c), writeSync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel), zap.WithCaller(true))
	logger.Info("string")
	globalLog = logger
}
func outputFromEnv(path string) io.WriteCloser {
	f, err := os.Create(path)
	if err != nil {
		log.Println("failed to create log file", err)
		return os.Stdout
	}
	return f
}

func Root() Logger {
	if root == nil {
		root = newMyLog()
	}
	return root
}

func Init(field Field) {
	root = newMyLogWithFields(field)
}

func NewWithPrefix(key string, value interface{}) Logger {
	return newMyLogWithField(key, value)
}

func NewContext(ctx context.Context, logger Logger) context.Context {
	if logger == nil {
		logger = Root()
	}
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		return Root()
	}
	if logger, ok := ctx.Value(loggerKey).(Logger); ok {
		return logger
	}
	return Root()
}

func LoadEnvConfig() *Config {
	var conf Config
	envconfig.Load(&conf)
	return &conf
}

func convertToZapField(field Field) []zap.Field{
	var zapFields []zap.Field
	for k, v := range field{
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return zapFields
}
