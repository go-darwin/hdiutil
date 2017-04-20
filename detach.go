// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import "os/exec"

type DetachFlag interface {
	detachFlag() string
}

type DetachForce bool

func (d DetachForce) detachFlag() string { return boolFlag(bool(d), "force") }

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
