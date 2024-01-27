package logus_core

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger              *slog.Logger
	enable_file_showing bool
	enable_json_format  bool
	log_level           *slog.LevelVar
	default_params      []SlogParam
}

type LoggerParam func(r *Logger)

func WithJsonFormat(state bool) LoggerParam {
	return func(logger *Logger) {
		logger.enable_json_format = state
	}
}

func WithFileShowing(state bool) LoggerParam {
	return func(logger *Logger) {
		logger.enable_file_showing = state
	}
}

func WithLogLevelStr(log_level_str string) LoggerParam {
	return WithLogLevel(LogLevel(log_level_str))
}

func WithLogLevel(log_level_str LogLevel) LoggerParam {
	return func(logger *Logger) {
		var log_level LogLevel = LogLevel(log_level_str)

		var programLevel = new(slog.LevelVar) // Info by default
		switch log_level {
		case LEVEL_DEBUG:
			programLevel.Set(slog.LevelDebug)
		case LEVEL_INFO:
			programLevel.Set(slog.LevelInfo)
		case LEVEL_WARN:
			programLevel.Set(slog.LevelWarn)
		case LEVEL_ERROR:
			programLevel.Set(slog.LevelError)
		default:
			programLevel.Set(slog.LevelWarn)
		}

		logger.log_level = programLevel
	}
}

func NewLogger(
	options ...LoggerParam,
) *Logger {

	logger := &Logger{}

	WithJsonFormat(bool(EnvTurnJSON))(logger)
	WithFileShowing(EnvTurnFileShowing)(logger)
	WithLogLevel(LEVEL_WARN)(logger)

	for _, opt := range options {
		opt(logger)
	}

	if logger.enable_json_format {
		logger.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logger.log_level}))
	} else {
		logger.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logger.log_level}))
	}
	return logger
}

func (l *Logger) WithFields(opts ...SlogParam) *Logger {
	var new_logger Logger = *l
	new_logger.default_params = append(l.default_params, opts...)
	return &new_logger
}
