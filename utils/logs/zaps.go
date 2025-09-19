package logs

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志级别
// Debug：调试信息（开发环境）
// Info：正常运行信息
// Warn：警告信息
// Error：错误信息
// DPanic：开发环境 panic（生产环境转为 Error）
// Panic：记录日志后触发 panic
// Fatal：记录日志后调用 os.Exit (1)

var Log = zap.NewExample()

func InitLog(appPath, errPath string) (err error) {
	// 日志级别
	level := zap.NewAtomicLevel()
	// 初始级别debug
	level.SetLevel(zapcore.DebugLevel)

	// 时间格式编码器
	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	// 日志配置
	config := zap.Config{
		Level:       level,
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			MessageKey:     "msg",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeTime:     timeEncoder,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		// 标准输出，应用日志文件
		OutputPaths: []string{
			"stdout",
			appPath,
		},
		// 标准错误输出，错误日志文件
		ErrorOutputPaths: []string{
			"stderr",
			errPath,
		},
	}

	// 根据配置生成log
	Log, err = config.Build()
	if err != nil {
		Log = zap.NewExample()
		return err
	}
	defer Log.Sync()
	return nil
}
