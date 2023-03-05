package internal

import (
	"reflect"

	"github.com/godbus/dbus/v5"
)

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

func Changes(adapter interface{}, props []string) (map[string]dbus.Variant, error) {
	changes := map[string]dbus.Variant{}
	for _, prop := range props {
		reflectValues := reflect.ValueOf(adapter).MethodByName(prop).Call([]reflect.Value{})
		variant := dbus.MakeVariant(reflectValues[0].Interface())
		err, _ := reflectValues[1].Interface().(error)
		if err != nil {
			return map[string]dbus.Variant{}, err
		}
		changes[prop] = variant
	}
	return changes, nil
}
