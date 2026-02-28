// Package logging は、ロギング機能の実装またはインターフェースを提供します。
package logging

import (
	"time"
)

type AccessLogEntry struct {
	Status    int
	Method    string
	Path      string
	IP        string
	Latency   time.Duration
	UserAgent string
	Time      time.Time
}

// Logging Loggingのインターフェース
type Logging interface {
	Close()
	Error(string)
	Warning(string)
	Info(string)
	Debug(string)
	AccessLog(*AccessLogEntry)
	SQLLog(string, string, string)
}
