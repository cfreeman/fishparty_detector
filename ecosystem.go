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

// import (
// 	"bitbucket.org/liamstask/gosc"
// 	"net"
// )

func updateEcosystem(deltaA chan float32, deltaL chan float32) {
	// Notify sound system via OSC.
	// DEMO OSC code.
	// m := &osc.Message{Address: "/my/message"}
	// m.Args = append(m.Args, int32(12345))
	// m.Args = append(m.Args, "important")

	// // error checking omitted for brevity...
	// addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	// conn, _ := net.DialUDP("udp", nil, addr)

	// _, err := m.WriteTo(conn)
	// if err != nil {
	// 	// handle error
	// }

	// Notify fishtank.
	// This will be a restful call to the YUN.
}
