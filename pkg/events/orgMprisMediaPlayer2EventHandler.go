package events

import (
	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/internal"
	"github.com/quarckster/go-mpris-server/pkg/types"
)

func allRootProps(adapter types.OrgMprisMediaPlayer2Adapter) []string {
	props := []string{
		"CanQuit",
		"CanRaise",
		"HasTrackList",
		"Identity",
		"SupportedUriSchemes",
		"SupportedMimeTypes",
	}
	var ok bool
	_, ok = adapter.(types.OrgMprisMediaPlayer2AdapterFullscreen)
	if ok {
		props = append(props, "Fullscreen")
	}
	_, ok = adapter.(types.OrgMprisMediaPlayer2AdapterCanSetFullscreen)
	if ok {
		props = append(props, "CanSetFullscreen")
	}
	_, ok = adapter.(types.OrgMprisMediaPlayer2AdapterDesktopEntry)
	if ok {
		props = append(props, "DesktopEntry")
	}
	return props
}

func newOrgMprisMediaPlayer2EventHandler(adapter types.OrgMprisMediaPlayer2Adapter) (*orgMprisMediaPlayer2EventHandler, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return &orgMprisMediaPlayer2EventHandler{}, err
	}
	eventHandler := orgMprisMediaPlayer2EventHandler{
		conn:     conn,
		iface:    "org.mpris.MediaPlayer2",
		adapter:  adapter,
		allProps: allRootProps(adapter),
	}
	return &eventHandler, nil
}

type orgMprisMediaPlayer2EventHandler struct {
	conn     *dbus.Conn
	iface    string
	adapter  types.OrgMprisMediaPlayer2Adapter
	allProps []string
}

func (o *orgMprisMediaPlayer2EventHandler) EmitChanges(props []string) error {
	changes, err := internal.Changes(o.adapter, props)
	if err != nil {
		return err
	}
	return internal.EmitPropertiesChanged(o.conn, o.iface, changes)
}

func (o *orgMprisMediaPlayer2EventHandler) OnAll() error {
	return o.EmitChanges(o.allProps)
}
