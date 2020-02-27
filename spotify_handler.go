package main

import (
	"errors"
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

func passToSpotifyMessageBus(action string) (bool, error) {

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to SystemBus bus:", err)
		os.Exit(1)
	}

	defer conn.Close()

	sessionName := "org.mpris.MediaPlayer2.spotify"
	var objectPath dbus.ObjectPath = "/org/mpris/MediaPlayer2"
	objectInterfaceForAction := "org.mpris.MediaPlayer2.Player." + action

	sessionObject := conn.Object(sessionName, objectPath)
	actionStatus := sessionObject.Call(objectInterfaceForAction, 0).Err
	if actionStatus != nil {
		fmt.Fprintln(os.Stderr, "Failed to introspect bluez", err)
		return false, err
	}
	return true, errors.New("")

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

	passToSpotifyMessageBus(action)

}
