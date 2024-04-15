package log

import (
	"sync"

	"go.uber.org/zap"
)

var defaultFieldPool = newFieldPool()

type fieldContainer struct {
	Fields []Field
}

type fieldPool struct {
	pool sync.Pool
}

func newFieldPool() *fieldPool {
	return &fieldPool{
		pool: sync.Pool{
			New: func() any {
				return &fieldContainer{make([]zap.Field, 0, 32)}
			},
		},
	}
}

func (p *fieldPool) Get() *fieldContainer {
	return p.pool.Get().(*fieldContainer)
}

func (p *fieldPool) Put(c *fieldContainer) {
	c.Fields = c.Fields[:0]
	p.pool.Put(c)
}
