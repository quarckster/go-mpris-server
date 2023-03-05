package internal

import (
	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

func exportOrgFreedesktopDBusIntrospectable(conn *dbus.Conn) error {
	v := introspect.Introspectable(Spec)
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"Introspect": v.Introspect,
	}, "/org/mpris/MediaPlayer2", "org.freedesktop.DBus.Introspectable")
}

func exportOrgMprisMediaPlayer2(conn *dbus.Conn, r *OrgMprisMediaPlayer2) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"Raise": r.Raise,
		"Quit":  r.Quit,
	}, "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2")
}

func exportOrgMprisMediaPlayer2Player(conn *dbus.Conn, p *OrgMprisMediaPlayer2Player) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"Next":      p.Next,
		"Previous":  p.Previous,
		"Pause":     p.Pause,
		"PlayPause": p.PlayPause,
		"Stop":      p.Stop,
		"Play":      p.Play,
		"Seek":      p.Seek,
	}, "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2.Player")
}

func exportOrgFreedesktopDBusProperties(conn *dbus.Conn, p *OrgFreedesktopDBusProperties) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"Get":    p.Get,
		"GetAll": p.GetAll,
		"Set":    p.Set,
	}, "/org/mpris/MediaPlayer2", "org.freedesktop.DBus.Properties")
}

func ExportMethods(
	conn *dbus.Conn,
	root *OrgMprisMediaPlayer2,
	player *OrgMprisMediaPlayer2Player,
	properties *OrgFreedesktopDBusProperties,
) error {
	var err error
	err = exportOrgFreedesktopDBusIntrospectable(conn)
	if err != nil {
		return err
	}
	err = exportOrgMprisMediaPlayer2(conn, root)
	if err != nil {
		return err
	}
	err = exportOrgMprisMediaPlayer2Player(conn, player)
	if err != nil {
		return err
	}
	err = exportOrgFreedesktopDBusProperties(conn, properties)
	if err != nil {
		return err
	}
	return nil
}

func UnexportMethods(conn *dbus.Conn) error {
	var err error
	err = conn.Export(nil, "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2")
	if err != nil {
		return err
	}
	err = conn.Export(nil, "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2.Player")
	if err != nil {
		return err
	}
	err = conn.Export(nil, "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2.Playlists")
	if err != nil {
		return err
	}
	err = conn.Export(nil, "/org/mpris/MediaPlayer2", "org.mpris.MediaPlayer2.TrackList")
	if err != nil {
		return err
	}
	err = conn.Export(nil, "/org/mpris/MediaPlayer2", "org.freedesktop.DBus.Properties")
	if err != nil {
		return err
	}
	err = conn.Export(nil, "/org/mpris/MediaPlayer2", "org.freedesktop.DBus.Introspectable")
	if err != nil {
		return err
	}
	return nil
}
