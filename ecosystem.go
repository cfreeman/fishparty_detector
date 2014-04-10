/*
 * Copyright (c) Clinton Freeman 2014
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"bitbucket.org/liamstask/gosc"
	"fmt"
	"net"
	"net/http"
)

// Representation of the audio state within Reaper.
const NUM_TRACKS int = 16

type Track struct {
	TrackID      int
	Muted        bool
	TriggerStart float32
	TriggerEnd   float32
}

func updateAudio(tracks []Track, tankLevel float32, config Configuration) []Track {
	newTrackState := make([]Track, NUM_TRACKS)

	for key, track := range tracks {
		if tankLevel >= track.TriggerStart && tankLevel < track.TriggerEnd && track.Muted {
			track.Muted = false
			notifyReaper(track, config)
		}

		// TODO: Handle the activity level distortion effects.

		newTrackState[key] = track
	}

	return newTrackState
}

func updateEcosystem(activityL chan float32, tankL chan float32, config Configuration) {
	audioTracks := make([]Track, NUM_TRACKS)

	// Setup each of the audio tracks that exist in Reaper.
	triggerFraction := (float32(1.0) / float32(16))
	for key, _ := range make([]Track, NUM_TRACKS) {
		audioTracks[key].TriggerStart = 1.0 - (float32(key+1) * triggerFraction)
		audioTracks[key].TriggerEnd = 1.0 - (float32(key) * triggerFraction)
		audioTracks[key].Muted = true
		audioTracks[key].TrackID = (key + 1)
	}
	activity := float32(0.0)
	level := float32(0.0)

	for {
		select {
		case activity = <-activityL:
			fmt.Printf("Activity: %f\n", activity)
			// TODO: Handle the activity level distortion effects.

		case level = <-tankL:
			fmt.Printf("TankL: %f\n", level)
			audioTracks = updateAudio(audioTracks, level, config)

			// Notify fishtank.
			url := fmt.Sprintf("http://%s/arduino/drain/%f", config.TankAddress, level)
			_, err := http.Get(url)
			if err != nil {
				fmt.Printf("Unable to notify fishtank YUN: " + url)
			}
		}
	}
}

func notifyReaper(track Track, config Configuration) {
	var m *osc.Message

	m = &osc.Message{Address: fmt.Sprintf("/track/%d/mute", track.TrackID)}
	if track.Muted {
		m.Args = append(m.Args, int32(1))
	} else {
		m.Args = append(m.Args, int32(0))
	}
	sendOSCMessage(m, config.OSCServerAddress)
}

func sendOSCMessage(msg *osc.Message, address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Printf("Unable to resolve OSC Server '" + address + "'\n")
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("Unable to connect to OSC Server '" + address + "'\n")
	}

	_, err = msg.WriteTo(conn)
	if err != nil {
		fmt.Printf("Unable to write to OSC Server '" + address + "'\n")
	}
}
