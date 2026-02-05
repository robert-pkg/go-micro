package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"go-micro.dev/v5/logger"

	zapLog "github.com/robert-pkg/go-micro/my-examples/hello-world/common/log/zap"
)

func Init(logFileName string, output2Console bool) error {
	zapConfig := zap.NewProductionConfig()

	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	zapConfig.Encoding = "json"
	zapConfig.EncoderConfig.TimeKey = "t"
	zapConfig.EncoderConfig.LevelKey = "l"
	zapConfig.EncoderConfig.NameKey = "logger"
	zapConfig.EncoderConfig.CallerKey = "c"
	zapConfig.EncoderConfig.MessageKey = "msg"
	zapConfig.EncoderConfig.StacktraceKey = "st"
	zapConfig.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	zapConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	zapConfig.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder //zapcore.ShortCallerEncoder

	coreList := make([]zapcore.Core, 0, 2)
	if output2Console {
		coreList = append(coreList, zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig.EncoderConfig), zapcore.Lock(os.Stdout), zapConfig.Level))
	}

	if len(logFileName) > 0 {
		var enc zapcore.Encoder
		if zapConfig.Encoding == "json" {
			enc = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
		} else {
			enc = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
		}

		hook := &lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    100, // megabytes
			MaxBackups: 10,
			MaxAge:     10, // days
			LocalTime:  true,
			Compress:   true,
		}

		coreList = append(coreList, zapcore.NewCore(enc, zapcore.AddSync(hook), zapConfig.Level))
	}

	zaplog := zap.New(zapcore.NewTee(coreList...),
		zap.WithCaller(true),
		zap.AddCallerSkip(2),
		zap.AddStacktrace(zapcore.ErrorLevel)) // error级别及以上输出调用栈

	log, err := zapLog.NewLogger(
		logger.WithLevel(logger.DebugLevel),
		zapLog.WithLogger(zaplog),
	)
	if err != nil {
		return err
	}

	logger.DefaultLogger = log
	return nil
}
