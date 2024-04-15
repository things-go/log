package log

import (
	"sync"

	"go.uber.org/zap"
)

var fieldPool = newFieldPool()

type innerFieldPool struct {
	pool sync.Pool
}

func newFieldPool() *innerFieldPool {
	return &innerFieldPool{
		pool: sync.Pool{
			New: func() any {
				return make([]zap.Field, 0, 16)
			},
		},
	}
}

func (p *innerFieldPool) Get() []zap.Field {
	return p.pool.Get().([]zap.Field)
}

func (p *innerFieldPool) Put(fields []zap.Field) {
	p.pool.Put(fields)
}
