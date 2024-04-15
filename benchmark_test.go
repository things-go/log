package log_test

import (
	"context"
	"io"
	"testing"

	"github.com/things-go/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Benchmark_NativeLogger(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	logger := zap.New(core)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("success",
			zap.String("name", `jack`),
			zap.Int("age", 18),
		)
	}
}

func newDiscardLogger() *log.Log {
	return log.NewLogger(log.WithAdapter("custom", io.Discard))
}
func dfltCtx(ctx context.Context) log.Field {
	return zap.String("dflt_key", "dflt_value")
}

func Benchmark_Logger(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.
			InfoxContext(
				ctx,
				"success",
				log.String("name", `jack`),
				log.Int("age", 18),
				dfltCtx(ctx),
			)
	}
}

func Benchmark_Logger_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	logger.SetDefaultValuer(dfltCtx)
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.
			InfoxContext(
				ctx,
				"success",
				log.String("name", `jack`),
				log.Int("age", 18),
			)
	}
}

func Benchmark_NativeSugar(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	logger := zap.New(core).Sugar()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Infow("success",
			"name", `jack`,
			"age", 18,
		)
	}
}

func Benchmark_SugarKeyValuePair(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.InfowContext(ctx,
			"success",
			log.String("name", `jack`),
			log.Int("age", 18),
			dfltCtx(ctx),
		)
	}
}

func Benchmark_SugarKeyValuePair_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	logger.SetDefaultValuer(dfltCtx)
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.InfowContext(ctx,
			"success",
			log.String("name", `jack`),
			log.Int("age", 18),
		)
	}
}

func Benchmark_SugarKeyValuePair_Use_WithFields(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.With(
			log.String("name", `jack`),
			log.Int("age", 18),
			dfltCtx(ctx),
		).InfowContext(ctx, "success")
	}
}

func Benchmark_SugarKeyValuePair_Use_WithFields_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	logger.SetDefaultValuer(dfltCtx)
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.With(
			log.String("name", `jack`),
			log.Int("age", 18),
		).InfowContext(ctx, "success")
	}
}

func Benchmark_SugarKeyValuePair_Use_WithValuer(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.WithValuer(
			log.ImmutString("name", `jack`),
			log.ImmutInt("age", 18),
			dfltCtx,
		).InfowContext(ctx, "success")
	}
}

func Benchmark_SugarKeyValuePair_Use_WithValuer_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	logger.SetDefaultValuer(dfltCtx)
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.WithValuer(
			log.ImmutString("name", `jack`),
			log.ImmutInt("age", 18),
		).InfowContext(ctx, "success")
	}
}

func Benchmark_SugarFormat(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.WithValuer(
			func(ctx context.Context) log.Field {
				return log.String("name", `jack`)
			},
			func(ctx context.Context) log.Field {
				return log.Int("age", 18)
			},
			dfltCtx,
		).InfofContext(ctx,
			"success",
		)
	}
}

func Benchmark_SugarFormat_Use_Hook(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := newDiscardLogger()
	logger.SetDefaultValuer(dfltCtx)
	b.StartTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		logger.WithValuer(
			log.ImmutString("name", `jack`),
			log.ImmutInt("age", 18),
		).InfofContext(ctx,
			"success",
		)
	}
}