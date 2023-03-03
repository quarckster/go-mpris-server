package internal

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/pkg/adapters"
)

func NewOrgMprisMediaPlayer2(adapter adapters.OrgMprisMediaPlayer2Adapter) *OrgMprisMediaPlayer2 {
	return &OrgMprisMediaPlayer2{
		Adapter: adapter,
	}
}

type OrgMprisMediaPlayer2 struct {
	Adapter adapters.OrgMprisMediaPlayer2Adapter
	mut     sync.RWMutex
}

func (r *OrgMprisMediaPlayer2) Raise() *dbus.Error {
	return makeError(r.Adapter.Raise())
}

func (r *OrgMprisMediaPlayer2) Quit() *dbus.Error {
	return makeError(r.Adapter.Quit())
}

func (r *OrgMprisMediaPlayer2) GetMethods() map[string]interface{} {
	return map[string]interface{}{
		"CanQuit":             r.Adapter.CanQuit,
		"Fullscreen":          r.Adapter.Fullscreen,
		"CanSetFullscreen":    r.Adapter.CanSetFullscreen,
		"CanRaise":            r.Adapter.CanRaise,
		"HasTrackList":        r.Adapter.HasTrackList,
		"Identity":            r.Adapter.Identity,
		"DesktopEntry":        r.Adapter.DesktopEntry,
		"SupportedUriSchemes": r.Adapter.SupportedUriSchemes,
		"SupportedMimeTypes":  r.Adapter.SupportedMimeTypes,
	}
}

func (r *OrgMprisMediaPlayer2) SetMethods() map[string]interface{} {
	return map[string]interface{}{
		"Fullscreen": r.Adapter.SetFullscreen,
	}
}
