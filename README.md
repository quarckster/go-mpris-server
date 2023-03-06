# go-mpris-server

Go implementation of the server part of [MPRIS](https://specifications.freedesktop.org/mpris-spec/2.2/)
D-Bus interface. The library provides an event handler that emits
`org.freedesktop.DBus.Properties.PropertiesChanged` signal in response to changes in a media player.
This allows for real-time updates from the media player to D-Bus.

## Usage

### Implement adapters

Implement `pkg.types.OrgMprisMediaPlayer2Adapter` and `pkg.types.OrgMprisMediaPlayer2PlayerAdapter`
interfaces. Instances should be passed to `pkg.server.Server`.

### Integrate event handlers

Instantiate `pkg.events.EventHandler` struct and integrate with your media player to emit changes
on certain events. E.g., if the user pauses the media player, call
`pkg.events.EventHandler.OnPlayPause()` in the player's code.

### Instantiate server and listen

Instantiate the server from `pkg.server.Server` struct, pass your adapters and run it with
`pkg.server.Server.Listen()`.

### Example

```go
package main

import (
	"log"

	"github.com/quarckster/go-mpris-server/pkg/events"
	"github.com/quarckster/go-mpris-server/pkg/server"
)

type Root struct{}

// Implement other methods of `pkg.types.OrgMprisMediaPlayer2Adapter`
func (r Root) Raise() error {
	log.Println("Raised")
	return nil
}

type Player struct {}

// Implement other methods of `pkg.types.OrgMprisMediaPlayer2PlayerAdapter`
func (p Player) Next() error {
	log.Println("Next")
	return nil
}


func main() {
	r := root{}
	p := player{}
	s := server.NewServer("MyPlayer", r, p)
	eventHandler := events.NewEventHandler(s)
	go s.Listen()
	// some blocking call should be here
}
```
