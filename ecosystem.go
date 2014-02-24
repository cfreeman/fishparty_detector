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
)

func updateEcosystem(activityL chan float32, tankL chan float32, config Configuration) {
	for {
		// Notify sound system via OSC.
		var m *osc.Message

		select {
		case activity := <-activityL:
			fmt.Printf("Activity %f\n", activity)
			m = &osc.Message{Address: "/tank/activity"}
			m.Args = append(m.Args, activity)

		case level := <-tankL:
			fmt.Printf("Level %f\n", level)
			m = &osc.Message{Address: "/tank/level"}
			m.Args = append(m.Args, level)
		}

		addr, err := net.ResolveUDPAddr("udp", config.OSCServerAddress)
		if err != nil {
			fmt.Printf("Unable to resolve OSC Server '" + config.OSCServerAddress + "'\n")
		}

		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			fmt.Printf("Unable to connect to OSC Server '" + config.OSCServerAddress + "'\n")
		}

		_, err = m.WriteTo(conn)
		if err != nil {
			fmt.Printf("Unable to write to OSC Server '" + config.OSCServerAddress + "'\n")
		}

		// Notify fishtank.
		// This will be a restful call to the YUN.
	}
}
