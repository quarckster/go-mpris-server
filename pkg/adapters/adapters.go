package adapters

import "github.com/quarckster/go-mpris-server/pkg/types"

type OrgMprisMediaPlayer2Adapter interface {
	Raise() error
	Quit() error
	CanQuit() (bool, error)
	Fullscreen() (bool, error)
	SetFullscreen(bool) error
	CanSetFullscreen() (bool, error)
	CanRaise() (bool, error)
	HasTrackList() (bool, error)
	Identity() (string, error)
	DesktopEntry() (string, error)
	SupportedUriSchemes() ([]string, error)
	SupportedMimeTypes() ([]string, error)
}

type OrgMprisMediaPlayer2PlayerAdapter interface {
	Next() error
	Previous() error
	Pause() error
	PlayPause() error
	Stop() error
	Play() error
	Seek(offset int64) error
	SetPosition(trackId string, position int64) error
	OpenUri(uri string) error
	PlaybackStatus() (types.PlaybackStatus, error)
	LoopStatus() (types.LoopStatus, error)
	SetLoopStatus(types.LoopStatus) error
	Rate() (float64, error)
	SetRate(float64) error
	Shuffle() (bool, error)
	SetShuffle(bool) error
	Metadata() (types.Metadata, error)
	Volume() (float64, error)
	SetVolume(float64) error
	Position() (int64, error)
	MinimumRate() (float64, error)
	MaximumRate() (float64, error)
	CanGoNext() (bool, error)
	CanGoPrevious() (bool, error)
	CanPlay() (bool, error)
	CanPause() (bool, error)
	CanSeek() (bool, error)
	CanControl() (bool, error)
}
