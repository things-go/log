package log

import (
	"context"
)

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
