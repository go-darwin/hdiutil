// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import "os/exec"

// sizeFlag implements a hdiutil create command size flag interface.
type sizeFlag interface {
	sizeFlag() []string
}

// CreateSize specify the size of the image in the style of mkfile(8) with the addition of tera-, peta-, and exa-bytes sizes.
//
// The larger sizes are useful for large sparse images.
type CreateSize string

func (c CreateSize) sizeFlag() []string { return stringFlag("size", string(c)) }

// CreateSectors specify the size of the image file in 512-byte sectors.
type CreateSectors int

func (c CreateSectors) sizeFlag() []string { return intFlag("sectors", int(c)) }

// CreateMegabytes specify the size of the image file in megabytes (1024*1024 bytes).
type CreateMegabytes int

func (c CreateMegabytes) sizeFlag() []string { return intFlag("megabytes", int(c)) }

// CreateSrcfolder copies file-by-file the contents of source into image, creating a fresh (theoretically defragmented) filesystem on the destination.
//
// The resulting image is thus recommended for use with asr(8) since it will have a minimal amount of unused space.
// Its size will be that of the source data plus some padding for filesystem overhead. The filesystem type of the image volume will match that of the source as closely as possible unless overridden with -fs.
//
// Other size specifiers, such as CreateSize, will override the default size calculation based on the source content, allowing for more or less free space in the resulting filesystem.
// CreateSrcfolder can be specified more than once, in which case the image volume will be populated at the top level with a copy of each specified filesystem object.
type CreateSrcfolder string

func (c CreateSrcfolder) sizeFlag() []string { return stringFlag("srcfolder", string(c)) }

// CreateSrcdir is a synonym to CreateSrcfolder.
type CreateSrcdir CreateSrcfolder

func (c CreateSrcdir) sizeFlag() []string { return stringFlag("srcdir", string(c)) }

// CreateSrcdevice specifies that the blocks of device should be used to create a new image.
//
// The image size will match the size of device. resize can be used to adjust the size of resizable filesystems and writable images.
// Both CreateSrcdevice and CreateSrcfolder can run into errors if there are bad blocks on a disk.
// One way around this problem is to write over the files in question in the hopes that the drive will remap the bad blocks.
// Data will be lost, but the image creation operation will subsequently succeed.
//
// Filesystem options (like createFS, CreateVolname, CreateStretch, or CreateSize) are invalid and ignored when using CreateSrcdevice.
type CreateSrcdevice string

func (c CreateSrcdevice) sizeFlag() []string { return stringFlag("srcdevice", string(c)) }

// createFlag implements a hdiutil create command flag interface.
type createFlag interface {
	createFlag() []string
}

// CreateAlign specifies a size to which the final data partition will be aligned. The default is 4K.
type CreateAlign int

func (c CreateAlign) createFlag() []string { return intFlag("align", int(c)) }

type createType int

const (
	// CreateUDIF is the default type. If specified, a UDRW of the specified size will be created.
	CreateUDIF createType = 1 << iota
	// CreateSPARSE creates a UDSP, a read/write single-file image which expands as is is filled with data.
	CreateSPARSE
	// CreateSPARSEBUNDLE creates a UDSB, a read/write image backed by a directory bundle.
	CreateSPARSEBUNDLE
)

func (c createType) String() string {
	switch c {
	case CreateUDIF:
		return "UDIF"
	case CreateSPARSE:
		return "SPARSE"
	case CreateSPARSEBUNDLE:
		return "SPARSEBUNDLE"
	default:
		return ""
	}
}

func (c createType) createFlag() []string { return stringFlag("type", c.String()) }

type createFS int

const (
	// CreateHFSPlus provides the HFS+.
	CreateHFSPlus createFS = 1 << iota
	// CreateHFSPlusJ provides the HFS+J.
	CreateHFSPlusJ
	// CreateJHFSPlus provides the JHFS+.
	CreateJHFSPlus
	// CreateHFSX provides the HFSX.
	CreateHFSX
	// CreateJHFSPlusX provides the JHFS+X.
	CreateJHFSPlusX
	// CreateAPFS provides the APFS.
	CreateAPFS
	// CreateFAT32 provides the FAT32.
	CreateFAT32
	// CreateExFAT provides the ExFAT.
	CreateExFAT
	// CreateUDF provides the UDF.
	CreateUDF
)

func (c createFS) String() string {
	switch c {
	case CreateHFSPlus:
		return "HFS+"
	case CreateHFSPlusJ:
		return "HFS+J"
	case CreateJHFSPlus:
		return "JHFS+"
	case CreateHFSX:
		return "HFSX"
	case CreateJHFSPlusX:
		return "JHFS+X"
	case CreateAPFS:
		return "APFS"
	case CreateFAT32:
		return "FAT32"
	case CreateExFAT:
		return "ExFAT"
	case CreateUDF:
		return "UDF"
	default:
		return ""
	}
}

func (c createFS) createFlag() []string { return stringFlag("fs", c.String()) }

// CreateVolname the newly-created filesystem will be named volname.
//
// The default depends the filesystem being used, The default volume name in both HFS+ and APFS is `untitled'.
//
// CreateVolname is invalid and ignored when using CreateSrcdevice.
type CreateVolname string

func (c CreateVolname) createFlag() []string { return stringFlag("volname", string(c)) }

// CreateUID the root of the newly-created volume will be owned by the given numeric user id. 99 maps to the magic 'unknown' user.
type CreateUID int

func (c CreateUID) createFlag() []string { return intFlag("uid", int(c)) }

// CreateGID the root of the newly-created volume will be owned by the given numeric group id. 99 maps to 'unknown'.
type CreateGID int

func (c CreateGID) createFlag() []string { return intFlag("gid", int(c)) }

// CreateMode the root of the newly-created volume will have mode (in octal) mode.
//
// The default mode is determined by the filesystem's newfs unless CreateSrcfolder is specified, in which case the default mode is derived from the specified filesystem object.
type CreateMode string

func (c CreateMode) createFlag() []string { return stringFlag("mode", string(c)) }

type createAutostretch bool

func (c createAutostretch) createFlag() []string { return boolNoFlag("autostretch", bool(c)) }

// CreateStretch initializes HFS+ filesystem data such that it can later be stretched on older systems (which could only stretch within predefined limits) using hdiutil resize or by asr(8). max_stretch(int) is specified like CreateSize.
//
// CreateStretch is invalid and ignored when using CreateSrcdevice.
type CreateStretch int

func (c CreateStretch) createFlag() []string { return intFlag("stretch", int(c)) }

// CreateFSArgs additional arguments to pass to whichever newfs program is implied by createFS.
//
// As an example with HFS+, newfs_hfs(8) has a number of options that can control the amount of space used by the filesystem's data structures.
type CreateFSArgs []string

func (c CreateFSArgs) createFlag() []string { return stringSliceFlag("fsargs", c) }

// CreateLayout specify the partition layout of the image.
//
// layout can be anything supported by MediaKit.framework.
// NONE creates an image with no partition map.When such an image is attached, a single /dev entry will be created (e.g. /dev/disk1).
//
// 'SPUD' causes a DDM and an Apple Partition Scheme partition map with a single entry to be written. 'GPTSPUD' creates a similar image but with a GUID Partition Scheme map instead.
// When attached, multiple /dev entries will be created, with either slice 1 (GPT) or slice 2 (APM) as the data partition. (e.g. /dev/disk1, /dev/disk1s1, /dev/disk1s2).
//
// Unless overridden by createFS, the default layout is 'GPTSPUD' (PPC systems used 'SPUD' prior to Mac OS X 10.6). Other layouts include 'MBRSPUD' and 'ISOCD'. create -help lists all supported layouts.
type CreateLayout string

func (c CreateLayout) createFlag() []string { return stringFlag("layout", string(c)) }

// CreateLibrary specify an alternate layout library. The default is MediaKit's MKDrivers.bundle.
type CreateLibrary string

func (c CreateLibrary) createFlag() []string { return stringFlag("library", string(c)) }

// CreatePartitionType change the type of partition in a single-partition disk image. The default is Apple_HFS unless createFS implies otherwise.
type CreatePartitionType string

func (c CreatePartitionType) createFlag() []string { return stringFlag("partitionType", string(c)) }

type createOV bool

func (c createOV) createFlag() []string { return boolFlag("ov", bool(c)) }

type createAttach bool

func (c createAttach) createFlag() []string { return boolFlag("attach", bool(c)) }

// CreateFormat specify the final image format. The default when a source is specified is UDZO. CreateFormat can be any of the format parameters used by convert.
type CreateFormat string

func (c CreateFormat) createFlag() []string { return stringFlag("format", string(c)) }

// CreateSegmentSize specify that the image should be written in segments no bigger than size_spec (which follows CreateSize conventions).
type CreateSegmentSize int

func (c CreateSegmentSize) createFlag() []string { return intFlag("segmentSize", int(c)) }

type createCrossdev bool

func (c createCrossdev) createFlag() []string { return boolNoFlag("crossdev", bool(c)) }

type createScrub bool

func (c createScrub) createFlag() []string { return boolNoFlag("scrub", bool(c)) }

type createAnyowners bool

func (c createAnyowners) createFlag() []string { return boolNoFlag("anyowners", bool(c)) }

type createSkipunreadable bool

func (c createSkipunreadable) createFlag() []string { return boolFlag("skipunreadable", bool(c)) }

type createAtomic bool

func (c createAtomic) createFlag() []string { return boolFlag("atomic", bool(c)) }

// CreateCopyuid perform the copy as the given user. Requires root privilege.
// If user can't read or create files with the needed owners, CreateAnyowners or CreateSkipunreadable must be used to prevent the operation from failing.
type CreateCopyuid string

func (c CreateCopyuid) createFlag() []string { return stringFlag("copyuid", string(c)) }

const (
	// CreateAutostretch do suppress automatically making backwards-compatible stretchable volumes when the volume size crosses the auto-stretch-size threshold (default: 256 MB). See also asr(8).
	CreateAutostretch createAutostretch = true

	// CreateNoAutostretch do not suppress automatically making backwards-compatible stretchable volumes when the volume size crosses the auto-stretch-size threshold (default: 256 MB). See also asr(8).
	CreateNoAutostretch createAutostretch = false

	// CreateOV overwrite an existing file. The default is not to overwrite existing files.
	CreateOV createOV = true

	// CreateAttach the image after creating it. If no filesystem is specified via createFS, the attach will fail per the default attach createMount required behavior.
	CreateAttach createAttach = true

	// CreateCrossdev do cross device boundaries on the source filesystem.
	CreateCrossdev createCrossdev = true

	// CreateNoCrossdev do not cross device boundaries on the source filesystem.
	CreateNoCrossdev createCrossdev = false

	// CreateScrub do cross device boundaries on the source filesystem.
	CreateScrub createScrub = true

	// CreateNoScrub do not cross device boundaries on the source filesystem.
	CreateNoScrub createScrub = false

	// CreateAnyowners do fail if the user invoking hdiutil can't ensure correct file ownership for the files in the image.
	CreateAnyowners createAnyowners = true

	// CreateNoAnyowners do not fail if the user invoking hdiutil can't ensure correct file ownership for the files in the image.
	CreateNoAnyowners createAnyowners = false

	// CreeteSkipunreadable skip files that can't be read by the copying user and don't authenticate.
	CreeteSkipunreadable createSkipunreadable = false

	// CreateAtomic do copy files to a temporary location and then rename them to their destination. Atomic copies are the default. Non-atomic copying may be slightly faster.
	CreateAtomic createAtomic = true

	// CreateNoAtomic do not copy files to a temporary location and then rename them to their destination. Atomic copies are the default. Non-atomic copying may be slightly faster.
	CreateNoAtomic createAtomic = false
)

// Create create a new image of the given size or from the provided data.
func Create(image string, sizeSpec sizeFlag, flags ...createFlag) error {
	cmd := exec.Command(hdiutilPath, "create")
	cmd.Args = append(cmd.Args, sizeSpec.sizeFlag()...)
	cmd.Args = append(cmd.Args, image)
	if len(flags) > 0 {
		for _, flag := range flags {
			cmd.Args = append(cmd.Args, flag.createFlag()...)
		}
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
