package playplay

import (
	"github.com/YutongGu/go-librespot/playplay/impl"
	"github.com/YutongGu/go-librespot/playplay/plugin"
)

var Plugin plugin.Interface = impl.Impl{}
