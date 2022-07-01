package logger

import (

	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.SugaredLogger

func initialize() {
	//Encoder configs
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = TimestampEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//lumberjack configs
	lumberjackLogger := &lumberjack.Logger{
		Filename:   propertymanager.GetStringProperty("logging.filename"),
		MaxSize:    propertymanager.GetIntProperty("logging.maxSize", 1500),
		MaxBackups: propertymanager.GetIntProperty("logging.maxBackups", 60),
		MaxAge:     propertymanager.GetIntProperty("logging.maxAge", 2),
		Compress:   true,
		LocalTime:  true,
	}
	writeSyncer := zapcore.AddSync(lumberjackLogger)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCallerSkip(1), zap.AddCaller())
	log = logger.Sugar()
}

func TimestampEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func Debug(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Error(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Fatal(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}

func Warn(template string, args ...interface{}) {
	log.Warnf(template, args...)
}
