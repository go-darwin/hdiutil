// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import "os/exec"

// verifyFlag implements a hdiutil verify command flag interface.
type verifyFlag interface {
	verifyFlag() []string
}

type verifyCache bool

func (v verifyCache) verifyFlag() []string { return boolFlag(bool(v), "force") }

const (
	// VerifyCache do cache checksum-verification.
	VerifyCache verifyCache = true

	// VerifyNoCache do not cache checksum-verification cache.
	VerifyNoCache verifyCache = false
)

// Verify compute the checksum of a "read-only" or "compressed" image and verify it against the value stored in the image.
func Verify(img string, flags ...verifyFlag) error {
	cmd := exec.Command(hdiutilPath, "verify", img)
	if len(flags) != 0 {
		for _, flag := range flags {
			cmd.Args = append(cmd.Args, flag.verifyFlag()...)
		}
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
