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

import "testing"

func TestMissingConfiguration(t *testing.T) {
	config, err := parseConfiguration("foo")
	if err == nil {
		t.Errorf("error not raised for invalid configuration file.")
	}

	if config.OpticalFlowScale != 300.0 {
		t.Errorf("incorrect default optical flow scale.")
	}

	if config.MovementThreshold != 1.0 {
		t.Errorf("incorrect default movement threshold.")
	}

	if config.ListenAddress != ":8080" {
		t.Errorf("incorrect default listen address")
	}
}

func TestValidConfiguration(t *testing.T) {
	config, err := parseConfiguration("testdata/test-config.json")
	if err != nil {
		t.Errorf("returned error when parsing valid configuration file")
	}

	if config.OpticalFlowScale != 0.23 {
		t.Errorf("parsed incorrect value for optical flow scale.")
	}

	if config.MovementThreshold != 0.10 {
		t.Errorf("parsed incorrect value for movement threshold.")
	}

	if config.ListenAddress != "10.1.1.1:8080" {
		t.Errorf("parsed incorrect listen address")
	}
}
