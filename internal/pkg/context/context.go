package context

import (
	"github.com/opencap/cappy/internal/pkg/key"
)

type Context interface {
	KeyManager() key.Manager
	SetKeyManager(key.Manager)
}

type context struct {
	keyManager key.Manager
}

var instance *context

func Instance() Context {
	if instance == nil {
		instance = &context{}
	}
	return instance
}

func (ctx *context) KeyManager() key.Manager {
	return ctx.keyManager
}

func (ctx *context) SetKeyManager(m key.Manager) {
	ctx.keyManager = m
}
