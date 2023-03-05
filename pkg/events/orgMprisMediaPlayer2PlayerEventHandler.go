package events

import (
	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/internal"
	"github.com/quarckster/go-mpris-server/pkg/types"
)

func allPlayerProps(adapter types.OrgMprisMediaPlayer2PlayerAdapter) []string {
	props := []string{
		"PlaybackStatus",
		"Rate",
		"Metadata",
		"Volume",
		"Position",
		"MinimumRate",
		"MaximumRate",
		"CanGoNext",
		"CanGoPrevious",
		"CanPlay",
		"CanPause",
		"CanSeek",
		"CanControl",
	}
	var ok bool
	_, ok = adapter.(types.OrgMprisMediaPlayer2PlayerAdapterLoopStatus)
	if ok {
		props = append(props, "LoopStatus")
	}
	_, ok = adapter.(types.OrgMprisMediaPlayer2PlayerAdapterShuffle)
	if ok {
		props = append(props, "Shuffle")
	}
	return props
}

func onOptionsProps(adapter types.OrgMprisMediaPlayer2PlayerAdapter) []string {
	props := []string{
		"CanGoNext",
		"CanGoPrevious",
		"CanPause",
		"CanPlay",
	}
	var ok bool
	_, ok = adapter.(types.OrgMprisMediaPlayer2PlayerAdapterLoopStatus)
	if ok {
		props = append(props, "LoopStatus")
	}
	_, ok = adapter.(types.OrgMprisMediaPlayer2PlayerAdapterShuffle)
	if ok {
		props = append(props, "Shuffle")
	}
	return props
}

func newOrgMprisMediaPlayer2PlayerEventHandler(
	adapter types.OrgMprisMediaPlayer2PlayerAdapter,
) (*orgMprisMediaPlayer2PlayerEventHandler, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return &orgMprisMediaPlayer2PlayerEventHandler{}, err
	}
	eventHandler := orgMprisMediaPlayer2PlayerEventHandler{
		conn:          conn,
		iface:         "org.mpris.MediaPlayer2.Player",
		adapter:       adapter,
		allProps:      allPlayerProps(adapter),
		onEndedProps:  []string{"PlaybackStatus"},
		onVolumeProps: []string{"Volume"},
		onPlaybackProps: []string{
			"CanControl",
			"MaximumRate",
			"Metadata",
			"MinimumRate",
			"PlaybackStatus",
			"Rate",
		},
		onPlayPauseProps: []string{"PlaybackStatus"},
		onTitleProps:     []string{"Metadata"},
		onSeekProps:      []string{"Position"},
		onOptionsProps:   onOptionsProps(adapter),
	}
	return &eventHandler, nil
}

type orgMprisMediaPlayer2PlayerEventHandler struct {
	conn             *dbus.Conn
	iface            string
	adapter          types.OrgMprisMediaPlayer2PlayerAdapter
	allProps         []string
	onEndedProps     []string
	onVolumeProps    []string
	onPlaybackProps  []string
	onPlayPauseProps []string
	onTitleProps     []string
	onSeekProps      []string
	onOptionsProps   []string
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) EmitChanges(props []string) error {
	changes, err := internal.Changes(o.adapter, props)
	if err != nil {
		return err
	}
	return internal.EmitPropertiesChanged(o.conn, o.iface, changes)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnEnded() error {
	return o.EmitChanges(o.onEndedProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnVolume() error {
	return o.EmitChanges(o.onVolumeProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnPlayback() error {
	return o.EmitChanges(o.onPlaybackProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnPlayPause() error {
	return o.EmitChanges(o.onPlayPauseProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnTitle() error {
	return o.EmitChanges(o.onTitleProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnSeek(position types.Microseconds) error {
	o.conn.Emit("/org/mpris/MediaPlayer2", o.iface+".Seeked", int64(position))
	return o.EmitChanges(o.onSeekProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnOptions() error {
	return o.EmitChanges(o.onOptionsProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnAll() error {
	return o.EmitChanges(o.allProps)
}
