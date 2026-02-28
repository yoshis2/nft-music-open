// Package logging は、ロギング機能の実装またはインターフェースを提供します。
package logging

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockLoggingはLoggingインターフェースのモック実装です
type MockLogging struct {
	entries []AccessLogEntry
}

func (m *MockLogging) Close() {}

func (m *MockLogging) Error(msg string) {
	m.entries = append(m.entries, AccessLogEntry{Time: time.Now(), Path: msg})
}

func (m *MockLogging) Warning(msg string) {
	// 警告ログの処理
}

func (m *MockLogging) Info(msg string) {
	// 情報ログの処理
}

func (m *MockLogging) Debug(msg string) {
	// デバッグログの処理
}

func (m *MockLogging) AccessLog(entry *AccessLogEntry) {
	m.entries = append(m.entries, *entry)
}

func (m *MockLogging) SQLLog(query, params, result string) {
	// SQLログの処理
}

func TestAccessLogEntry(t *testing.T) {
	mockLogger := &MockLogging{}

	entry := &AccessLogEntry{
		Status:    200,
		Method:    "GET",
		Path:      "/api/test",
		IP:        "192.168.1.1",
		Latency:   100 * time.Millisecond,
		UserAgent: "Mozilla/5.0",
		Time:      time.Now(),
	}

	// AccessLogメソッドを呼び出す
	mockLogger.AccessLog(entry)

	// ログエントリが正しく追加されたか確認
	assert.Equal(t, 1, len(mockLogger.entries))
	assert.Equal(t, entry.Status, mockLogger.entries[0].Status)
	assert.Equal(t, entry.Method, mockLogger.entries[0].Method)
	assert.Equal(t, entry.Path, mockLogger.entries[0].Path)
	assert.Equal(t, entry.IP, mockLogger.entries[0].IP)
	assert.Equal(t, entry.Latency, mockLogger.entries[0].Latency)
	assert.Equal(t, entry.UserAgent, mockLogger.entries[0].UserAgent)
}
