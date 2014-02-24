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
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func updateDispensers(tankL chan float32, config Configuration) {
	// Initialise all the Beverage dispensers in the ecosystem (1.0 is full, 0.0 is empty)
	BeverageDispensers := [4]float32{1.0, 1.0, 1.0, 1.0}

	// Generate URL endpoint for each dispenser in the ecosystem.
	for index, _ := range BeverageDispensers {
		http.HandleFunc("/"+strconv.Itoa(index)+"/", func(w http.ResponseWriter, r *http.Request) {

			// Derive the ID of the dispenser that just squrited out a beverage.
			id, err := strconv.Atoi(strings.TrimRight(strings.TrimLeft(r.URL.Path, "/"), "/"))
			if err != nil {
				fmt.Printf("ERR: Unable to parse the dispenser id\n")
				return
			}

			// Parse the new level to use for the dispenser.
			i, err := strconv.ParseFloat(r.FormValue("l"), 32)
			if err != nil {
				fmt.Printf("ERR: Unable to parse dispenser level\n")
				return
			}

			// Update the level of dispenser that was altered.
			BeverageDispensers[id] = float32(i)
			fmt.Printf("INFO: B[%d]: %f\n", id, BeverageDispensers[id])

			// ecosystem causality - determine the new level of the fish tank.
			tank := float32(0.0)
			for _, dispenserLevel := range BeverageDispensers {
				tank += dispenserLevel
			}
			tankL <- tank / float32(4.0)
		})
	}

	http.ListenAndServe(config.ListenAddress, nil)
}
