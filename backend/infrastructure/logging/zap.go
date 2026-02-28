// Package logging は、ロギング機能の実装またはインターフェースを提供します。
package logging

import (
	"time"

	"nft-music/usecases/logging"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ZapLogging STDOUT,STDERRにStackdriverLogging形式のログを出力し、StackdriverLoggingに自動的に取り込んでもらう場合のlogger
type ZapLogging struct {
	Client *zap.SugaredLogger
}

// NewZapLogging New ZapLogging
func NewZapLogging() *ZapLogging {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	accessSync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/go/src/nft-music/logs/access.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     10,   // days
		Compress:   true, // disabled by default
	})

	errSync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/go/src/nft-music/logs/error.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     10,
		Compress:   true,
	})

	accessCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		accessSync,
		zapcore.DebugLevel,
	)
	errorCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		errSync,
		zapcore.ErrorLevel,
	)

	opts := []zap.Option{}
	opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewSamplerWithOptions(
			core,
			time.Second,
			100, // Initial
			100, // Thereafter,
		)
	}))
	logger := zap.New(zapcore.NewTee(accessCore, errorCore), opts...)
	// logger, _ := zap.NewProduction()

	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()
	sugar := logger.Sugar()

	return &ZapLogging{
		Client: sugar,
	}
}

// Close 特になにもしていない
func (zap *ZapLogging) Close() {
}

// Error Errorレベルのアプリケーションログの出力
func (zap *ZapLogging) Error(data string) {
	zap.Client.Error(data)
}

// Warning Warningレベルのアプリケーションログの出力
func (zap *ZapLogging) Warning(data string) {
	zap.Client.Warn(data)
}

// Info Infoレベルのアプリケーションログの出力
func (zap *ZapLogging) Info(data string) {
	zap.Client.Info(data)
}

// Debug Debugレベルのアプリケーションログの出力
func (zap *ZapLogging) Debug(data string) {
	zap.Client.Debug(data)
}

// AccessLog echoのアクセスログの出力
func (zap *ZapLogging) AccessLog(data *logging.AccessLogEntry) {
	zap.Client.Infoln(data)
}

// SQLLog gormのSQLログの出力
func (zap *ZapLogging) SQLLog(v1 string, v2 string, v3 string) {
	zap.Client.Debug(v1)
	zap.Client.Debug(v2)
	zap.Client.Debug(v3)
}
