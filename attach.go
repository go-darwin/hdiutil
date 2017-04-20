// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

// attachFlag implements a hdiutil attach command flag interface.
type attachFlag interface {
	attachFlag() []string
}

type attachRWType int

const (
	// Readonly force the resulting device to be read-only.
	readonly attachRWType = 1 << iota
	// Readwrite attempt to override the DiskImages framework's decision to attach a particular image read-only.
	// For example, -readwrite can be used to modify the HFS+ filesystem on a HFS+/ISO hybrid CD image.
	readwrite
)

func (a attachRWType) attachFlag() string {
	switch a {
	case readonly:
		return "-readonly"
	case readwrite:
		return "-readwrite"
	default:
		return ""
	}
}

type attachKernel bool

func (a attachKernel) attachFlag() []string { return boolNoFlag(bool(a), "kernel") }

type attachNotRemovable bool

func (a attachNotRemovable) attachFlag() []string { return boolFlag(bool(a), "notremovable") }

type attachMount string

func (a attachMount) attachFlag() []string { return stringFlag(string(a), "mount") }

type attachNoMount bool

func (a attachNoMount) attachFlag() []string { return boolFlag(bool(a), "nomount") }

// AttachMountRoot mount volumes on subdirectories of path instead of under /Volumes. path must exist.
//
// Full mount point paths must be less than MNAMELEN characters (increased from 90 to 1024 in Mac OS X 10.6).
type AttachMountRoot string

func (a AttachMountRoot) attachFlag() []string { return stringFlag(string(a), "mountroot") }

// AttachMountRandom like -mountroot, but mount point directory names are randomized with mkdtemp(3).
type AttachMountRandom string

func (a AttachMountRandom) attachFlag() []string { return stringFlag(string(a), "mountrandom") }

// AttachMountPoint assuming only one volume, mount it at path instead of in /Volumes.
//
// See fstab(5) for ways a system administrator can make particular volumes automatically mount in particular filesystem locations by editing the file /etc/fstab.
type AttachMountPoint string

func (a AttachMountPoint) attachFlag() []string { return stringFlag(string(a), "mountpoint") }

type attachNoBrowse bool

func (a attachNoBrowse) attachFlag() []string { return boolFlag(bool(a), "nobrowse") }

type attachOwners string

const (
	ownersOn  attachOwners = "on"
	ownersOff attachOwners = "off"
)

func (a attachOwners) attachFlag() []string { return stringFlag(string(a), "owners") }

// AttachDrivekey specify a key/value pair to be set on the device in the IOKit registry.
type AttachDrivekey [2]string

func (a AttachDrivekey) attachFlag() []string {
	return stringFlag(a[0]+"="+a[1], "drivekey")
}

// AttachSection attach a subsection of a disk image.
// subspec is any of <offset>, <first-last>, or <start,count> in 0-based sectors.
// Ranges are inclusive.
type AttachSection [2]int

func (a AttachSection) attachFlag() []string {
	var arg string
	for v := range a {
		arg = arg + strconv.Itoa(v)
	}
	return stringFlag(arg, "section")
}

type attachVerify bool

func (a attachVerify) attachFlag() []string { return boolNoFlag(bool(a), "verify") }

type attachIgnoreBadChecksums bool

func (a attachIgnoreBadChecksums) attachFlag() []string {
	return boolNoFlag(bool(a), "ignoreBadChecksums")
}

type attachIdme bool

func (a attachIdme) attachFlag() []string { return boolNoFlag(bool(a), "idme") }

type atachIdmeReveal bool

func (a atachIdmeReveal) attachFlag() []string { return boolNoFlag(bool(a), "idmereveal") }

type attachIdmeTrash bool

func (a attachIdmeTrash) attachFlag() []string { return boolNoFlag(bool(a), "idmetrash") }

type attachAutoOpen bool

func (a attachAutoOpen) attachFlag() []string { return boolNoFlag(bool(a), "autoopen") }

type attachAutoOpenRO bool

func (a attachAutoOpenRO) attachFlag() []string { return boolNoFlag(bool(a), "autoopenro") }

type attachAutoOpenRW bool

func (a attachAutoOpenRW) attachFlag() []string { return boolNoFlag(bool(a), "autoopenrw") }

type attachAutoFsck bool

func (a attachAutoFsck) attachFlag() []string { return boolNoFlag(bool(a), "autofsck") }

const (
	// AttachReadonly force the resulting device to be read-only.
	AttachReadonly attachRWType = readonly

	// AttachReadWrite attempt to override the DiskImages framework's decision to attach a particular image read-only.
	//
	// For example, -readwrite can be used to modify the HFS+ filesystem on a HFS+/ISO hybrid CD image.
	AttachReadWrite attachRWType = readwrite

	// AttachKernel attempt to attach this image without a helper process; fail if unsupported.
	//
	// Only UDRW, UDRO, UDZO, ULFO, and UDSP images are supported in-kernel. Encryption and HTTP are supported by the kernel driver.
	AttachKernel attachKernel = true

	// AttachNoKernel attach with a helper process.  This is (again) the default as of Mac OS X 10.5.
	AttachNoKernel attachKernel = false

	// AttachNotRemovable prevent this image from being detached. Only root can use this option.
	//
	// A reboot is necessary to cleanly detach an image attached with -notremovable.
	AttachNotRemovable attachNotRemovable = true

	// AttachMountRequired indicate to -mount required.
	AttachMountRequired attachMount = "required"

	// AttachMountOptional indicate to -mount optional.
	AttachMountOptional attachMount = "optional"

	// AttachMountSuppressed indicate to -mount suppressed.
	AttachMountSuppressed attachMount = "suppressed"

	// AttachNoMount identical to -mount suppressed
	AttachNoMount attachNoMount = true

	// AttachNoBrowse render any volumes invisible in applications such as the macOS Finder.
	AttachNoBrowse attachNoBrowse = true

	// AttachOwnersOn owners on any filesystems be honored.
	AttachOwnersOn attachOwners = ownersOn

	// AttachOwnersOff owners on any filesystems be not honored.
	AttachOwnersOff attachOwners = ownersOff

	// AttachVerify do verify the image.
	AttachVerify attachVerify = true

	// AttachNoVerify do not verify the image.
	AttachNoVerify attachVerify = false

	// AttachIgnoreBadChecksums bad checksums should be ignored.
	AttachIgnoreBadChecksums attachIgnoreBadChecksums = true

	// AttachNoIgnoreBadChecksums bad checksums should be not ignored.
	AttachNoIgnoreBadChecksums attachIgnoreBadChecksums = false

	// AttachIdme do perform IDME actions on IDME images.
	AttachIdme attachIdme = true

	// AttachNoIdme do not perform IDME actions on IDME images.
	AttachNoIdme attachIdme = false

	// AttachIdmeReveal do reveal (in the Finder) the results of IDME processing.
	AttachIdmeReveal atachIdmeReveal = true

	// AttachNoIdmeReveal do not reveal (in the Finder) the results of IDME processing.
	AttachNoIdmeReveal atachIdmeReveal = false

	// AttachIdmeTrash do put IDME images in the trash after processing.
	AttachIdmeTrash attachIdmeTrash = true
	// AttachNoIdmeTrash do not put IDME images in the trash after processing.
	AttachNoIdmeTrash attachIdmeTrash = false

	// AttachAutoOpen do not auto-open volumes (in the Finder) after attaching an image.
	AttachAutoOpen attachAutoOpen = true

	// AttachNoAutoOpen do not auto-open volumes (in the Finder) after attaching an image.
	AttachNoAutoOpen attachAutoOpen = false

	// AttachAutoOpenRO do auto-open read-only volumes.
	AttachAutoOpenRO attachAutoOpenRO = true

	// AttachNoAutoOpenRO do not auto-open read-only volumes.
	AttachNoAutoOpenRO attachAutoOpenRO = false

	// AttachAutoOpenRW do auto-open read/write volumes.
	AttachAutoOpenRW attachAutoOpenRW = true

	// AttachNoAutoOpenRW do not auto-open read/write volumes.
	AttachNoAutoOpenRW attachAutoOpenRW = false

	// AttachAutoFsck do force automatic file system checking before mounting a disk image.
	AttachAutoFsck attachAutoFsck = true

	// AttachNoAutoFsck do not force automatic file system checking before mounting a disk image.
	AttachNoAutoFsck attachAutoFsck = false
)

var attachRe = regexp.MustCompile(`/dev/disk[\d]+`)

// Attach attach the image file. The returns device node path and error.
func Attach(image string, flags ...attachFlag) (string, error) {
	cmd := exec.Command(hdiutilPath, "attach", image)

	if len(flags) != 0 {
		for _, f := range flags {
			cmd.Args = append(cmd.Args, f.attachFlag()...)
		}
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, out)
	}

	return string(attachRe.Find(out)), nil
}
