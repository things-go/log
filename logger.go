package log

import (
	"context"
	"fmt"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log wrap zap logger
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
	if !l.level.Enabled(level) {
		return
	}
	if len(l.fn) == 0 {
		l.Sugar().Log(level, args...)
	} else {
		l.Logx(ctx, level, getMessage("", args))
	}
}

func (l *Log) Logf(ctx context.Context, level Level, template string, args ...any) {
	if !l.level.Enabled(level) {
		return
	}
	if len(l.fn) == 0 {
		l.Sugar().Logf(level, template, args...)
	} else {
		l.Logx(ctx, level, getMessage(template, args))
	}
}

func (l *Log) Logw(ctx context.Context, level Level, msg string, keysAndValues ...any) {
	if !l.level.Enabled(level) {
		return
	}
	if len(l.fn) == 0 {
		l.Sugar().Logw(level, msg, keysAndValues...)
	} else {
		l.Logx(ctx, level, msg, l.sweetenFields(keysAndValues)...)
	}
}

func (l *Log) Logx(ctx context.Context, level Level, msg string, fields ...Field) {
	if !l.level.Enabled(level) {
		return
	}

	if len(l.fn) == 0 {
		l.log.Log(level, msg, fields...)
	} else {
		tmpFields := fieldPool.Get()
		defer fieldPool.Put(tmpFields)
		for _, f := range l.fn {
			tmpFields = append(tmpFields, f(ctx))
		}
		tmpFields = append(tmpFields, fields...)
		l.log.Log(level, msg, tmpFields...)
	}
}

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
	_multipleErrMsg     = "Multiple errors without a key."
)

// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs []interface{}) string {
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

func (l *Log) sweetenFields(args []interface{}) []Field {
	if len(args) == 0 {
		return nil
	}

	var (
		// Allocate enough space for the worst case; if users pass only structured
		// fields, we shouldn't penalize them with extra allocations.
		fields    = make([]Field, 0, len(args))
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
	key, value interface{}
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
