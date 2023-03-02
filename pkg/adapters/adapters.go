package adapters

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
