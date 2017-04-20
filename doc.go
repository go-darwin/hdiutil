// Copyright 2016 The darwinutil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hdiutil implements a macOS hdiutil command wrapper.
//
//  HDIUTIL(1)              BSD General Commands Manual              HDIUTIL(1)
//  NAME
//       hdiutil -- manipulate disk images (attach, verify, create, etc)
//
//  SYNOPSIS
//       hdiutil verb [options]
//
//  DESCRIPTION
//       hdiutil uses the DiskImages framework to manipulate disk images.
//       Common verbs include attach, detach, verify, create, convert, and
//       compact.
//
//       The rest of the verbs are currently: help, info, burn, checksum,
//       chpass, erasekeys, unflatten, flatten, imageinfo, isencrypted,
//       mountvol, unmount, plugins, udifrez, udifderez, internet-enable,
//       resize, segment, makehybrid, and pmap.
//
//  BACKGROUND
//       Disk images are data containers that emulate disks.  Like disks, they
//       can be partitioned and formatted.  Many common uses of disk images
//       blur the distinction between the disk image container and its content,
//       but this distinction is critical to understanding how disk images
//       work.
//       The terms "attach" and "detach" are used to distinguish the way disk
//       images are connected to and disconnected from a system.
//       "Mount" and "unmount" are the parallel filesystems options.
//
//       For example, when you double-click a disk image in the macOS Finder,
//       two separate things happen.
//       First, the image is "attached" to the system just like an external
//       drive.
//       Then, the kernel and Disk Arbitration probe the new device for
//       recognized file structures.
//       If any are discovered that should be mounted, the associated volumes
//       will mount and appear on the desktop.
//
//       When using disk images, always consider whether an operation applies
//       to the blocks of the disk image container or to the
//       (often file-oriented) content of the image.
//       For example, hdiutil verify verifies that the blocks stored in a
//       read-only disk image have not changed since it was created.
//       It does not check whether the filesystem stored within the image is
//       self-consistent (as diskutil verifyVolume would).
//       On the other hand, hdiutil create -srcfolder creates a disk image
//       container, puts a filesystem in it, and then copies the specified
//       files to the new filesystem.
package hdiutil
