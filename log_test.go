package log_test

import (
	"context"
	"testing"

	"github.com/things-go/log"
)

func init() {
	l, lv := log.New(log.WithConfig(log.Config{Level: "debug", Format: "json"}))
	log.ReplaceGlobals(log.NewLoggerWith(l, lv))
	log.SetDefaultValuer(
		log.Caller(3),
		func(ctx context.Context) log.Field {
			return log.String("deft_key1", "deft_val1")
		},
	)
}

func Test_LoggerNormal(t *testing.T) {
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Info("info")
	log.Error("Error")
	log.DPanic("DPanic")
}

func Test_LoggerFormater(t *testing.T) {
	log.Debugf("Debugf: %s", "debug")
	log.Infof("Infof: %s", "info")
	log.Warnf("Warnf: %s", "warn")
	log.Infof("Infof: %s", "info")
	log.Errorf("Errorf: %s", "error")
	log.DPanicf("DPanicf: %s", "dPanic")
}

func Test_LoggerKeyValue(t *testing.T) {
	log.Debugw("Debugw", "Debugw", "w")
	log.Infow("Infow", "Infow", "w")
	log.Warnw("Warnw", "Warnw", "w")
	log.Infow("Infow", "Infow", "w")
	log.Errorw("Errorw", "Errorw", "w")
	log.DPanicw("DPanicw", "DPanicw", "w")
}

func TestPanic(t *testing.T) {
	shouldPanic(t, func() {
		log.Panic("Panic")
	})
	shouldPanic(t, func() {
		log.Panicf("Panicf: %s", "panic")
	})
	shouldPanic(t, func() {
		log.Panicw("Panicw: %s", "panic", "w")
	})
}

func Test_LoggerWith(t *testing.T) {
	log.With(
		log.String("string", "bb"),
		log.Int16("int16", 100),
	).
		Debug("debug with")
}

func Test_LoggerNamed(t *testing.T) {
	log.Named("another").Debug("debug named")
}
func Test_Logger_ZapLogger(t *testing.T) {
	log.Logger().Debug("desugar")
}

func Test_LoggerNamespace(t *testing.T) {
	log.Logger().With(log.Namespace("aaaa")).With(log.String("xx", "yy"), log.String("aa", "bb")).Debug("with namespace")

	_ = log.Sync()
}

type ctxKey struct{}

func Test_Logger_WithContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey{}, "ctx_value")
	ctxValuer := func(ctx context.Context) log.Field {
		s, ok := ctx.Value(ctxKey{}).(string)
		if !ok {
			return log.Skip()
		}
		return log.String("ctx_key", s)
	}
	log.WithContext(ctx).
		WithValuer(ctxValuer).
		Debug("with context")

	log.WithContext(ctx).
		WithValuer(ctxValuer).
		Debug("with field fn")
}

func shouldPanic(t *testing.T, f func()) {
	defer func() {
		e := recover()
		if e == nil {
			t.Errorf("should panic but not")
		}
	}()
	f()
}
