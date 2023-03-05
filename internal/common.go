package internal

import "github.com/godbus/dbus/v5"

func makeError(err error) *dbus.Error {
	if err != nil {
		return dbus.MakeFailedError(err)
	}
	return nil
}

func EmitPropertiesChanged(conn *dbus.Conn, iface string, changes map[string]dbus.Variant) error {
	return conn.Emit(
		"/org/mpris/MediaPlayer2",
		"org.freedesktop.DBus.Properties.PropertiesChanged",
		iface,
		changes,
		[]string{},
	)
}

