package types

import (
	"github.com/godbus/dbus/v5"
)

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
