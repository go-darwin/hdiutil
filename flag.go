// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import (
	"fmt"
)

func boolFlag(b bool, name string) string {
	if b {
		return "-" + name
	}
	return ""
}

func boolNoFlag(b bool, name string) string {
	if b {
		return "-" + name
	}
	return "-no" + name
}

func stringFlag(s, name string) string {
	return fmt.Sprintf("-%s %s", name, s)
}
