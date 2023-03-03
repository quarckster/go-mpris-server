package internal

import "github.com/godbus/dbus/v5"

func makeError(err error) *dbus.Error {
	if err != nil {
		return dbus.MakeFailedError(err)
	}
	return nil
}
