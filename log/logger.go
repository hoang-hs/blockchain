package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

const callerSkip = 2

type logger struct {
	zap *zap.SugaredLogger
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func NewLogger() (Logging, error) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeTime:   SyslogTimeEncoder,
		EncodeLevel:  CustomLevelEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writeSyncer := zapcore.AddSync(os.Stderr)
	level := zap.DebugLevel

	core := zapcore.NewCore(encoder, writeSyncer, level)

	return &logger{
		zap: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(callerSkip)).Sugar(),
	}, nil
}

func (l *logger) Info(args ...interface{}) {
	l.zap.Info(args...)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	l.zap.Infof(msg, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.zap.Error(args...)
}

func (l *logger) Errorf(msg string, args ...interface{}) {
	l.zap.Errorf(msg, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.zap.Fatal(args...)
}

func (l *logger) Fatalf(msgFormat string, args ...interface{}) {
	l.zap.Fatalf(msgFormat, args...)
}

func (l *logger) GetZap() *zap.SugaredLogger {
	return l.zap
}
