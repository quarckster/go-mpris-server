package types

import (
	"github.com/godbus/dbus/v5"
)

type OrgMprisMediaPlayer2Adapter interface {
	Raise() error
	Quit() error
	CanQuit() (bool, error)
	CanRaise() (bool, error)
	HasTrackList() (bool, error)
	Identity() (string, error)
	SupportedUriSchemes() ([]string, error)
	SupportedMimeTypes() ([]string, error)
}

type OrgMprisMediaPlayer2AdapterFullscreen interface {
	Fullscreen() (bool, error)
	SetFullscreen(bool) error
}

type OrgMprisMediaPlayer2AdapterCanSetFullscreen interface {
	CanSetFullscreen() (bool, error)
}

type OrgMprisMediaPlayer2AdapterDesktopEntry interface {
	DesktopEntry() (string, error)
}

type OrgMprisMediaPlayer2PlayerAdapter interface {
	Next() error
	Previous() error
	Pause() error
	PlayPause() error
	Stop() error
	Play() error
	Seek(offset Microseconds) error
	SetPosition(trackId string, position Microseconds) error
	OpenUri(uri string) error
	PlaybackStatus() (PlaybackStatus, error)
	Rate() (float64, error)
	SetRate(float64) error
	Metadata() (Metadata, error)
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

type OrgMprisMediaPlayer2PlayerAdapterLoopStatus interface {
	LoopStatus() (LoopStatus, error)
	SetLoopStatus(LoopStatus) error
}

type OrgMprisMediaPlayer2PlayerAdapterShuffle interface {
	Shuffle() (bool, error)
	SetShuffle(bool) error
}

type Microseconds int64

type Metadata struct {
	TrackId        dbus.ObjectPath
	Length         Microseconds
	ArtUrl         string
	Album          string
	AlbumArtist    []string
	Artist         []string
	AsText         string
	AudioBPM       int
	AutoRating     float64
	Comment        []string
	Composer       []string
	ContentCreated string
	DiscNumber     int
	FirstUsed      string
	Genre          []string
	LastUsed       string
	Lyricist       []string
	Title          string
	TrackNumber    int
	Url            string
	UseCount       int
	UserRating     float64
}

func (m *Metadata) MakeMap() map[string]dbus.Variant {
	return map[string]dbus.Variant{
		"mpris:trackid":        dbus.MakeVariant(m.TrackId),
		"mpris:length":         dbus.MakeVariant(m.Length),
		"mpris:artUrl":         dbus.MakeVariant(m.ArtUrl),
		"xesam:album":          dbus.MakeVariant(m.Album),
		"xesam:albumArtist":    dbus.MakeVariant(m.AlbumArtist),
		"xesam:artist":         dbus.MakeVariant(m.Artist),
		"xesam:asText":         dbus.MakeVariant(m.AsText),
		"xesam:audioBPM":       dbus.MakeVariant(m.AudioBPM),
		"xesam:autoRating":     dbus.MakeVariant(m.AutoRating),
		"xesam:comment":        dbus.MakeVariant(m.Comment),
		"xesam:composer":       dbus.MakeVariant(m.Composer),
		"xesam:contentCreated": dbus.MakeVariant(m.ContentCreated),
		"xesam:discNumber":     dbus.MakeVariant(m.DiscNumber),
		"xesam:firstUsed":      dbus.MakeVariant(m.FirstUsed),
		"xesam:genre":          dbus.MakeVariant(m.Genre),
		"xesam:lastUsed":       dbus.MakeVariant(m.LastUsed),
		"xesam:lyricist":       dbus.MakeVariant(m.Lyricist),
		"xesam:title":          dbus.MakeVariant(m.Title),
		"xesam:trackNumber":    dbus.MakeVariant(m.TrackNumber),
		"xesam:url":            dbus.MakeVariant(m.Url),
		"xesam:useCount":       dbus.MakeVariant(m.UseCount),
		"xesam:userRating":     dbus.MakeVariant(m.UserRating),
	}
}

type PlaybackStatus string

const (
	PlaybackStatusPlaying PlaybackStatus = "Playing"
	PlaybackStatusPaused  PlaybackStatus = "Paused"
	PlaybackStatusStopped PlaybackStatus = "Stopped"
)

type LoopStatus string

const (
	LoopStatusNone     LoopStatus = "None"
	LoopStatusTrack    LoopStatus = "Track"
	LoopStatusPlaylist LoopStatus = "Playlist"
)
