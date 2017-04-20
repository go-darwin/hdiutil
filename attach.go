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

// kernel attempt to attach this image without a helper process; fail if unsupported.
// Only UDRW, UDRO, UDZO, ULFO, and UDSP images are supported in-kernel.
// Encryption and HTTP are supported by the kernel driver.
// -nokernel attach with a helper process. This is (again) the default as of Mac OS X 10.5.
type kernel bool

func (k kernel) attachFlag() string { return boolNoFlag(bool(k), "kernel") }

// notremovable prevent this image from being detached. Only root can use this option.
// A reboot is necessary to cleanly detach an image attached with -notremovable.
type notRemovable bool

func (n notRemovable) attachFlag() string { return boolFlag(bool(n), "notremovable") }

// mount indicate whether filesystems in the image should be mounted or not.
// The default is required (attach will fail if no filesystems mount).
type mount string

func (m mount) attachFlag() string { return stringFlag(string(m), "mount") }

type noMount bool

func (n noMount) attachFlag() string { return boolFlag(bool(n), "nomount") }

// MountRoot mount volumes on subdirectories of path instead of under /Volumes. path must exist.
// Full mount point paths must be less than MNAMELEN characters (increased from 90 to 1024 in Mac OS X 10.6).
type MountRoot string

func (m MountRoot) attachFlag() string { return stringFlag(string(m), "mountroot") }

// MountRandom like -mountroot, but mount point directory names are randomized with mkdtemp(3).
type MountRandom string

func (m MountRandom) attachFlag() string { return stringFlag(string(m), "mountrandom") }

// MountPoint assuming only one volume, mount it at path instead of in /Volumes.
// See fstab(5) for ways a system administrator can make particular volumes automatically mount in particular filesystem locations by editing the file /etc/fstab.
type MountPoint string

func (m MountPoint) attachFlag() string { return stringFlag(string(m), "mountpoint") }

// noBrowse render any volumes invisible in applications such as the macOS Finder.
type noBrowse bool

func (n noBrowse) attachFlag() string { return boolFlag(bool(n), "nobrowse") }

// owners specify that owners on any filesystems be honored or not.
type owners string

const (
	// OwnersOn is -owners on
	ownersOn owners = "on"
	// OwnersOff is -owners off
	ownersOff owners = "off"
)

func (o owners) attachFlag() string { return stringFlag(string(o), "owners") }

// drivekey specify a key/value pair to be set on the device in the IOKit registry.
type drivekey struct {
	Key   string
	Value string
}

// section attach a subsection of a disk image.
// subspec is any of <offset>, <first-last>, or <start,count> in 0-based sectors.
// Ranges are inclusive.
type section [2]int

func (s section) attachFlag() string {
	var arg string
	for v := range s {
		arg = arg + strconv.Itoa(v)
	}
	return stringFlag(arg, "section")
}

// verify do [not] verify the image.
// By default, hdiutil attach attempts to intelligently verify images that contain checksums before attaching them.
// If hdiutil can write to an image it has verified, attach will store an attribute with
// the image so that it will not be verified again unless its timestamp changes.
// To maintain backwards compatibility, hdid(8) does not attempt to verify images before attaching them.
// Preferences keys: skip-verify, skip-verify-remote, skip-verify-locked, skip-previously-verified
type verify bool

func (v verify) attachFlag() string { return boolNoFlag(bool(v), "verify") }

// ignoreBadChecksums specify whether bad checksums should be ignored.
// The default is to abort when a bad checksum is detected.
// Preferences key: ignore-bad-checksums
type ignoreBadChecksums bool

func (i ignoreBadChecksums) attachFlag() string { return boolNoFlag(bool(i), "ignoreBadChecksums") }

// idme do perform IDME actions on IDME images.
// IDME actions are not performed by default.
// Preferences key: skip-idme
type idme bool

func (i idme) attachFlag() string { return boolNoFlag(bool(i), "idme") }

// idmeReveal do [not] reveal (in the Finder) the results of IDME processing.
// Preferences key: skip-idme-reveal
type idmeReveal bool

func (i idmeReveal) attachFlag() string { return boolNoFlag(bool(i), "idmereveal") }

// idmeTrash do [not] put IDME images in the trash after processing.
// Preferences key: skip-idme-trash
type idmeTrash bool

func (i idmeTrash) attachFlag() string { return boolNoFlag(bool(i), "idmetrash") }

// autoOpen do [not] auto-open volumes (in the Finder) after attaching an image.
// By default, double-clicking a read-only disk image causes the resulting volume to be opened in the Finder.
// hdiutil defaults to -noautoopen.
type autoOpen bool

func (a autoOpen) attachFlag() string { return boolNoFlag(bool(a), "autoopen") }

// autoOpenRO do [not] auto-open read-only volumes.
// Preferences key: auto-open-ro-root
type autoOpenRO bool

func (a autoOpenRO) attachFlag() string { return boolNoFlag(bool(a), "autoopenro") }

// autoOpenRW do [not] auto-open read/write volumes.
// Preferences key: auto-open-rw-root
type autoOpenRW bool

func (a autoOpenRW) attachFlag() string { return boolNoFlag(bool(a), "autoopenrw") }

// autoFsck do [not] force automatic file system checking before mounting a disk image.
// By default, only quarantined images (e.g. downloaded from the Internet) that have not previously passed fsck are checked.
// Preferences key: auto-fsck
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
