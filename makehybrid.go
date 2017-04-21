// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import "os/exec"

// makehybridFlag implements a hdiutil makehybrid command flag interface.
type makehybridFlag interface {
	makehybridFlag() []string
}

type makehybridHFS bool

func (m makehybridHFS) makehybridFlag() []string { return boolFlag("hfs", bool(m)) }

type makehybridISO bool

func (m makehybridISO) makehybridFlag() []string { return boolFlag("iso", bool(m)) }

type makehybridJoliet bool

func (m makehybridJoliet) makehybridFlag() []string { return boolFlag("joliet", bool(m)) }

type makehybridUDF bool

func (m makehybridUDF) makehybridFlag() []string { return boolFlag("udf", bool(m)) }

type makehybridHFSBlessedDirectory bool

func (m makehybridHFSBlessedDirectory) makehybridFlag() []string {
	return boolFlag("hfs-blessed-directory", bool(m))
}

type makehybridHFSOpenfolder bool

func (m makehybridHFSOpenfolder) makehybridFlag() []string { return boolFlag("hfs-openfolder", bool(m)) }

type makehybridHFSStartupfileSize bool

func (m makehybridHFSStartupfileSize) makehybridFlag() []string {
	return boolFlag("hfs-startupfile-size", bool(m))
}

type makehybridAbstractFile bool

func (m makehybridAbstractFile) makehybridFlag() []string { return boolFlag("abstract-file", bool(m)) }

type makehybridBibliographyFile bool

func (m makehybridBibliographyFile) makehybridFlag() []string {
	return boolFlag("bibliography-file", bool(m))
}

type makehybridCopyrightFile bool

func (m makehybridCopyrightFile) makehybridFlag() []string { return boolFlag("copyright-file", bool(m)) }

type makehybridApplication bool

func (m makehybridApplication) makehybridFlag() []string { return boolFlag("application", bool(m)) }

type makehybridPreparer bool

func (m makehybridPreparer) makehybridFlag() []string { return boolFlag("preparer", bool(m)) }

type makehybridPublisher bool

func (m makehybridPublisher) makehybridFlag() []string { return boolFlag("publisher", bool(m)) }

type makehybridSystemID bool

func (m makehybridSystemID) makehybridFlag() []string { return boolFlag("system-id", bool(m)) }

type makehybridKeepMacSpecific bool

func (m makehybridKeepMacSpecific) makehybridFlag() []string {
	return boolFlag("keep-mac-specific", bool(m))
}

type makehybridEltoritoBoot bool

func (m makehybridEltoritoBoot) makehybridFlag() []string { return boolFlag("eltorito-boot", bool(m)) }

type makehybridHardDiskBoot bool

func (m makehybridHardDiskBoot) makehybridFlag() []string { return boolFlag("hard-disk-boot", bool(m)) }

type makehybridNoEmulBoot bool

func (m makehybridNoEmulBoot) makehybridFlag() []string { return boolFlag("no-emul-boot", bool(m)) }

type makehybridNoBoot bool

func (m makehybridNoBoot) makehybridFlag() []string { return boolFlag("no-boot", bool(m)) }

type makehybridBootLoadSeg bool

func (m makehybridBootLoadSeg) makehybridFlag() []string { return boolFlag("boot-load-seg", bool(m)) }

type makehybridBootLoadSize bool

func (m makehybridBootLoadSize) makehybridFlag() []string { return boolFlag("boot-load-seg", bool(m)) }

type makehybridEltoritoPlatform bool

func (m makehybridEltoritoPlatform) makehybridFlag() []string {
	return boolFlag("eltorito-platform", bool(m))
}

type makehybridEltoritoSpecification bool

func (m makehybridEltoritoSpecification) makehybridFlag() []string {
	return boolFlag("eltorito-specification", bool(m))
}

type makehybridUDFVersion bool

func (m makehybridUDFVersion) makehybridFlag() []string { return boolFlag("udf-version", bool(m)) }

type makehybridDefaultVolumeName bool

func (m makehybridDefaultVolumeName) makehybridFlag() []string {
	return boolFlag("default-volume-name", bool(m))
}

type makehybridHFSVolumeName bool

func (m makehybridHFSVolumeName) makehybridFlag() []string {
	return boolFlag("hfs-volume-name", bool(m))
}

type makehybridISOVolumeName bool

func (m makehybridISOVolumeName) makehybridFlag() []string {
	return boolFlag("iso-volume-name", bool(m))
}

type makehybridJolietVolumeName bool

func (m makehybridJolietVolumeName) makehybridFlag() []string {
	return boolFlag("joliet-volume-name", bool(m))
}

type makehybridUDFVolumeName bool

func (m makehybridUDFVolumeName) makehybridFlag() []string {
	return boolFlag("udf-volume-name", bool(m))
}

type makehybridHideAll bool

func (m makehybridHideAll) makehybridFlag() []string { return boolFlag("hide-all", bool(m)) }

type makehybridHideHFS bool

func (m makehybridHideHFS) makehybridFlag() []string { return boolFlag("hide-hfs", bool(m)) }

type makehybridHideISO bool

func (m makehybridHideISO) makehybridFlag() []string { return boolFlag("hide-iso", bool(m)) }

type makehybridHideJoliet bool

func (m makehybridHideJoliet) makehybridFlag() []string { return boolFlag("hide-joliet", bool(m)) }

type makehybridHideUDF bool

func (m makehybridHideUDF) makehybridFlag() []string { return boolFlag("hide-udf", bool(m)) }

type makehybridOnlyUDF bool

func (m makehybridOnlyUDF) makehybridFlag() []string { return boolFlag("only-udf", bool(m)) }

type makehybridOnlyISO bool

func (m makehybridOnlyISO) makehybridFlag() []string { return boolFlag("only-iso", bool(m)) }

type makehybridOnlyJoliet bool

func (m makehybridOnlyJoliet) makehybridFlag() []string { return boolFlag("only-joliet", bool(m)) }

type makehybridPrintSize bool

func (m makehybridPrintSize) makehybridFlag() []string { return boolFlag("print-size", bool(m)) }

type makehybridPlistin bool

func (m makehybridPlistin) makehybridFlag() []string { return boolFlag("plistin", bool(m)) }

const (
	// MakehybridHFS generate an HFS+ filesystem.
	//
	// This filesystem can be present on an image simultaneously with an ISO9660 or Joliet or UDF filesystem.
	// On operating systems that understand HFS+ as well as ISO9660 and UDF, like Mac OS 9 or OS X, HFS+ is usually the preferred filesystem for hybrid images.
	MakehybridHFS makehybridHFS = true

	// MakehybridISO generate an ISO9660 Level 2 filesystem with Rock Ridge extensions.
	//
	// This filesystem can be present on an image simultaneously with an HFS+ or Joliet or UDF filesystem.
	// ISO9660 is the standard cross-platform interchange format for CDs and some DVDs, and is understood by virtually all operating systems.
	//
	// If an ISO9660 or Joliet filesystem is present on a disk image or CD, but not HFS+, OS X will use the ISO9660 (or Joliet) filesystem.
	MakehybridISO makehybridISO = true

	// MakeHybridJoliet generate joliet extensions to ISO9660.
	//
	// This view of the filesystem can be present on an image simultaneously with HFS+, and requires the presence of an ISO9660 filesystem.
	// Joliet supports Unicode filenames, but is only supported on some operating systems.
	//
	// If both an ISO9660 and Joliet filesystem are present on a disk image or CD, but not HFS+, OS X will prefer the Joliet filesystem.
	MakeHybridJoliet makehybridJoliet = true

	// MakeHybridUDF generate a UDF filesystem.
	//
	// This filesystem can be present on an image simultaneously with HFS+, ISO9660, and Joliet.
	// UDF is the standard interchange format for DVDs, although operating system support varies based on OS version and UDF version.
	MakeHybridUDF makehybridUDF = true

	// MakehybridHFSBlessedDirectory path to directory which should be "blessed" for OS X booting on the generated filesystem.
	//
	// This assumes the directory has been otherwise prepared, for example with bless -bootinfo to create a valid BootX file. (HFS+ only).
	MakehybridHFSBlessedDirectory makehybridHFSBlessedDirectory = true

	// MakehybridHFSOpenfolder path to a directory that will be opened by the Finder automatically.  See also the -openfolder option in bless(8) (HFS+ only).
	MakehybridHFSOpenfolder makehybridHFSOpenfolder = true

	// MakehybridHFSStartupfileSize allocate an empty HFS+ Startup File of the specified size, in bytes (HFS+ only).
	MakehybridHFSStartupfileSize makehybridHFSStartupfileSize = true

	// MakehybridAbstractFile path to a file in the source directory (and thus the root of the generated filesystem) for use as the ISO9660/Joliet Abstract file (ISO9660/Joliet).
	MakehybridAbstractFile makehybridAbstractFile = true

	// MakehybridBibliographyFile path to a file in the source directory (and thus the root of the generated filesystem) for use as the ISO9660/Joliet Bibliography file (ISO9660/Joliet).
	MakehybridBibliographyFile makehybridBibliographyFile = true

	// MakehybridCopyrightFile path to a file in the source directory (and thus the root of the generated filesystem) for use as the ISO9660/Joliet Copyright file (ISO9660/Joliet).
	MakehybridCopyrightFile makehybridCopyrightFile = true

	// MakehybridApplication Application string (ISO9660/Joliet).
	MakehybridApplication makehybridApplication = true

	// MakehybridPreparer preparer string (ISO9660/Joliet).
	MakehybridPreparer makehybridPreparer = true

	// MakehybridPublisher publisher string (ISO9660/Joliet).
	MakehybridPublisher makehybridPublisher = true

	// MakehybridSystemID system Identification string (ISO9660/Joliet).
	MakehybridSystemID makehybridSystemID = true

	// MakehybridKeepMacSpecific Expose Macintosh-specific files (such as .DS_Store) in non-HFS+ filesystems (ISO9660/Joliet).
	MakehybridKeepMacSpecific makehybridKeepMacSpecific = true

	// MakehybridEltoritoBoot path to an El Torito boot image within the source directory. By default, floppy drive emulation is used, so the image must be one of 1200KB, 1440KB, or 2880KB. If the image has a different size, either -no-emul-boot or
	// -hard-disk-boot must be used to enable "No Emulation" or "Hard Disk Emulation" mode, respectively (ISO9660/Joliet).
	MakehybridEltoritoBoot makehybridEltoritoBoot = true

	// MakehybridHardDiskBoot use El Torito Hard Disk Emulation mode. The image must represent a virtual device with an MBR partition map and a single partition.
	MakehybridHardDiskBoot makehybridHardDiskBoot = true

	// MakehybridNoEmulBoot use El Torito No Emulation mode. The system firmware will load the number of sectors specified by -boot-load-size and execute it, without emulating any devices (ISO9660/Joliet).
	MakehybridNoEmulBoot makehybridNoEmulBoot = true

	// MakehybridNoBoot mark the El Torito image as non-bootable. The system firmware may still create a virtual device backed by this data. This option is not recommended (ISO9660/Joliet).
	MakehybridNoBoot makehybridNoBoot = true

	// MakehybridBootLoadSeg for a No Emulation boot image, load the data at the specified segment address.  This options is not recommended, so that the system firmware can use its default address (ISO9660/Joliet)
	MakehybridBootLoadSeg makehybridBootLoadSeg = true

	// MakehybridBootLoadSize for a No Emulation boot image, load the specified number of 512-byte emulated sectors into memory and execute it. By default, 4 sectors (2KB) will be loaded (ISO9660/Joliet).
	MakehybridBootLoadSize makehybridBootLoadSize = true

	// MakehybridEltoritoPlatform use the specified numeric platform ID in the El Torito Boot Catalog Validation Entry or Section Header. Defaults to 0 to identify x86 hardware (ISO/Joliet).
	MakehybridEltoritoPlatform makehybridEltoritoPlatform = true

	// MakehybridEltoritoSpecification for complex layouts involving multiple boot images, a plist-formatted string can be provided, using either OpenStep-style syntax or XML syntax, representing an array of dictionaries.
	//
	// Any of the El Torito options can be set in the sub-dictionaries and will apply to that boot image only.
	// If -eltorito-specification is provided in addition to the normal El Torito command-line options, the specification will be used to populate secondary non-default boot entries.
	MakehybridEltoritoSpecification makehybridEltoritoSpecification = true

	// MakehybridUDFVersion version of UDF filesystem to generate. This can be either "1.02" or "1.50".  If not specified, it defaults to "1.50" (UDF).
	MakehybridUDFVersion makehybridUDFVersion = true

	// MakehybridDefaultVolumeName default volume name for all filesystems, unless overridden.
	//
	// If not specified, defaults to the last path component of source.
	MakehybridDefaultVolumeName makehybridDefaultVolumeName = true

	// MakehybridHFSVolumeName volume name for just the HFS+ filesystem if it should be different (HFS+ only).
	MakehybridHFSVolumeName makehybridHFSVolumeName = true

	// MakehybridISOVolumeName volume name for just the ISO9660 filesystem if it should be different (ISO9660 only).
	MakehybridISOVolumeName makehybridISOVolumeName = true

	// MakehybridJolietVolumeName volume name for just the Joliet filesystem if it should be different (Joliet only).
	MakehybridJolietVolumeName makehybridJolietVolumeName = true

	// MakehybridUDFVolumeName volume name for just the UDF filesystem if it should be different (UDF only).
	MakehybridUDFVolumeName makehybridUDFVolumeName = true

	// MakehybridHideAll a glob expression of files and directories that should not be exposed in the generated filesystems.
	//
	// The string may need to be quoted to avoid shell expansion, and will be passed to glob(3) for evaluation.
	// Although this option can not be used multiple times, an arbitrarily complex glob expression can be used.
	MakehybridHideAll makehybridHideAll = true

	// MakehybridHideHFS a glob expression of files and directories that should not be exposed via the HFS+ filesystem, although the data may still be present for use by other filesystems (HFS+ only).
	MakehybridHideHFS makehybridHideHFS = true

	// MakehybridHideISO a glob expression of files and directories that should not be exposed via the ISO filesystem, although the data may still be present for use by other filesystems (ISO9660 only).
	//
	// Per above, the Joliet hierarchy will supersede the ISO hierarchy when the hybrid is mounted as an ISO 9660 filesystem on OS X.
	// Therefore, if Joliet is being generated (the default) -hide-joliet will also be needed to hide the file from mount_cd9660(8).
	MakehybridHideISO makehybridHideISO = true

	// MakehybridHideJoliet a glob expression of files and directories that should not be exposed via the Joliet filesystem, although the data may still be present for use by other filesystems (Joliet only).
	//
	// Because OS X's ISO 9660 filesystem uses the Joliet catalog if it is available, -hide-joliet effectively supersedes -hide-iso when the resulting filesystem is mounted as ISO on OS X.
	MakehybridHideJoliet makehybridHideJoliet = true

	// MakehybridHideUDF a glob expression of files and directories that should not be exposed via the UDF filesystem, although the data may still be present for use by other filesystems (UDF only).
	MakehybridHideUDF makehybridHideUDF = true

	// MakehybridOnlyUDF a glob expression of objects that should only be exposed in UDF.
	MakehybridOnlyUDF makehybridOnlyUDF = true

	// MakehybridOnlyISO a glob expression of objects that should only be exposed in ISO.
	MakehybridOnlyISO makehybridOnlyISO = true

	// MakehybridOnlyJoliet a glob expression of objects that should only be exposed in Joleit.
	MakehybridOnlyJoliet makehybridOnlyJoliet = true

	// MakehybridPrintSize preflight the data and calculate an upper bound on the size of the image.  The actual size of the generated image is guaranteed to be less than or equal to this estimate.
	MakehybridPrintSize makehybridPrintSize = true

	// MakehybridPlistin instead of using command-line parameters, use a standard plist from standard input to specific the parameters of the hybrid image generation.
	//
	// Each command-line option should be a key in the dictionary, without the leading "-", and the value should be a string for path and string arguments, a number for number arguments, and a boolean for toggle options.
	// The source argument should use a key of "source" and the image should use a key of "output".
	MakehybridPlistin makehybridPlistin = true
)

// Makehybrid generate a potentially-hybrid filesystem in a read-only disk image using the DiscRecording framework's content creation system.
func Makehybrid(image, source string, flags ...makehybridFlag) error {
	cmd := exec.Command(hdiutilPath, "makehybrid", image, source)
	if len(flags) > 0 {
		for _, flag := range flags {
			cmd.Args = append(cmd.Args, flag.makehybridFlag()...)
		}
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
