package typelog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func (l *Logger) Debug(msg string, opts ...LogType) {
	if IsMsgEnabled(l.level_log, LEVEL_DEBUG) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles())
		}
		l.logger.Debug(msg, args...)
	}
}

func (l *Logger) DebugCtx(ctx context.Context, msg string, opts ...LogType) {
	l.Debug(msg, CtxOpts(ctx, opts...)...)
}

func (l *Logger) Debugln(a ...any) {
	l.Debug(fmt.Sprintln(a...))
}

func (l *Logger) Info(msg string, opts ...LogType) {
	MetricMsgsTotal.WithLabelValues("info", l.GetFile(), strconv.FormatBool(IsMsgEnabled(l.level_log, LEVEL_INFO))).Inc()

	if IsMsgEnabled(l.level_log, LEVEL_INFO) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles())
		}
		l.logger.Info(msg, args...)
	}
}

func (l *Logger) Infoln(a ...any) {
	SetCallerSkip(4)
	l.Info(fmt.Sprintln(a...))
	UnsetCallerSkip()
}

func (l *Logger) InfoCtx(ctx context.Context, msg string, opts ...LogType) {
	SetCallerSkip(4)
	l.Info(msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
}

// Warn is for just potentially bad behavior to be aware of
func (l *Logger) Warn(msg string, opts ...LogType) {
	MetricMsgsTotal.WithLabelValues("warn", l.GetFile(), strconv.FormatBool(IsMsgEnabled(l.level_log, LEVEL_WARN))).Inc()

	if IsMsgEnabled(l.level_log, LEVEL_WARN) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles())
		}
		l.logger.Warn(msg, args...)
	}
}

func (l *Logger) WarnCtx(ctx context.Context, msg string, opts ...LogType) {
	SetCallerSkip(4)
	l.Warn(msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
}

func (l *Logger) Warnln(a ...any) {
	SetCallerSkip(4)
	l.Warn(fmt.Sprintln(a...))
	UnsetCallerSkip()
}

// Error is bad but program can recover from it
func (l *Logger) Error(msg string, opts ...LogType) {
	MetricMsgsTotal.WithLabelValues("error", l.GetFile(), strconv.FormatBool(IsMsgEnabled(l.level_log, LEVEL_ERROR))).Inc()

	if IsMsgEnabled(l.level_log, LEVEL_ERROR) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles())
		}
		l.logger.Error(msg, args...)
	}

}

func (l *Logger) Errorln(a ...any) {
	SetCallerSkip(4)
	l.Error(fmt.Sprintln(a...))
	UnsetCallerSkip()
}

func (l *Logger) ErrorCtx(ctx context.Context, msg string, opts ...LogType) {
	SetCallerSkip(4)
	l.Error(msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
}

// Fatal when encountered, Program is not allowed to run further with fatal
func (l *Logger) Fatal(msg string, opts ...LogType) {
	MetricMsgsTotal.WithLabelValues("panic", l.GetFile(), strconv.FormatBool(true)).Inc()

	if l.enable_scope_shown {
		opts = append(opts, String("scope", l.scope))
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	if l.enable_file_shown {
		args = append(args, logGroupFiles())
	}
	l.logger.Error(msg, args...)

	os.Exit(1)
}

func (l *Logger) FatalCtx(ctx context.Context, msg string, opts ...LogType) {
	SetCallerSkip(4)
	l.Fatal(msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
}
func (l *Logger) Fatalln(a ...any) {
	SetCallerSkip(4)
	l.Fatal(fmt.Sprintln(a...))
	UnsetCallerSkip()
}

func (l *Logger) Panic(msg string, opts ...LogType) {
	MetricMsgsTotal.WithLabelValues("panic", l.GetFile(), strconv.FormatBool(true)).Inc()

	if l.enable_scope_shown {
		opts = append(opts, String("scope", l.scope))
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	if l.enable_file_shown {
		args = append(args, logGroupFiles())
	}
	l.logger.Error(msg, args...)
	panic(msg)
}

func (l *Logger) PanicCtx(ctx context.Context, msg string, opts ...LogType) {
	SetCallerSkip(4)
	l.Panic(msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
}
func (l *Logger) Panicln(a ...any) {
	SetCallerSkip(4)
	l.Panic(fmt.Sprintln(a...))
	UnsetCallerSkip()
}
func (l *Logger) CheckDebug(err error, msg string, opts ...LogType) bool {
	if err == nil {
		return false
	}
	if !IsMsgEnabled(l.level_log, LEVEL_DEBUG) {
		return true
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	args = append(args, slog.String("error", fmt.Sprintf("%v", err)))

	l.logger.Debug(msg, args...)
	return true
}

func (l *Logger) CheckDebugCtx(ctx context.Context, err error, msg string, opts ...LogType) bool {
	return l.CheckDebug(err, msg, CtxOpts(ctx, opts...)...)
}
func (l *Logger) CheckDebugln(err error, a ...any) bool {
	return l.CheckDebug(err, fmt.Sprintln(a...))
}

func (l *Logger) CheckWarn(err error, msg string, opts ...LogType) bool {
	if err == nil {
		return false
	}
	if !IsMsgEnabled(l.level_log, LEVEL_WARN) {
		return true
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	args = append(args, slog.String("error", fmt.Sprintf("%v", err)))
	SetCallerSkip(4)
	l.logger.Warn(msg, args...)
	UnsetCallerSkip()
	return true
}

func (l *Logger) CheckWarnCtx(ctx context.Context, err error, msg string, opts ...LogType) bool {
	SetCallerSkip(5)
	result := l.CheckWarn(err, msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
	return result
}

func (l *Logger) CheckWarnln(err error, a ...any) bool {
	SetCallerSkip(5)
	result := l.CheckWarn(err, fmt.Sprintln(a...))
	UnsetCallerSkip()
	return result
}

func (l *Logger) CheckError(err error, msg string, opts ...LogType) bool {
	if err == nil {
		return false
	}
	if !IsMsgEnabled(l.level_log, LEVEL_ERROR) {
		return true
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	args = append(args, slog.String("error", fmt.Sprintf("%v", err)))
	SetCallerSkip(4)
	l.logger.Error(msg, args...)
	UnsetCallerSkip()
	return true
}

func (l *Logger) CheckErrorCtx(ctx context.Context, err error, msg string, opts ...LogType) bool {
	SetCallerSkip(5)
	result := l.CheckError(err, msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
	return result
}

func (l *Logger) CheckErrorln(err error, a ...any) bool {
	SetCallerSkip(5)
	result := l.CheckError(err, fmt.Sprintln(a...))
	UnsetCallerSkip()
	return result
}

// CheckFatal has shorter error output in comparison to CheckPanic
func (l *Logger) CheckFatal(err error, msg string, opts ...LogType) {
	if err == nil {
		return
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	args = append(args, slog.String("error", fmt.Sprintf("%v", err)))
	SetCallerSkip(4)
	l.logger.Error(msg, args...)
	UnsetCallerSkip()
	os.Exit(1)
}

func (l *Logger) CheckFatalCtx(ctx context.Context, err error, msg string, opts ...LogType) {
	SetCallerSkip(5)
	l.CheckFatal(err, msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
}

func (l *Logger) CheckFatalln(err error, a ...any) {
	SetCallerSkip(5)
	l.CheckFatal(err, fmt.Sprintln(a...))
	UnsetCallerSkip()
}

func (l *Logger) CheckPanic(err error, msg string, opts ...LogType) {
	if err == nil {
		return
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	args = append(args,
		slog.String("err_msg", fmt.Sprintf("%v", err)),
		slog.String("err_type", fmt.Sprintf("%T", err)),
	)
	SetCallerSkip(4)
	l.logger.Error(msg, args...)
	UnsetCallerSkip()
	panic_logger.Error(msg, args...)
	panic(panic_str.String())
}

func (l *Logger) CheckPanicCtx(ctx context.Context, err error, msg string, opts ...LogType) {
	SetCallerSkip(5)
	l.CheckPanic(err, msg, CtxOpts(ctx, opts...)...)
	UnsetCallerSkip()
}
func (l *Logger) CheckPanicln(err error, a ...any) {
	SetCallerSkip(5)
	l.CheckPanic(err, fmt.Sprintln(a...))
	UnsetCallerSkip()
}
