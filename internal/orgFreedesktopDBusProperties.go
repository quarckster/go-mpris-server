package internal

import (
	"reflect"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/prop"
)

type iface = string
type property = string
type methodsMap = map[iface]map[property]interface{}

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
	conn, _ := dbus.SessionBus()
	return &OrgFreedesktopDBusProperties{
		getMethods: gm,
		setMethods: sm,
		conn:       conn,
	}
}

type OrgFreedesktopDBusProperties struct {
	mut        sync.RWMutex
	getMethods methodsMap
	setMethods methodsMap
	conn       *dbus.Conn
}

func (p *OrgFreedesktopDBusProperties) Get(iface string, property string) (dbus.Variant, *dbus.Error) {
	p.mut.RLock()
	defer p.mut.RUnlock()
	properties, ok := p.getMethods[iface]
	if !ok {
		return dbus.Variant{}, prop.ErrIfaceNotFound
	}
	method, ok := properties[property]
	if !ok {
		return dbus.Variant{}, prop.ErrPropNotFound
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
		return nil, prop.ErrIfaceNotFound
	}
	result := map[string]dbus.Variant{}
	for k, v := range properties {
		reflectValue := reflect.ValueOf(v).Call([]reflect.Value{})
		variant := dbus.MakeVariant(reflectValue[0].Interface())
		err, _ := reflectValue[1].Interface().(error)
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
		return prop.ErrIfaceNotFound
	}
	method, ok := properties[property]
	if !ok {
		return prop.ErrPropNotFound
	}
	arg := reflect.ValueOf(newv.Value())
	args := []reflect.Value{arg}
	// set methods should always return an error
	err, _ := reflect.ValueOf(method).Call(args)[0].Interface().(error)
	if err != nil {
		return dbus.MakeFailedError(err)
	}
	err = EmitPropertiesChanged(p.conn, iface, map[string]dbus.Variant{property: newv})
	if err != nil {
		return dbus.MakeFailedError(err)
	}
	return nil
}
