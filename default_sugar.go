package log

import "context"

func Debug(args ...any) {
	defaultLogger.Debug(args...)
}
func DebugContext(ctx context.Context, args ...any) {
	defaultLogger.DebugContext(ctx, args...)
}
func Info(args ...any) {
	defaultLogger.Info(args...)
}
func InfoContext(ctx context.Context, args ...any) {
	defaultLogger.InfoContext(ctx, args...)
}
func Warn(args ...any) {
	defaultLogger.Warn(args...)
}
func WarnContext(ctx context.Context, args ...any) {
	defaultLogger.WarnContext(ctx, args...)
}
func Error(args ...any) {
	defaultLogger.Error(args...)
}
func ErrorContext(ctx context.Context, args ...any) {
	defaultLogger.ErrorContext(ctx, args...)
}
func DPanic(args ...any) {
	defaultLogger.DPanic(args...)
}
func DPanicContext(ctx context.Context, args ...any) {
	defaultLogger.DPanicContext(ctx, args...)
}
func Panic(args ...any) {
	defaultLogger.Panic(args...)
}
func PanicContext(ctx context.Context, args ...any) {
	defaultLogger.PanicContext(ctx, args...)
}
func Fatal(args ...any) {
	defaultLogger.Fatal(args...)
}
func FatalContext(ctx context.Context, args ...any) {
	defaultLogger.FatalContext(ctx, args...)
}
func Debugf(template string, args ...any) {
	defaultLogger.Debugf(template, args...)
}
func DebugfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.DebugfContext(ctx, template, args...)
}
func Infof(template string, args ...any) {
	defaultLogger.Infof(template, args...)
}
func InfofContext(ctx context.Context, template string, args ...any) {
	defaultLogger.InfofContext(ctx, template, args...)
}
func Warnf(template string, args ...any) {
	defaultLogger.Warnf(template, args...)
}
func WarnfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.WarnfContext(ctx, template, args...)
}
func Errorf(template string, args ...any) {
	defaultLogger.Errorf(template, args...)
}
func ErrorfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.ErrorfContext(ctx, template, args...)
}
func DPanicf(template string, args ...any) {
	defaultLogger.DPanicf(template, args...)
}
func DPanicfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.DPanicfContext(ctx, template, args...)
}
func Panicf(template string, args ...any) {
	defaultLogger.Panicf(template, args...)
}
func PanicfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.PanicfContext(ctx, template, args...)
}
func Fatalf(template string, args ...any) {
	defaultLogger.Fatalf(template, args...)
}
func FatalfContext(ctx context.Context, template string, args ...any) {
	defaultLogger.FatalfContext(ctx, template, args...)
}
func Debugw(msg string, keysAndValues ...any) {
	defaultLogger.Debugw(msg, keysAndValues...)
}
func DebugwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.DebugwContext(ctx, msg, keysAndValues...)
}
func Infow(msg string, keysAndValues ...any) {
	defaultLogger.Infow(msg, keysAndValues...)
}
func InfowContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.InfowContext(ctx, msg, keysAndValues...)
}
func Warnw(msg string, keysAndValues ...any) {
	defaultLogger.Warnw(msg, keysAndValues...)
}
func WarnwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.WarnwContext(ctx, msg, keysAndValues...)
}
func Errorw(msg string, keysAndValues ...any) {
	defaultLogger.Errorw(msg, keysAndValues...)
}
func ErrorwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.ErrorwContext(ctx, msg, keysAndValues...)
}
func DPanicw(msg string, keysAndValues ...any) {
	defaultLogger.DPanicw(msg, keysAndValues...)
}
func DPanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.DPanicwContext(ctx, msg, keysAndValues...)
}
func Panicw(msg string, keysAndValues ...any) {
	defaultLogger.Panicw(msg, keysAndValues...)
}
func PanicwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.PanicwContext(ctx, msg, keysAndValues...)
}
func Fatalw(msg string, keysAndValues ...any) {
	defaultLogger.Fatalw(msg, keysAndValues...)
}
func FatalwContext(ctx context.Context, msg string, keysAndValues ...any) {
	defaultLogger.FatalwContext(ctx, msg, keysAndValues...)
}
