package server

import (
	"errors"
	"log"

	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/internal"
	"github.com/quarckster/go-mpris-server/pkg/types"
)

type Server struct {
	ServiceName string
	conn        *dbus.Conn
	root        *internal.OrgMprisMediaPlayer2
	player      *internal.OrgMprisMediaPlayer2Player
	properties  *internal.OrgFreedesktopDBusProperties
	stop        chan bool
}

// Create a new server with a given name and initialize needed data.
func NewServer(
	name string,
	rootAdapter types.OrgMprisMediaPlayer2Adapter,
	playerAdapter types.OrgMprisMediaPlayer2PlayerAdapter,
) *Server {
	root := internal.NewOrgMprisMediaPlayer2(rootAdapter)
	player := internal.NewOrgMprisMediaPlayer2Player(playerAdapter)
	properties := internal.NewOrgFreedesktopDBusProperties(root, player)
	server := Server{
		ServiceName: "org.mpris.MediaPlayer2." + name,
		root:        root,
		player:      player,
		properties:  properties,
		stop:        make(chan bool, 1),
	}
	properties.EmitPropertyChanged = server.EmitPropertyChanged
	return &server
}

// Start the server and block.
func (s *Server) Listen() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	s.conn = conn
	reply, err := s.conn.RequestName(s.ServiceName, dbus.NameFlagReplaceExisting)
	if err != nil || reply != dbus.RequestNameReplyPrimaryOwner {
		s.conn.Close()
		return errors.New("Unable to claim " + s.ServiceName)
	}
	err = internal.ExportMethods(s.conn, s.root, s.player, s.properties)
	if err != nil {
		s.conn.ReleaseName(s.ServiceName)
		s.conn.Close()
		return err
	}
	log.Println("Started DBus server on " + s.ServiceName)
	<-s.stop
	return nil
}

// Release the claimed bus name and close the connection.
func (s *Server) Stop() error {
	var err error
	err = internal.UnexportMethods(s.conn)
	if err != nil {
		s.stop <- true
		return err
	}
	_, err = s.conn.ReleaseName(s.ServiceName)
	if err != nil {
		s.stop <- true
		return err
	}
	err = s.conn.Close()
	if err != nil {
		s.stop <- true
		return err
	}
	log.Println("Finished " + s.ServiceName)
	s.stop <- true
	return nil
}

// Emit sends the given signal to the bus.
func (s *Server) EmitPropertyChanged(property string, newv dbus.Variant) error {
	return s.conn.Emit(
		"/org/mpris/MediaPlayer2",
		"org.freedesktop.DBus.Properties.PropertiesChanged",
		s.ServiceName,
		map[string]dbus.Variant{property: newv},
		[]string{},
	)
}
