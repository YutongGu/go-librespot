package plugin

import (
	librespot "github.com/YutongGu/go-librespot"
	"github.com/YutongGu/go-librespot/mercury"
	"github.com/YutongGu/go-librespot/player"
	"github.com/YutongGu/go-librespot/spclient"
)

type Interface interface {
	NewEventManager(log librespot.Logger, state *librespot.AppState, hg *mercury.Client, sp *spclient.Spclient, username string) (player.EventManager, error)
}
