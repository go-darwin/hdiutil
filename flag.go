// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import "strconv"

func boolFlag(name string, b bool) []string {
	if b {
		return []string{"-" + name}
	}
	return nil
}

func boolNoFlag(name string, b bool) []string {
	if b {
		return []string{"-" + name}
	}
	return []string{"-no" + name}
}

func stringFlag(name, s string) []string {
	return []string{"-" + name, s}
}

func stringSliceFlag(name string, s []string) []string {
	a := []string{"-" + name}
	a = append(a, s...)
	return a
}

func intFlag(name string, i int) []string {
	return []string{"-" + name, strconv.Itoa(i)}
}
