package log

import "context"

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
