package internal

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/pkg/adapters"
)

func NewOrgMprisMediaPlayer2Player(adapter adapters.OrgMprisMediaPlayer2PlayerAdapter) *OrgMprisMediaPlayer2Player {
	return &OrgMprisMediaPlayer2Player{
		Adapter: adapter,
	}
}

type OrgMprisMediaPlayer2Player struct {
	Adapter adapters.OrgMprisMediaPlayer2PlayerAdapter
	mut     sync.RWMutex
}

func (p *OrgMprisMediaPlayer2Player) Next() *dbus.Error {
	return makeError(p.Adapter.Next())
}

func (p *OrgMprisMediaPlayer2Player) Previous() *dbus.Error {
	return makeError(p.Adapter.Previous())
}

func (p *OrgMprisMediaPlayer2Player) Pause() *dbus.Error {
	return makeError(p.Adapter.Pause())
}

func (p *OrgMprisMediaPlayer2Player) PlayPause() *dbus.Error {
	return makeError(p.Adapter.PlayPause())
}

func (p *OrgMprisMediaPlayer2Player) Stop() *dbus.Error {
	return makeError(p.Adapter.Stop())
}

func (p *OrgMprisMediaPlayer2Player) Play() *dbus.Error {
	return makeError(p.Adapter.Play())
}

func (p *OrgMprisMediaPlayer2Player) Seek(offset int64) *dbus.Error {
	return makeError(p.Adapter.Seek(offset))
}

func (p *OrgMprisMediaPlayer2Player) SetPosition(trackId string, position int64) *dbus.Error {
	return makeError(p.Adapter.SetPosition(trackId, position))
}

func (p *OrgMprisMediaPlayer2Player) OpenUri(uri string) *dbus.Error {
	return makeError(p.Adapter.OpenUri(uri))
}

func (p *OrgMprisMediaPlayer2Player) Metadata() (map[string]dbus.Variant, error) {
	metadata, err := p.Adapter.Metadata()
	if err != nil {
		return map[string]dbus.Variant{}, err
	}
	return metadata.MakeMap(), nil
}

func (p *OrgMprisMediaPlayer2Player) GetMethods() map[string]interface{} {
	return map[string]interface{}{
		"PlaybackStatus": p.Adapter.PlaybackStatus,
		"LoopStatus":     p.Adapter.LoopStatus,
		"Rate":           p.Adapter.Rate,
		"Shuffle":        p.Adapter.Shuffle,
		"Metadata":       p.Metadata,
		"Volume":         p.Adapter.Volume,
		"Position":       p.Adapter.Position,
		"MinimumRate":    p.Adapter.MinimumRate,
		"MaximumRate":    p.Adapter.MaximumRate,
		"CanGoNext":      p.Adapter.CanGoNext,
		"CanGoPrevious":  p.Adapter.CanGoPrevious,
		"CanPlay":        p.Adapter.CanPlay,
		"CanPause":       p.Adapter.CanPause,
		"CanSeek":        p.Adapter.CanSeek,
		"CanControl":     p.Adapter.CanControl,
	}
}

func (p *OrgMprisMediaPlayer2Player) SetMethods() map[string]interface{} {
	return map[string]interface{}{
		"LoopStatus": p.Adapter.SetLoopStatus,
		"Rate":       p.Adapter.SetRate,
		"Shuffle":    p.Adapter.SetShuffle,
		"Volume":     p.Adapter.SetVolume,
	}
}