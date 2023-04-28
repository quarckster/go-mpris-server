package internal

import (
	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/pkg/types"
)

func NewOrgMprisMediaPlayer2(adapter types.OrgMprisMediaPlayer2Adapter) *OrgMprisMediaPlayer2 {
	return &OrgMprisMediaPlayer2{
		Adapter: adapter,
	}
}

type OrgMprisMediaPlayer2 struct {
	Adapter types.OrgMprisMediaPlayer2Adapter
}

func (r *OrgMprisMediaPlayer2) Raise() *dbus.Error {
	return makeError(r.Adapter.Raise())
}

func (r *OrgMprisMediaPlayer2) Quit() *dbus.Error {
	return makeError(r.Adapter.Quit())
}

func (r *OrgMprisMediaPlayer2) GetMethods() map[string]interface{} {
	methods := map[string]interface{}{
		"CanQuit":             r.Adapter.CanQuit,
		"CanRaise":            r.Adapter.CanRaise,
		"HasTrackList":        r.Adapter.HasTrackList,
		"Identity":            r.Adapter.Identity,
		"SupportedUriSchemes": r.Adapter.SupportedUriSchemes,
		"SupportedMimeTypes":  r.Adapter.SupportedMimeTypes,
	}
	fullscreen, ok := r.Adapter.(types.OrgMprisMediaPlayer2AdapterFullscreen)
	if ok {
		methods["Fullscreen"] = fullscreen.Fullscreen
	}
	canSetFullscreen, ok := r.Adapter.(types.OrgMprisMediaPlayer2AdapterCanSetFullscreen)
	if ok {
		methods["CanSetFullscreen"] = canSetFullscreen.CanSetFullscreen
	}
	desktopEntry, ok := r.Adapter.(types.OrgMprisMediaPlayer2AdapterDesktopEntry)
	if ok {
		methods["DesktopEntry"] = desktopEntry.DesktopEntry
	}
	return methods
}

func (r *OrgMprisMediaPlayer2) SetMethods() map[string]interface{} {
	methods := make(map[string]interface{})
	setFullscreen, ok := r.Adapter.(types.OrgMprisMediaPlayer2AdapterFullscreen)
	if ok {
		methods["SetFullscreen"] = setFullscreen.SetFullscreen
	}
	return methods
}
