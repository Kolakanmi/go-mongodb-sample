package log

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
)

type Config struct {
	Level string `envconfig:"LOG_LEVEL"`
	FilePath string `envconfig:"LOG_FILE_PATH"`
}

type Field map[string]interface{}

func NewLogger(conf Config) *zap.Logger{
	out := outputFromEnv(conf.FilePath)
	writeSync := zapcore.AddSync(out)
	c := zap.NewProductionEncoderConfig()
	c.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	c.CallerKey = "caller"
	core := zapcore.NewCore(zapcore.NewJSONEncoder(c), writeSync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel), zap.WithCaller(true))
	logger.Info("string")
	return logger
}
func outputFromEnv(path string) io.WriteCloser {
	f, err := os.Create(path)
	if err != nil {
		log.Println("failed to create log file", err)
	}
	return f
}

func LoadEnvConfig() *Config {
	var conf Config
	envconfig.Load(&conf)
	return &conf
}

func WithField(field Field) []zap.Field{
	zapFields := []zap.Field{}
	for k, v := range field{
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return zapFields
}

/*func NewLogger() (*zap.Logger, error) {
  cfg := zap.NewProductionConfig()
  cfg.OutputPaths = []string{
    "/var/log/myproject/myproject.log",
  }
  return cfg.Build()
}

 */