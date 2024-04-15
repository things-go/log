package log

import "context"

func Debugx(msg string, fields ...Field) {
	defaultLogger.Debugx(msg, fields...)
}
func DebugxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.DebugxContext(ctx, msg, fields...)
}
func Infox(msg string, fields ...Field) {
	defaultLogger.Infox(msg, fields...)
}
func InfoxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.InfoxContext(ctx, msg, fields...)
}
func Warnx(msg string, fields ...Field) {
	defaultLogger.Warnx(msg, fields...)
}
func WarnxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.WarnxContext(ctx, msg, fields...)
}
func Errorx(msg string, fields ...Field) {
	defaultLogger.Errorx(msg, fields...)
}
func ErrorxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.ErrorxContext(ctx, msg, fields...)
}
func DPanicx(msg string, fields ...Field) {
	defaultLogger.DPanicx(msg, fields...)
}
func DPanicxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.DPanicxContext(ctx, msg, fields...)
}
func Panicx(msg string, fields ...Field) {
	defaultLogger.Panicx(msg, fields...)
}
func PanicxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.PanicxContext(ctx, msg, fields...)
}
func Fatalx(msg string, fields ...Field) {
	defaultLogger.Fatalx(msg, fields...)
}
func FatalxContext(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.FatalxContext(ctx, msg, fields...)
}
