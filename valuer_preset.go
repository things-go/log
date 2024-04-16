package log

import (
	"context"
	"runtime"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func caller(depth int) (file string, line int) {
	ok := false
	for i := depth; i < depth+10; i++ {
		_, file, line, ok = runtime.Caller(i)
		if ok && !skipPackage(file, "github.com/things-go/log") {
			return file, line
		}
	}
	return file, line
}

func skipPackage(file string, skipPackages ...string) bool {
	if strings.HasSuffix(file, "_test.go") {
		return false
	}
	for _, p := range skipPackages {
		if strings.Contains(file, p) {
			return true
		}
	}
	return false
}

// Caller returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) Valuer {
	return func(context.Context) Field {
		file, line := caller(depth)
		idx := strings.LastIndexByte(file, '/')
		return zap.String("caller", file[idx+1:]+":"+strconv.Itoa(line))
	}
}

// File returns a Valuer that returns a pkg/file:line description of the caller.
func File(depth int) Valuer {
	return func(context.Context) Field {
		file, line := caller(depth)
		return zap.String("file", file+":"+strconv.Itoa(line))
	}
}

// Package returns a Valuer that returns an immutable Valuer which key is pkg
func Package(v string) Valuer {
	return ImmutString("pkg", v)
}

func App(v string) Valuer {
	return ImmutString("app", v)
}
func Component(v string) Valuer {
	return ImmutString("component", v)
}
func Module(v string) Valuer {
	return ImmutString("module", v)
}
func Unit(v string) Valuer {
	return ImmutString("unit", v)
}
func Kind(v string) Valuer {
	return ImmutString("kind", v)
}
func Type(v string) Valuer {
	return ImmutString("type", v)
}
func TraceId(f func(c context.Context) string) Valuer {
	return FromString("traceId", f)
}
func RequestId(f func(c context.Context) string) Valuer {
	return FromString("requestId", f)
}
