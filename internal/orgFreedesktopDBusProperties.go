package internal

import (
	"reflect"
	"sync"

	"github.com/godbus/dbus/v5"
)

type iface = string
type property = string
type methodsMap = map[iface]map[property]interface{}

// ErrIfaceNotFound is the error returned to peers who try to access properties
// on interfaces that aren't found.
var ErrIfaceNotFound = dbus.NewError("org.freedesktop.DBus.Properties.Error.InterfaceNotFound", nil)

// ErrPropNotFound is the error returned to peers trying to access properties
// that aren't found.
var ErrPropNotFound = dbus.NewError("org.freedesktop.DBus.Properties.Error.PropertyNotFound", nil)

// ErrReadOnly is the error returned to peers trying to set a read-only
// property.
var ErrReadOnly = dbus.NewError("org.freedesktop.DBus.Properties.Error.ReadOnly", nil)

// ErrInvalidArg is returned to peers if the type of the property that is being
// changed and the argument don't match.
var ErrInvalidArg = dbus.NewError("org.freedesktop.DBus.Properties.Error.InvalidArg", nil)

func NewOrgFreedesktopDBusProperties(
	root *OrgMprisMediaPlayer2,
	player *OrgMprisMediaPlayer2Player,
) *OrgFreedesktopDBusProperties {
	gm := make(methodsMap)
	gm["org.mpris.MediaPlayer2"] = root.GetMethods()
	gm["org.mpris.MediaPlayer2.Player"] = player.GetMethods()
	sm := make(methodsMap)
	sm["org.mpris.MediaPlayer2"] = root.SetMethods()
	sm["org.mpris.MediaPlayer2.Player"] = player.SetMethods()
	return &OrgFreedesktopDBusProperties{getMethods: gm, setMethods: sm}
}

type OrgFreedesktopDBusProperties struct {
	mut        sync.RWMutex
	getMethods methodsMap
	setMethods methodsMap
	Emit       func(string, dbus.Variant) error
}

func (p *OrgFreedesktopDBusProperties) Get(iface string, property string) (dbus.Variant, *dbus.Error) {
	p.mut.RLock()
	defer p.mut.RUnlock()
	properties, ok := p.getMethods[iface]
	if !ok {
		return dbus.Variant{}, ErrIfaceNotFound
	}
	method, ok := properties[property]
	if !ok {
		return dbus.Variant{}, ErrPropNotFound
	}
	reflectValue := reflect.ValueOf(method).Call([]reflect.Value{})
	// get methods should return a value and an error
	variant := dbus.MakeVariant(reflectValue[0].Interface())
	err, _ := reflectValue[1].Interface().(error)
	if err != nil {
		return dbus.Variant{}, dbus.MakeFailedError(err)
	}
	return variant, nil
}

func (p *OrgFreedesktopDBusProperties) GetAll(iface string) (map[string]dbus.Variant, *dbus.Error) {
	p.mut.RLock()
	defer p.mut.RUnlock()
	properties, ok := p.getMethods[iface]
	if !ok {
		return nil, ErrIfaceNotFound
	}
	result := make(map[string]dbus.Variant, len(properties))
	var err error
	for k, v := range properties {
		reflectValue := reflect.ValueOf(v).Call([]reflect.Value{})
		variant := dbus.MakeVariant(reflectValue[0].Interface())
		err, _ = reflectValue[1].Interface().(error)
		if err != nil {
			return map[string]dbus.Variant{}, dbus.MakeFailedError(err)
		}
		result[k] = variant
	}
	return result, nil
}

func (p *OrgFreedesktopDBusProperties) Set(iface string, property string, newv dbus.Variant) *dbus.Error {
	p.mut.Lock()
	defer p.mut.Unlock()
	properties, ok := p.setMethods[iface]
	if !ok {
		return ErrIfaceNotFound
	}
	method, ok := properties[property]
	if !ok {
		return ErrPropNotFound
	}
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(newv.Value())
	// set methods should always return an error
	err, _ := reflect.ValueOf(method).Call(args)[0].Interface().(error)
	if err != nil {
		return dbus.MakeFailedError(err)
	}
	err = p.Emit(property, newv)
	if err != nil {
		return dbus.MakeFailedError(err)
	}
	return nil
}
