package events

import "github.com/quarckster/go-mpris-server/pkg/types"

func NewEventHandler(
	rootAdapter types.OrgMprisMediaPlayer2Adapter,
	playerAdapter types.OrgMprisMediaPlayer2PlayerAdapter,
) (*EventHandler, error) {
	rootEventHandler, err := newOrgMprisMediaPlayer2EventHandler(rootAdapter)
	if err != nil {
		return nil, err
	}
	playerEventHandler, err := newOrgMprisMediaPlayer2PlayerEventHandler(playerAdapter)
	if err != nil {
		return nil, err
	}
	return &EventHandler{Root: rootEventHandler, Player: playerEventHandler}, nil
}

type EventHandler struct {
	Root   types.OrgMprisMediaPlayer2EventAdapter
	Player types.OrgMprisMediaPlayer2PlayerEventAdapter
}
