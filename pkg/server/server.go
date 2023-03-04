package server

import (
	"errors"
	"log"

	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/internal"
	"github.com/quarckster/go-mpris-server/pkg/types"
)

type Server struct {
	serviceName   string
	conn          *dbus.Conn
	rootAdapter   types.OrgMprisMediaPlayer2Adapter
	playerAdapter types.OrgMprisMediaPlayer2PlayerAdapter
	stop          chan bool
}

// Create a new server with a given name and initialize needed data.
func NewServer(
	name string,
	rootAdapter types.OrgMprisMediaPlayer2Adapter,
	playerAdapter types.OrgMprisMediaPlayer2PlayerAdapter,
) *Server {
	server := Server{
		serviceName:   "org.mpris.MediaPlayer2." + name,
		rootAdapter:   rootAdapter,
		playerAdapter: playerAdapter,
		stop:          make(chan bool, 1),
	}
	return &server
}

func (s *Server) exportMethods() error {
	root := internal.NewOrgMprisMediaPlayer2(s.rootAdapter)
	player := internal.NewOrgMprisMediaPlayer2Player(s.playerAdapter)
	properties := internal.NewOrgFreedesktopDBusProperties(s.serviceName, s.conn, root, player)
	return internal.ExportMethods(s.conn, root, player, properties)
}

// Start the server and block.
func (s *Server) Listen() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	s.conn = conn
	reply, err := s.conn.RequestName(s.serviceName, dbus.NameFlagReplaceExisting)
	if err != nil || reply != dbus.RequestNameReplyPrimaryOwner {
		s.conn.Close()
		return errors.New("Unable to claim " + s.serviceName)
	}
	err = s.exportMethods()
	if err != nil {
		s.conn.ReleaseName(s.serviceName)
		s.conn.Close()
		return err
	}
	log.Println("Started DBus server on " + s.serviceName)
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
	_, err = s.conn.ReleaseName(s.serviceName)
	if err != nil {
		s.stop <- true
		return err
	}
	err = s.conn.Close()
	if err != nil {
		s.stop <- true
		return err
	}
	log.Println("Finished " + s.serviceName)
	s.stop <- true
	return nil
}
