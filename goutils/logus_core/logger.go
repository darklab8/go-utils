package logus_core

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/darklab8/darklab_goutils/goutils/logus_core/logus_types"
)

func (l *Logger) Debug(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogGroup(opts...))
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogGroup(opts...))
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Info(msg, args...)
}

// Just potentially bad behavior to be aware of
func (l *Logger) Warn(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogGroup(opts...))
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Warn(msg, args...)
}

// It is bad but program can recover from it
func (l *Logger) Error(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogGroup(opts...))
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Error(msg, args...)
}

// Program is not allowed to run further with fatal
func (l *Logger) Fatal(msg string, opts ...SlogParam) {

	args := append([]any{}, newSlogGroup(opts...))
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Error(msg, args...)
	panic(msg)
}

func (l *Logger) CheckDebug(err error, msg string, opts ...SlogParam) bool {
	if err == nil {
		return false
	}
	args := append([]any{}, newSlogGroup(opts...))
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Debug(msg, args...)
	return true
}

func (l *Logger) CheckWarn(err error, msg string, opts ...SlogParam) bool {
	if err == nil {
		return false
	}
	args := append([]any{}, newSlogGroup(opts...))
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Warn(msg, args...)
	return true
}

func (l *Logger) CheckError(err error, msg string, opts ...SlogParam) bool {
	if err == nil {
		return false
	}
	args := append([]any{}, newSlogGroup(opts...))
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Error(msg, args...)
	return true
}

// It has shorter error output in comparison to CheckPanic
func (l *Logger) CheckFatal(err error, msg string, opts ...SlogParam) {
	if err == nil {
		return
	}
	args := append([]any{}, newSlogGroup(opts...))
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Error(msg, args...)
	os.Exit(1)
}

func (l *Logger) CheckPanic(err error, msg string, opts ...SlogParam) {
	if err == nil {
		return
	}
	args := append([]any{}, newSlogGroup(opts...))
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Error(msg, args...)
	panic(msg)
}

func (l *Logger) Debugf(msg string, varname string, value any, opts ...SlogParam) {
	args := append([]any{}, newSlogGroup(opts...))
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	args = append(args, varname)
	args = append(args, fmt.Sprintf("%v", value))
	l.logger.Debug(msg, args...)
}

func (l *Logger) Infof(msg string, varname string, value any, opts ...SlogParam) {
	args := append([]any{}, newSlogGroup(opts...))
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	args = append(args, varname)
	args = append(args, fmt.Sprintf("%v", value))
	l.logger.Info(msg, args...)
}

type Logger struct {
	logger              *slog.Logger
	enable_file_showing logus_types.EnableFileShowing
}

func NewLogger(
	log_level_str logus_types.LogLevel,
	enable_json_format logus_types.EnableJsonFormat,
	enable_file_showing logus_types.EnableFileShowing,
) *Logger {
	var programLevel = new(slog.LevelVar) // Info by default

	switch log_level_str {
	case LEVEL_DEBUG:
		programLevel.Set(slog.LevelDebug)
	case LEVEL_INFO:
		programLevel.Set(slog.LevelInfo)
	case LEVEL_WARN:
		programLevel.Set(slog.LevelWarn)
	case LEVEL_ERROR:
		programLevel.Set(slog.LevelError)
	}

	logger := &Logger{}

	if enable_json_format {
		logger.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))
	}
	logger.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))

	logger.enable_file_showing = enable_file_showing
	return logger
}

func GetCallingFile(level int) string {
	GetTwiceParentFunctionLocation := level
	_, filename, _, _ := runtime.Caller(GetTwiceParentFunctionLocation)
	filename = filepath.Base(filename)
	return fmt.Sprintf("f:%s ", filename)
}
