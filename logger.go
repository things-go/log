package log

import (
	"context"
	"fmt"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log wrap zap logger
// - methods named after the log level or ending in "Context" for log.Print-style logging
// - methods ending in "w" or "wContext" for loosely-typed structured logging
// - methods ending in "f" or "fContext" for log.Printf-style logging
// - methods ending in "x" or "xContext" for structured logging
type Log struct {
	log   *zap.Logger
	level zap.AtomicLevel
	fn    []Valuer
	ctx   context.Context
}

// NewLoggerWith new logger with zap logger and atomic level
func NewLoggerWith(logger *zap.Logger, lv zap.AtomicLevel) *Log {
	return &Log{
		log:   logger,
		level: lv,
		fn:    nil,
		ctx:   context.Background(),
	}
}

// NewLogger new logger
func NewLogger(opts ...Option) *Log { return NewLoggerWith(New(opts...)) }

// SetLevelWithText alters the logging level.
// ParseAtomicLevel set the logging level based on a lowercase or all-caps ASCII
// representation of the log level.
// If the provided ASCII representation is
// invalid an error is returned.
// see zapcore.Level
func (l *Log) SetLevelWithText(text string) error {
	lv, err := zapcore.ParseLevel(text)
	if err != nil {
		return err
	}
	l.level.SetLevel(lv)
	return nil
}

// SetLevel alters the logging level.
func (l *Log) SetLevel(lv zapcore.Level) *Log {
	l.level.SetLevel(lv)
	return l
}

// SetDefaultValuer set default Valuer function, which hold always until you call WithContext.
func (l *Log) SetDefaultValuer(fs ...Valuer) *Log {
	fn := make([]Valuer, 0, len(fs)+len(l.fn))
	fn = append(fn, l.fn...)
	fn = append(fn, fs...)
	l.fn = fn
	return l
}

// GetLevel returns the minimum enabled log level.
func (l *Log) GetLevel() Level { return l.level.Level() }

// Enabled returns true if the given level is at or above this level.
func (l *Log) Enabled(lvl Level) bool { return l.level.Enabled(lvl) }

// V returns true if the given level is at or above this level.
// same as Enabled
func (l *Log) V(lvl int) bool { return l.level.Enabled(zapcore.Level(lvl)) }

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func (l *Log) Sugar() *zap.SugaredLogger { return l.log.Sugar() }

// Logger return internal logger
func (l *Log) Logger() *zap.Logger { return l.log }

// WithValuer with Valuer function.
func (l *Log) WithValuer(fs ...Valuer) *Log {
	fn := make([]Valuer, 0, len(fs)+len(l.fn))
	fn = append(fn, l.fn...)
	fn = append(fn, fs...)
	return &Log{
		log:   l.log,
		level: l.level,
		fn:    fn,
		ctx:   l.ctx,
	}
}

// WithNewValuer return log with new Valuer function without default Valuer.
func (l *Log) WithNewValuer(fs ...Valuer) *Log {
	return &Log{
		log:   l.log,
		level: l.level,
		fn:    fs,
		ctx:   l.ctx,
	}
}

// WithContext return log with inject context.
//
// Deprecated: you should use XXXContext to inject context.
func (l *Log) WithContext(ctx context.Context) *Log {
	return &Log{
		log:   l.log,
		level: l.level,
		fn:    l.fn,
		ctx:   ctx,
	}
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (l *Log) With(fields ...Field) *Log {
	return &Log{
		log:   l.log.With(fields...),
		level: l.level,
		fn:    l.fn,
		ctx:   l.ctx,
	}
}

// Named adds a sub-scope to the logger's name. See Log.Named for details.
func (l *Log) Named(name string) *Log {
	return &Log{
		log:   l.log.Named(name),
		level: l.level,
		fn:    l.fn,
		ctx:   l.ctx,
	}
}

// Sync flushes any buffered log entries.
func (l *Log) Sync() error {
	return l.log.Sync()
}

func (l *Log) Log(ctx context.Context, level Level, args ...any) {
	l.Logx(ctx, level, formatMessage("", args))
}

func (l *Log) Logf(ctx context.Context, level Level, template string, args ...any) {
	l.Logx(ctx, level, formatMessage(template, args))
}

func (l *Log) Logw(ctx context.Context, level Level, msg string, keysAndValues ...any) {
	if !l.level.Enabled(level) {
		return
	}
	fc := defaultFieldPool.Get()
	defer defaultFieldPool.Put(fc)
	for _, f := range l.fn {
		fc.Fields = append(fc.Fields, f(ctx))
	}
	fc.Fields = l.appendSweetenFields(fc.Fields, keysAndValues)
	l.log.Log(level, msg, fc.Fields...)
}

func (l *Log) Logx(ctx context.Context, level Level, msg string, fields ...Field) {
	if !l.level.Enabled(level) {
		return
	}
	if len(l.fn) == 0 {
		l.log.Log(level, msg, fields...)
	} else {
		fc := defaultFieldPool.Get()
		defer defaultFieldPool.Put(fc)
		for _, f := range l.fn {
			fc.Fields = append(fc.Fields, f(ctx))
		}
		fc.Fields = append(fc.Fields, fields...)
		l.log.Log(level, msg, fc.Fields...)
	}
}

//****** named after the log level or ending in "Context" for log.Print-style logging

// Debug (see DebugContext)
func (l *Log) Debug(args ...any) {
	l.DebugContext(l.ctx, args...)
}

// DebugContext uses fmt.Sprint to construct and log a message.
func (l *Log) DebugContext(ctx context.Context, args ...any) {
	l.Log(ctx, DebugLevel, args...)
}

// Info see InfoContext
func (l *Log) Info(args ...any) {
	l.InfoContext(l.ctx, args...)
}

// InfoContext uses fmt.Sprint to construct and log a message.
func (l *Log) InfoContext(ctx context.Context, args ...any) {
	l.Log(ctx, InfoLevel, args...)
}

// Warn see WarnContext
func (l *Log) Warn(args ...any) {
	l.WarnContext(l.ctx, args...)
}

// WarnContext uses fmt.Sprint to construct and log a message.
func (l *Log) WarnContext(ctx context.Context, args ...any) {
	l.Log(ctx, WarnLevel, args...)
}

// Error see ErrorContext
func (l *Log) Error(args ...any) {
	l.ErrorContext(l.ctx, args...)
}

// ErrorContext uses fmt.Sprint to construct and log a message.
func (l *Log) ErrorContext(ctx context.Context, args ...any) {
	l.Log(ctx, ErrorLevel, args...)
}

// DPanic see DPanicContext
func (l *Log) DPanic(args ...any) {
	l.DPanicContext(l.ctx, args...)
}

// DPanicContext uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (see DPanicLevel for details.)
func (l *Log) DPanicContext(ctx context.Context, args ...any) {
	l.Log(ctx, DPanicLevel, args...)
}

// Panic see PanicContext
func (l *Log) Panic(args ...any) {
	l.PanicContext(l.ctx, args...)
}

// PanicContext uses fmt.Sprint to to construct and log a message, then panics.
func (l *Log) PanicContext(ctx context.Context, args ...any) {
	l.Log(ctx, PanicLevel, args...)
}

// Fatal see FatalContext
func (l *Log) Fatal(args ...any) {
	l.FatalContext(l.ctx, args...)
}

// FatalContext uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Log) FatalContext(ctx context.Context, args ...any) {
	l.Log(ctx, FatalLevel, args...)
}

//****** ending in "f" or "fContext" for log.Printf-style logging

// Debugf see DebugfContext
func (l *Log) Debugf(template string, args ...any) {
	l.DebugfContext(l.ctx, template, args...)
}

// DebugfContext uses fmt.Sprintf to log a templated message.
func (l *Log) DebugfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, DebugLevel, template, args...)
}

// Infof see InfofContext
func (l *Log) Infof(template string, args ...any) {
	l.InfofContext(l.ctx, template, args...)
}

// InfofContext uses fmt.Sprintf to log a templated message.
func (l *Log) InfofContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, InfoLevel, template, args...)
}

// Warnf see WarnfContext
func (l *Log) Warnf(template string, args ...any) {
	l.WarnfContext(l.ctx, template, args...)
}

// WarnfContext uses fmt.Sprintf to log a templated message.
func (l *Log) WarnfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, WarnLevel, template, args...)
}

// Errorf see ErrorfContext
func (l *Log) Errorf(template string, args ...any) {
	l.ErrorfContext(l.ctx, template, args...)
}

// ErrorfContext uses fmt.Sprintf to log a templated message.
func (l *Log) ErrorfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, ErrorLevel, template, args...)
}

// DPanicf see DPanicfContext
func (l *Log) DPanicf(template string, args ...any) {
	l.DPanicfContext(l.ctx, template, args...)
}

// DPanicfContext uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (see DPanicLevel for details.)
func (l *Log) DPanicfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, DPanicLevel, template, args...)
}

// Panicf see PanicfContext
func (l *Log) Panicf(template string, args ...any) {
	l.PanicfContext(l.ctx, template, args...)
}

// PanicfContext uses fmt.Sprintf to log a templated message, then panics.
func (l *Log) PanicfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, PanicLevel, template, args...)
}

// Fatalf see FatalfContext
func (l *Log) Fatalf(template string, args ...any) {
	l.FatalfContext(l.ctx, template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Log) FatalfContext(ctx context.Context, template string, args ...any) {
	l.Logf(ctx, FatalLevel, template, args...)
}

//****** ending in "w" or "wContext" for loosely-typed structured logging

// Debugw see DebugwContext
func (l *Log) Debugw(msg string, keysAndValues ...any) {
	l.DebugwContext(l.ctx, msg, keysAndValues...)
}

// DebugwContext logs a message with some additional context. The variadic key-value or Field
// pairs or Field are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(fields).Debug(msg)
func (l *Log) DebugwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, DebugLevel, msg, keysAndValues...)
}

// Infow see InfowContext
func (l *Log) Infow(msg string, keysAndValues ...any) {
	l.InfowContext(l.ctx, msg, keysAndValues...)
}

// InfowContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) InfowContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, InfoLevel, msg, keysAndValues...)
}

// Warnw see WarnwContext
func (l *Log) Warnw(msg string, keysAndValues ...any) {
	l.WarnwContext(l.ctx, msg, keysAndValues...)
}

// WarnwContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) WarnwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, WarnLevel, msg, keysAndValues...)
}

// Errorw see ErrorwContext
func (l *Log) Errorw(msg string, keysAndValues ...any) {
	l.ErrorwContext(l.ctx, msg, keysAndValues...)
}

// ErrorwContext logs a message with some additional context. The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) ErrorwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, ErrorLevel, msg, keysAndValues...)
}

// DPanicw see DPanicwContext
func (l *Log) DPanicw(msg string, keysAndValues ...any) {
	l.DPanicwContext(l.ctx, msg, keysAndValues...)
}

// DPanicwContext logs a message with some additional context. In development, the
// logger then panics. (see DPanicLevel for details.) The variadic key-value
// pairs or Field are treated as they are in With.
func (l *Log) DPanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, DPanicLevel, msg, keysAndValues...)
}

// Panicw see PanicwContext
func (l *Log) Panicw(msg string, keysAndValues ...any) {
	l.PanicwContext(l.ctx, msg, keysAndValues...)
}

// PanicwContext logs a message with some additional context, then panics. The
// variadic key-value pairs or Field are treated as they are in With.
func (l *Log) PanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, PanicLevel, msg, keysAndValues...)
}

func (l *Log) Fatalw(msg string, keysAndValues ...any) {
	l.FatalwContext(l.ctx, msg, keysAndValues...)
}

// FatalwContext logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs or Field are treated as they are in With.
func (l *Log) FatalwContext(ctx context.Context, msg string, keysAndValues ...any) {
	l.Logw(ctx, FatalLevel, msg, keysAndValues...)
}

//****** ending in "x" or "xContext" for structured logging

// Debug (see DebugContext)
func (l *Log) Debugx(msg string, fields ...Field) {
	l.DebugxContext(l.ctx, msg, fields...)
}

// DebugContext uses fmt.Sprint to construct and log a message.
func (l *Log) DebugxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, DebugLevel, msg, fields...)
}

// Info see InfoContext
func (l *Log) Infox(msg string, fields ...Field) {
	l.InfoxContext(l.ctx, msg, fields...)
}

// InfoContext uses fmt.Sprint to construct and log a message.
func (l *Log) InfoxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, InfoLevel, msg, fields...)
}

// Warn see WarnContext
func (l *Log) Warnx(msg string, fields ...Field) {
	l.WarnxContext(l.ctx, msg, fields...)
}

// WarnContext uses fmt.Sprint to construct and log a message.
func (l *Log) WarnxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, WarnLevel, msg, fields...)
}

// Error see ErrorContext
func (l *Log) Errorx(msg string, fields ...Field) {
	l.ErrorxContext(l.ctx, msg, fields...)
}

// ErrorContext uses fmt.Sprint to construct and log a message.
func (l *Log) ErrorxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, ErrorLevel, msg, fields...)
}

// DPanic see DPanicContext
func (l *Log) DPanicx(msg string, fields ...Field) {
	l.DPanicxContext(l.ctx, msg, fields...)
}

// DPanicContext uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (see DPanicLevel for details.)
func (l *Log) DPanicxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, DPanicLevel, msg, fields...)
}

// Panic see PanicContext
func (l *Log) Panicx(msg string, fields ...Field) {
	l.PanicxContext(l.ctx, msg, fields...)
}

// PanicContext uses fmt.Sprint to to construct and log a message, then panics.
func (l *Log) PanicxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, PanicLevel, msg, fields...)
}

// Fatal see FatalContext
func (l *Log) Fatalx(msg string, fields ...Field) {
	l.FatalxContext(l.ctx, msg, fields...)
}

// FatalContext uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Log) FatalxContext(ctx context.Context, msg string, fields ...Field) {
	l.Logx(ctx, FatalLevel, msg, fields...)
}

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
	_multipleErrMsg     = "Multiple errors without a key."
)

// formatMessage format with Sprint, Sprintf, or neither.
// copy from zap(sugar.go)
func formatMessage(template string, fmtArgs []any) string {
	if len(fmtArgs) == 0 {
		return template
	}
	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}
	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

// copy from zap(sugar.go)
func (l *Log) appendSweetenFields(fields []Field, args []any) []Field {
	if len(args) == 0 {
		return nil
	}

	var (
		invalid   invalidPairs
		seenError bool
	)

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// If it is an error, consume it and move on.
		if err, ok := args[i].(error); ok {
			if !seenError {
				seenError = true
				fields = append(fields, zap.Error(err))
			} else {
				l.Errorx(_multipleErrMsg, zap.Error(err))
			}
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			l.Errorx(_oddNumberErrMsg, Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, Any(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		l.Errorx(_nonStringKeyErrMsg, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value any
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	Any("key", p.key).AddTo(enc)
	Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}
