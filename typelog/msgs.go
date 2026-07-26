package typelog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func (l *Logger) DebugCtx(ctx context.Context, msg string, opts ...LogType) {
	opts = CtxOpts(ctx, opts...)
	if IsMsgEnabled(l.level_log, LEVEL_DEBUG) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles(ctx))
		}
		l.logger.Debug(msg, args...)
	}
}

func (l *Logger) Debug(msg string, opts ...LogType) {
	l.DebugCtx(context.Background(), msg, opts...)
}

func (l *Logger) Debugln(a ...any) {
	l.DebugCtx(context.Background(), fmt.Sprintln(a...))
}

func (l *Logger) InfoCtx(ctx context.Context, msg string, opts ...LogType) {
	opts = CtxOpts(ctx, opts...)
	MetricMsgsTotal.WithLabelValues("info", l.GetFileCtx(ctx), strconv.FormatBool(IsMsgEnabled(l.level_log, LEVEL_INFO))).Inc()

	if IsMsgEnabled(l.level_log, LEVEL_INFO) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles(ctx))
		}
		l.logger.Info(msg, args...)
	}
}

func (l *Logger) Info(msg string, opts ...LogType) {
	l.InfoCtx(WithCtxCallerSkip(context.Background(), 4), msg, opts...)
}

func (l *Logger) Infoln(a ...any) {
	l.InfoCtx(WithCtxCallerSkip(context.Background(), 4), fmt.Sprintln(a...))
}

// Warn is for just potentially bad behavior to be aware of
func (l *Logger) Warn(msg string, opts ...LogType) {
	l.WarnCtx(WithCtxCallerSkip(context.Background(), 4), msg, opts...)

}

func (l *Logger) WarnCtx(ctx context.Context, msg string, opts ...LogType) {
	opts = CtxOpts(ctx, opts...)
	MetricMsgsTotal.WithLabelValues("warn", l.GetFileCtx(ctx), strconv.FormatBool(IsMsgEnabled(l.level_log, LEVEL_WARN))).Inc()

	if IsMsgEnabled(l.level_log, LEVEL_WARN) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles(ctx))
		}
		l.logger.Warn(msg, args...)
	}
}

func (l *Logger) Warnln(a ...any) {
	l.WarnCtx(WithCtxCallerSkip(context.Background(), 4), fmt.Sprintln(a...))
}

// Error is bad but program can recover from it
func (l *Logger) Error(msg string, opts ...LogType) {
	l.ErrorCtx(WithCtxCallerSkip(context.Background(), 4), msg, opts...)
}

func (l *Logger) Errorln(a ...any) {
	l.ErrorCtx(WithCtxCallerSkip(context.Background(), 4), fmt.Sprintln(a...))
}

func (l *Logger) ErrorCtx(ctx context.Context, msg string, opts ...LogType) {
	opts = CtxOpts(ctx, opts...)
	MetricMsgsTotal.WithLabelValues("error", l.GetFileCtx(ctx), strconv.FormatBool(IsMsgEnabled(l.level_log, LEVEL_ERROR))).Inc()

	if IsMsgEnabled(l.level_log, LEVEL_ERROR) {
		if l.enable_scope_shown {
			opts = append(opts, String("scope", l.scope))
		}
		args := append([]SlogAttr{}, newSlogArgs(opts...)...)
		if l.enable_file_shown {
			args = append(args, logGroupFiles(ctx))
		}
		l.logger.Error(msg, args...)
	}
}

// Fatal when encountered, Program is not allowed to run further with fatal
func (l *Logger) Fatal(msg string, opts ...LogType) {
	l.FatalCtx(WithCtxCallerSkip(context.Background(), 4), msg, opts...)
}

func (l *Logger) FatalCtx(ctx context.Context, msg string, opts ...LogType) {
	opts = CtxOpts(ctx, opts...)
	MetricMsgsTotal.WithLabelValues("panic", l.GetFileCtx(ctx), strconv.FormatBool(true)).Inc()

	if l.enable_scope_shown {
		opts = append(opts, String("scope", l.scope))
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	if l.enable_file_shown {
		args = append(args, logGroupFiles(ctx))
	}
	l.logger.Error(msg, args...)

	os.Exit(1)
}
func (l *Logger) Fatalln(a ...any) {
	l.FatalCtx(WithCtxCallerSkip(context.Background(), 4), fmt.Sprintln(a...))
}

func (l *Logger) Panic(msg string, opts ...LogType) {
	l.PanicCtx(WithCtxCallerSkip(context.Background(), 4), msg, opts...)

}

func (l *Logger) PanicCtx(ctx context.Context, msg string, opts ...LogType) {
	opts = CtxOpts(ctx, opts...)
	MetricMsgsTotal.WithLabelValues("panic", l.GetFileCtx(ctx), strconv.FormatBool(true)).Inc()

	if l.enable_scope_shown {
		opts = append(opts, String("scope", l.scope))
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	if l.enable_file_shown {
		args = append(args, logGroupFiles(ctx))
	}
	l.logger.Error(msg, args...)
	panic(msg)
}
func (l *Logger) Panicln(a ...any) {
	l.PanicCtx(WithCtxCallerSkip(context.Background(), 4), fmt.Sprintln(a...))
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

func (l *Logger) CheckWarnCtx(ctx context.Context, err error, msg string, opts ...LogType) bool {
	if err == nil {
		return false
	}
	if !IsMsgEnabled(l.level_log, LEVEL_WARN) {
		return true
	}
	opts = append(opts, OptError(err))
	l.WarnCtx(WithCtxCallerSkip(ctx, 4), msg, opts...)
	return true
}

func (l *Logger) CheckWarn(err error, msg string, opts ...LogType) bool {
	result := l.CheckWarnCtx(WithCtxCallerSkip(context.Background(), 5), err, msg, opts...)
	return result
}

func (l *Logger) CheckWarnln(err error, a ...any) bool {
	result := l.CheckWarnCtx(WithCtxCallerSkip(context.Background(), 5), err, fmt.Sprintln(a...))
	return result
}

func (l *Logger) CheckErrorCtx(ctx context.Context, err error, msg string, opts ...LogType) bool {
	if err == nil {
		return false
	}
	if !IsMsgEnabled(l.level_log, LEVEL_ERROR) {
		return true
	}
	opts = append(opts, OptError(err))
	l.ErrorCtx(WithCtxCallerSkip(ctx, 4), msg, opts...)
	return true
}

func (l *Logger) CheckError(err error, msg string, opts ...LogType) bool {
	return l.CheckErrorCtx(WithCtxCallerSkip(context.Background(), 5), err, msg, opts...)
}

func (l *Logger) CheckErrorln(err error, a ...any) bool {
	return l.CheckErrorCtx(WithCtxCallerSkip(context.Background(), 5), err, fmt.Sprintln(a...))
}

// CheckFatal has shorter error output in comparison to CheckPanic
func (l *Logger) CheckFatalCtx(ctx context.Context, err error, msg string, opts ...LogType) {
	if err == nil {
		return
	}
	opts = append(opts, OptError(err))
	l.ErrorCtx(WithCtxCallerSkip(ctx, 4), msg, opts...)
	os.Exit(1)
}

func (l *Logger) CheckFatal(err error, msg string, opts ...LogType) {
	l.CheckFatalCtx(WithCtxCallerSkip(context.Background(), 5), err, msg, opts...)
}

func (l *Logger) CheckFatalln(err error, a ...any) {
	l.CheckFatalCtx(WithCtxCallerSkip(context.Background(), 5), err, fmt.Sprintln(a...))
}

func (l *Logger) CheckPanicCtx(ctx context.Context, err error, msg string, opts ...LogType) {
	if err == nil {
		return
	}
	args := append([]SlogAttr{}, newSlogArgs(opts...)...)
	args = append(args,
		slog.String("err_msg", fmt.Sprintf("%v", err)),
		slog.String("err_type", fmt.Sprintf("%T", err)),
	)
	opts = append(opts, OptError(err))
	l.Error(msg, opts...)
	panic_logger.Error(msg, args...)
	panic(panic_str.String())
}

func (l *Logger) CheckPanic(err error, msg string, opts ...LogType) {
	l.CheckPanicCtx(WithCtxCallerSkip(context.Background(), 5), err, msg, opts...)
}
func (l *Logger) CheckPanicln(err error, a ...any) {
	l.CheckPanicCtx(WithCtxCallerSkip(context.Background(), 5), err, fmt.Sprintln(a...))
}
