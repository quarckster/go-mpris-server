package events

import (
	"github.com/quarckster/go-mpris-server/pkg/server"
	"github.com/quarckster/go-mpris-server/pkg/types"
)

func NewEventHandler(mpris *server.Server) *EventHandler {
	rootEventHandler := newOrgMprisMediaPlayer2EventHandler(mpris)
	playerEventHandler := newOrgMprisMediaPlayer2PlayerEventHandler(mpris)
	return &EventHandler{Root: rootEventHandler, Player: playerEventHandler}
}

type EventHandler struct {
	Root   types.OrgMprisMediaPlayer2EventHandler
	Player types.OrgMprisMediaPlayer2PlayerEventHandler
}
