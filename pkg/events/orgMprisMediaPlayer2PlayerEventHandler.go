package events

import (
	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/internal"
	"github.com/quarckster/go-mpris-server/pkg/server"
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
	mpris *server.Server,
) *orgMprisMediaPlayer2PlayerEventHandler {
	eventHandler := orgMprisMediaPlayer2PlayerEventHandler{
		mpris:         mpris,
		iface:         "org.mpris.MediaPlayer2.Player",
		adapter:       mpris.PlayerAdapter,
		allProps:      allPlayerProps(mpris.PlayerAdapter),
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
		onOptionsProps:   onOptionsProps(mpris.PlayerAdapter),
	}
	return &eventHandler
}

type orgMprisMediaPlayer2PlayerEventHandler struct {
	mpris            *server.Server
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
	return internal.EmitPropertiesChanged(o.mpris.Conn, o.iface, changes)
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
	o.mpris.Conn.Emit("/org/mpris/MediaPlayer2", o.iface+".Seeked", int64(position))
	return o.EmitChanges(o.onSeekProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnOptions() error {
	return o.EmitChanges(o.onOptionsProps)
}

func (o *orgMprisMediaPlayer2PlayerEventHandler) OnAll() error {
	return o.EmitChanges(o.allProps)
}
