package events

import (
	"github.com/YutongGu/go-librespot/events/impl"
	"github.com/YutongGu/go-librespot/events/plugin"
)

var Plugin plugin.Interface = impl.Impl{}
