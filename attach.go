// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import (
	"bytes"
	"os/exec"
	"strconv"
)

// AttachFlag implements a hdiutil attach command flag interface.
type AttachFlag interface {
	attachFlag() string
}

type rwType int

const (
	// Readonly force the resulting device to be read-only.
	readonly rwType = 1 << iota
	// Readwrite attempt to override the DiskImages framework's decision to attach a particular image read-only.
	// For example, -readwrite can be used to modify the HFS+ filesystem on a HFS+/ISO hybrid CD image.
	readwrite
)

func (rw rwType) attachFlag() string {
	switch rw {
	case readonly:
		return "-readonly"
	case readwrite:
		return "-readwrite"
	default:
		return ""
	}
}

type kernel bool

func (k kernel) attachFlag() string { return boolNoFlag(bool(k), "kernel") }

type notRemovable bool

func (n notRemovable) attachFlag() string { return boolFlag(bool(n), "notremovable") }

type mount string

func (m mount) attachFlag() string { return stringFlag(string(m), "mount") }

type noMount bool

func (n noMount) attachFlag() string { return boolFlag(bool(n), "nomount") }

// AttachMountRoot mount volumes on subdirectories of path instead of under /Volumes. path must exist.
// Full mount point paths must be less than MNAMELEN characters (increased from 90 to 1024 in Mac OS X 10.6).
type AttachMountRoot string

func (m AttachMountRoot) attachFlag() string { return stringFlag(string(m), "mountroot") }

// AttachMountRandom like -mountroot, but mount point directory names are randomized with mkdtemp(3).
type AttachMountRandom string

func (m AttachMountRandom) attachFlag() string { return stringFlag(string(m), "mountrandom") }

// AttachMountPoint assuming only one volume, mount it at path instead of in /Volumes.
// See fstab(5) for ways a system administrator can make particular volumes automatically mount in particular filesystem locations by editing the file /etc/fstab.
type AttachMountPoint string

func (m AttachMountPoint) attachFlag() string { return stringFlag(string(m), "mountpoint") }

type noBrowse bool

func (n noBrowse) attachFlag() string { return boolFlag(bool(n), "nobrowse") }

type owners string

const (
	ownersOn  owners = "on"
	ownersOff owners = "off"
)

func (o owners) attachFlag() string { return stringFlag(string(o), "owners") }

// AttachDrivekey specify a key/value pair to be set on the device in the IOKit registry.
type AttachDrivekey [2]string

func (d AttachDrivekey) attachFlag() string {
	return stringFlag(d[0]+"="+d[1], "drivekey")
}

// AttachSection attach a subsection of a disk image.
// subspec is any of <offset>, <first-last>, or <start,count> in 0-based sectors.
// Ranges are inclusive.
type AttachSection [2]int

func (s AttachSection) attachFlag() string {
	var arg string
	for v := range s {
		arg = arg + strconv.Itoa(v)
	}
	return stringFlag(arg, "section")
}

type verify bool

func (v verify) attachFlag() string { return boolNoFlag(bool(v), "verify") }

type ignoreBadChecksums bool

func (i ignoreBadChecksums) attachFlag() string { return boolNoFlag(bool(i), "ignoreBadChecksums") }

type idme bool

func (i idme) attachFlag() string { return boolNoFlag(bool(i), "idme") }

type idmeReveal bool

func (i idmeReveal) attachFlag() string { return boolNoFlag(bool(i), "idmereveal") }

type idmeTrash bool

func (i idmeTrash) attachFlag() string { return boolNoFlag(bool(i), "idmetrash") }

type autoOpen bool

func (a autoOpen) attachFlag() string { return boolNoFlag(bool(a), "autoopen") }

type autoOpenRO bool

func (a autoOpenRO) attachFlag() string { return boolNoFlag(bool(a), "autoopenro") }

type autoOpenRW bool

func (a autoOpenRW) attachFlag() string { return boolNoFlag(bool(a), "autoopenrw") }

type autoFsck bool

func (a autoFsck) attachFlag() string { return boolNoFlag(bool(a), "autofsck") }

const (
	// AttachReadonly force the resulting device to be read-only.
	AttachReadonly rwType = readonly

	// AttachReadWrite attempt to override the DiskImages framework's decision to attach a particular image read-only.
	//
	// For example, -readwrite can be used to modify the HFS+ filesystem on a HFS+/ISO hybrid CD image.
	AttachReadWrite rwType = readwrite

	// AttachKernel attempt to attach this image without a helper process; fail if unsupported.
	//
	// Only UDRW, UDRO, UDZO, ULFO, and UDSP images are supported in-kernel. Encryption and HTTP are supported by the kernel driver.
	AttachKernel kernel = true

	// AttachNoKernel attach with a helper process.  This is (again) the default as of Mac OS X 10.5.
	AttachNoKernel kernel = false

	// AttachNotRemovable prevent this image from being detached. Only root can use this option.
	//
	// A reboot is necessary to cleanly detach an image attached with -notremovable.
	AttachNotRemovable notRemovable = true

	// AttachMountRequired indicate to -mount required.
	AttachMountRequired mount = "required"

	// AttachMountOptional indicate to -mount optional.
	AttachMountOptional mount = "optional"

	// AttachMountSuppressed indicate to -mount suppressed.
	AttachMountSuppressed mount = "suppressed"

	// AttachNoMount identical to -mount suppressed
	AttachNoMount noMount = true

	// AttachNoBrowse render any volumes invisible in applications such as the macOS Finder.
	AttachNoBrowse noBrowse = true

	// AttachOwnersOn owners on any filesystems be honored.
	AttachOwnersOn owners = ownersOn

	// AttachOwnersOff owners on any filesystems be not honored.
	AttachOwnersOff owners = ownersOff

	// AttachVerify do verify the image.
	AttachVerify verify = true

	// AttachNoVerify do not verify the image.
	AttachNoVerify verify = false

	// AttachIgnoreBadChecksums bad checksums should be ignored.
	AttachIgnoreBadChecksums ignoreBadChecksums = true

	// AttachNoIgnoreBadChecksums bad checksums should be not ignored.
	AttachNoIgnoreBadChecksums ignoreBadChecksums = false

	// AttachIdme do perform IDME actions on IDME images.
	AttachIdme idme = true

	// AttachNoIdme do not perform IDME actions on IDME images.
	AttachNoIdme idme = false

	// AttachIdmeReveal do reveal (in the Finder) the results of IDME processing.
	AttachIdmeReveal idmeReveal = true

	// AttachNoIdmeReveal do not reveal (in the Finder) the results of IDME processing.
	AttachNoIdmeReveal idmeReveal = false

	// AttachIdmeTrash do put IDME images in the trash after processing.
	AttachIdmeTrash idmeTrash = true
	// AttachNoIdmeTrash do not put IDME images in the trash after processing.
	AttachNoIdmeTrash idmeTrash = false

	// AttachAutoOpen do not auto-open volumes (in the Finder) after attaching an image.
	AttachAutoOpen autoOpen = true

	// AttachNoAutoOpen do not auto-open volumes (in the Finder) after attaching an image.
	AttachNoAutoOpen autoOpen = false

	// AttachAutoOpenRO do auto-open read-only volumes.
	AttachAutoOpenRO autoOpenRO = true

	// AttachNoAutoOpenRO do not auto-open read-only volumes.
	AttachNoAutoOpenRO autoOpenRO = false

	// AttachAutoOpenRW do auto-open read/write volumes.
	AttachAutoOpenRW autoOpenRW = true

	// AttachNoAutoOpenRW do not auto-open read/write volumes.
	AttachNoAutoOpenRW autoOpenRW = false

	// AttachAutoFsck do force automatic file system checking before mounting a disk image.
	AttachAutoFsck autoFsck = true

	// AttachNoAutoFsck do not force automatic file system checking before mounting a disk image.
	AttachNoAutoFsck autoFsck = false
)

// Attach attach the image file. The returns device node path and error.
func Attach(image string, flags ...AttachFlag) (string, error) {
	cmd := exec.Command(hdiutilPath, "attach", image)

	if len(flags) != 0 {
		args := []string{}
		for _, f := range flags {
			args = append(args, f.attachFlag())
		}
		cmd.Args = append(cmd.Args, args...)
	}

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(out)), nil
}
