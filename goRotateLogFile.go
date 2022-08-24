package gorotatelogfile

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func encoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.MessageKey = "Message"
	config.TimeKey = "TS"
	config.LevelKey = "Level"

	return zapcore.NewJSONEncoder(config)
}

func LogNewCore(writer *lumberjack.Logger, level zapcore.LevelEnabler) zapcore.Core {
	syncer := zapcore.AddSync(writer)
	return zapcore.NewCore(encoder(), syncer, level)
}

func LogNewLogger(cores []zapcore.Core, opts ...zap.Option) *zap.Logger {

	return zap.New(zapcore.NewTee(cores...), opts...)
}

func LogNewGlobalLogger(cores []zapcore.Core, opts ...zap.Option) {
	zap.ReplaceGlobals(LogNewLogger(cores, opts...))
}

func LogInit(logFilePath string, maxSizeMB int, compress bool, maxAgeDay int, maxBackupFiles int) {
	LogNewGlobalLogger([]zapcore.Core{
		LogNewCore(&lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    maxSizeMB,      // if maxSizeMB == 0 , default 100
			Compress:   compress,       // true / false
			MaxAge:     maxAgeDay,      // if maxAgeDay == 0 , default no
			MaxBackups: maxBackupFiles, // if maxBackupFiles == 0 , default no
		}, zap.InfoLevel),
	})
}

type Logs struct {
	FixedFieldsString map[string]string
	FixedFieldsInt    map[string]int
}

func (l Logs) setFixedFields() (fixedFields []zap.Field) {
	fixedFields = make([]zap.Field, 0)
	for k, v := range l.FixedFieldsString {
		fixedFields = append(fixedFields, zap.String(k, v))
	}
	for k, v := range l.FixedFieldsInt {
		fixedFields = append(fixedFields, zap.Int(k, v))
	}
	return
}

func (l Logs) Info(message string) {
	fixedFields := l.setFixedFields()
	zap.L().Info(message, fixedFields...)
}

func (l Logs) Debug(message string) {
	fixedFields := l.setFixedFields()
	zap.L().Debug(message, fixedFields...)
}

func (l Logs) Error(message string) {
	fixedFields := l.setFixedFields()
	zap.L().Error(message, fixedFields...)
}

func (l Logs) Warn(message string) {
	fixedFields := l.setFixedFields()
	zap.L().Warn(message, fixedFields...)
}
