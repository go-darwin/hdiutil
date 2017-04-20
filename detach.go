// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import "os/exec"

// DetachFlag implements a hdiutil detach command flag interface.
type DetachFlag interface {
	detachFlag() string
}

// DetachForce ignore open files on mounted volumes, etc.
type DetachForce bool

func (d DetachForce) detachFlag() string { return boolFlag(bool(d), "force") }

// Detach detach a disk image and terminate any associated process.
func Detach(deviceNode string, flags ...DetachFlag) error {
	cmd := exec.Command(hdiutilPath, "detach", deviceNode)
	if len(flags) != 0 {
		for _, flag := range flags {
			cmd.Args = append(cmd.Args, flag.detachFlag())
		}
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
