package log

import (
	"sync"

	"go.uber.org/zap"
)

var defaultFieldPool = newFieldPool()

type fieldContainer struct {
	Fields []Field
}

func (c *fieldContainer) reset() *fieldContainer {
	c.Fields = c.Fields[:0]
	return c
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
	c := p.pool.Get().(*fieldContainer)
	return c.reset()
}

func (p *fieldPool) Put(c *fieldContainer) {
	p.pool.Put(c.reset())
}
