package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger = NewLoggerWith(zap.NewNop(), zap.NewAtomicLevel())

// ReplaceGlobals replaces the global Log,
func ReplaceGlobals(logger *Log) { defaultLogger = logger }

// SetLevelWithText alters the logging level.
// ParseAtomicLevel set the logging level based on a lowercase or all-caps ASCII
// representation of the log level.
// If the provided ASCII representation is
// invalid an error is returned.
func SetLevelWithText(text string) error { return defaultLogger.SetLevelWithText(text) }

// SetLevel alters the logging level.
func SetLevel(lv zapcore.Level) *Log { return defaultLogger.SetLevel(lv) }

// GetLevel returns the minimum enabled log level.
func GetLevel() zapcore.Level { return defaultLogger.GetLevel() }

// Enabled returns true if the given level is at or above this level.
func Enabled(lvl zapcore.Level) bool { return defaultLogger.Enabled(lvl) }

// V returns true if the given level is at or above this level.
// same as Enabled
func V(lvl int) bool { return defaultLogger.V(lvl) }

// SetDefaultValuer set default field function, which hold always until you call WithContext.
func SetDefaultValuer(vs ...Valuer) *Log {
	return defaultLogger.SetDefaultValuer(vs...)
}

// WithValuer with field function.
func WithValuer(vs ...Valuer) *Log {
	return defaultLogger.WithValuer(vs...)
}

// WithNewValuer return log with new Valuer function without default Valuer.
func WithNewValuer(fs ...Valuer) *Log {
	return defaultLogger.WithNewValuer(fs...)
}

// WithContext return log with inject context.
func WithContext(ctx context.Context) *Log {
	return defaultLogger.WithContext(ctx)
}

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func Sugar() *zap.SugaredLogger { return defaultLogger.Sugar() }

// Logger return internal logger
func Logger() *zap.Logger { return defaultLogger.Logger() }

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//
//	 sugaredLogger.With(
//	   "hello", "world",
//	   "failure", errors.New("oh no"),
//	   "count", 42,
//	   "user", User{Name: "alice"},
//	)
//
// is the equivalent of
//
//	unsugared.With(
//	  String("hello", "world"),
//	  String("failure", "oh no"),
//	  Stack(),
//	  Int("count", 42),
//	  Object("user", User{Name: "alice"}),
//	)
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func With(fields ...Field) *Log { return defaultLogger.With(fields...) }

// Named adds a sub-scope to the logger's name. See Log.Named for details.
func Named(name string) *Log { return defaultLogger.Named(name) }

// Sync flushes any buffered log entries.
func Sync() error { return defaultLogger.Sync() }

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...any) { defaultLogger.Debug(args...) }

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...any) { defaultLogger.Info(args...) }

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...any) { defaultLogger.Warn(args...) }

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...any) { defaultLogger.Error(args...) }

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...any) { defaultLogger.DPanic(args...) }

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...any) { defaultLogger.Panic(args...) }

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...any) { defaultLogger.Fatal(args...) }

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...any) { defaultLogger.Debugf(template, args...) }

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...any) { defaultLogger.Infof(template, args...) }

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...any) { defaultLogger.Warnf(template, args...) }

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...any) { defaultLogger.Errorf(template, args...) }

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...any) { defaultLogger.DPanicf(template, args...) }

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...any) { defaultLogger.Panicf(template, args...) }

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...any) { defaultLogger.Fatalf(template, args...) }

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...any) { defaultLogger.Debugw(msg, keysAndValues...) }

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...any) { defaultLogger.Infow(msg, keysAndValues...) }

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...any) { defaultLogger.Warnw(msg, keysAndValues...) }

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...any) { defaultLogger.Errorw(msg, keysAndValues...) }

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...any) { defaultLogger.DPanicw(msg, keysAndValues...) }

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...any) { defaultLogger.Panicw(msg, keysAndValues...) }

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...any) { defaultLogger.Fatalw(msg, keysAndValues...) }
