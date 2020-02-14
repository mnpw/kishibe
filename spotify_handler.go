package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus"
)

// List of valid messages that can be sent to Spotify
var controls = []string{"Next", "Previous", "OpenUri", "PlayPause", "Play", "Pause", "Seek"}

func isValidAction(ac string) bool {
	for _, f := range controls {
		if f == ac {
			return true
		}
	}
	return false
}

func main() {

	var action string
	// options := os.Args[2:]

	if len(os.Args) > 1 {
		action = os.Args[1]
	} else {
		fmt.Fprintln(os.Stderr, "No action passed.")
		os.Exit(1)
	}

	if isValidAction(action) == false {
		fmt.Fprintln(os.Stderr, "Invalid action passed.")
		os.Exit(1)
	}

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to SystemBus bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// var s string
	err = conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").Call("org.mpris.MediaPlayer2.Player.PlayPause", 0).Err
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to introspect bluez", err)
		os.Exit(1)
	}

	// fmt.Println(s)
}
