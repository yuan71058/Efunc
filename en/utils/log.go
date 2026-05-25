package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logger 日志记录器结构体。
// 支持同时输出到控制台和文件，带时间戳和日志级别。
type Logger struct {
	file     *os.File
	prefix   string
	logLevel int
}

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// Log_New 创建一个新的日志记录器。
// 如果日志路径为空，默认写入程序同目录下的日志文件。
//
// 参数:
//   - logDir: 日志文件所在目录
//   - prefix: 日志前缀
//
// 返回:
//   - *Logger: 日志记录器实例
func Log_New(logDir string, prefix string) *Logger {
	if logDir == "" {
		logDir = Dir_GetRunDir()
	}
	Dir_Create(logDir)
	logPath := filepath.Join(logDir, time.Now().Format("2006-01-02")+".log")
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("日志文件创建失败: %v", err)
		return nil
	}

	return &Logger{
		file:     f,
		prefix:   prefix,
		logLevel: LogLevelDebug,
	}
}

// write 内部写入方法。
func (l *Logger) write(level string, format string, v ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	content := timestamp + " [" + level + "] " + l.prefix + " " + format
	if len(v) > 0 {
		content = timestamp + " [" + level + "] " + l.prefix + " " + format
		log.Printf(content, v...)
		if l.file != nil {
			l.file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " [" + level + "] " + l.prefix + " " + format + "\n")
			for _, val := range v {
				l.file.WriteString(ToString(val))
			}
			l.file.WriteString("\n")
		}
	} else {
		log.Println(content)
		if l.file != nil {
			l.file.WriteString(content + "\n")
		}
	}
}

// Debug 输出调试级别日志。
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.logLevel <= LogLevelDebug {
		l.write("DEBUG", format, v...)
	}
}

// Info 输出信息级别日志。
func (l *Logger) Info(format string, v ...interface{}) {
	if l.logLevel <= LogLevelInfo {
		l.write("INFO", format, v...)
	}
}

// Warn 输出警告级别日志。
func (l *Logger) Warn(format string, v ...interface{}) {
	if l.logLevel <= LogLevelWarn {
		l.write("WARN", format, v...)
	}
}

// Error 输出错误级别日志。
func (l *Logger) Error(format string, v ...interface{}) {
	if l.logLevel <= LogLevelError {
		l.write("ERROR", format, v...)
	}
}

// Close 关闭日志文件。
func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

// Log_Simple 简单的日志写入函数。
// 写入程序同目录下的日志文件，每行带时间戳。
//
// 参数:
//   - content: 要记录的日志文本
//   - logPath: 日志文件路径，为空时默认程序同目录下的"运行日志.txt"
func Log_Simple(content string, logPath string) {
	if logPath == "" {
		logPath = Dir_GetRunDir() + "/运行日志.txt"
	}
	Dir_Create(File_GetParentDir(logPath))

	if !File_Exists(logPath) {
		File_WriteAny(logPath, []byte{})
	}

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}
	defer f.Close()

	line := time.Now().Format("2006-01-02 15:04:05") + "   " + content + "\n"
	f.WriteString(line)
}

// IOWriter 实现 io.Writer 接口，用于将日志重定向到文件。
type IOWriter struct {
	file *os.File
}

func (w *IOWriter) Write(p []byte) (n int, err error) {
	return w.file.Write(p)
}

// NewIOWriter 创建 io.Writer 日志输出器。
//
// 参数:
//   - filePath: 目标文件路径
//
// 返回:
//   - *IOWriter: io.Writer 实例
func NewIOWriter(filePath string) *IOWriter {
	Dir_Create(File_GetParentDir(filePath))
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}
	return &IOWriter{file: f}
}

// Close 关闭 io.Writer。
func (w *IOWriter) Close() error {
	if w.file != nil {
		return w.file.Close()
	}
	return nil
}

// 确保 IOWriter 实现了 io.Writer 接口。
var _ io.Writer = (*IOWriter)(nil)