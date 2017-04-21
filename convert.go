// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import "os/exec"

// formatFlag implements a hdiutil convert command format flag interface.
type formatFlag interface {
	formatFlag() []string
}

type convertFormot int

const (
	// ConvertUDRW UDIF read/write image.
	ConvertUDRW convertFormot = 1 << iota
	// ConvertUDRO UDIF read-only image.
	ConvertUDRO
	// ConvertUDCO UDIF ADC-compressed image.
	ConvertUDCO
	// ConvertUDZO UDIF zlib-compressed image.
	ConvertUDZO
	// ConvertULFO UDIF lzfse-compressed image (OS X 10.11+ only).
	ConvertULFO
	// ConvertUDBZ UDIF bzip2-compressed image (Mac OS X 10.4+ only).
	ConvertUDBZ
	// ConvertUDTO DVD/CD-R master for export.
	ConvertUDTO
	// ConvertUDSP SPARSE (grows with content).
	ConvertUDSP
	// ConvertUDSB SPARSEBUNDLE (grows with content; bundle-backed).
	ConvertUDSB
	// ConvertUFBI UDIF entire image with MD5 checksum.
	ConvertUFBI
	// ConvertUDRo UDIF read-only (obsolete format).
	ConvertUDRo
	// ConvertUDCo UDIF compressed (obsolete format).
	ConvertUDCo
	// ConvertRdWr NDIF read/write image (deprecated).
	ConvertRdWr
	// ConvertRdxx NDIF read-only image (Disk Copy 6.3.3 format; deprecated).
	ConvertRdxx
	// ConvertROCo NDIF compressed image (deprecated).
	ConvertROCo
	// ConvertRken NDIF compressed (obsolete format).
	ConvertRken
	// ConvertDC42 Disk Copy 4.2 image (obsolete format).
	ConvertDC42
)

// convertFlag implements a hdiutil convert command flag interface.
type convertFlag interface {
	convertFlag() []string
}

// ConvertAlign default is 4 (2K).
type ConvertAlign int

func (c ConvertAlign) convertFlag() []string { return intFlag("align", int(c)) }

type convertPmap bool

func (c convertPmap) convertFlag() []string { return boolFlag("pmap", bool(c)) }

// ConvertSegmentSize specify segmentation into size_spec-sized segments as outfile is being written.
//
// The default size_spec when ConvertSegmentSize is specified alone is 2*1024*1024 (1 GB worth of sectors) for UDTO images and 4*1024*1024 (2 GB segments) for all other image types.
//
// size_spec(string) can also be specified ??b|??k|??m|??g|??t|??p|??e like create's CreateSize flag.
type ConvertSegmentSize string

func (c ConvertSegmentSize) convertFlag() []string { return stringFlag("segmentSize", string(c)) }

// ConvertTasks when converting an image into a compressed format, specify the number of threads to use for the compression operation.
//
// The default is the number of processors active in the current system.
type ConvertTasks int

func (c ConvertTasks) convertFlag() []string { return intFlag("tasks", int(c)) }

const (
	// ConvertPmap add partition map.
	ConvertPmap convertPmap = true
)

// Convert convert image to type format and write the result to outfile.
func Convert(image string, format formatFlag, outfile string, flags ...convertFlag) error {
	cmd := exec.Command(hdiutilPath, "convert", image)
	cmd.Args = append(cmd.Args, format.formatFlag()...)
	cmd.Args = append(cmd.Args, outfile)
	if len(flags) > 0 {
		for _, flag := range flags {
			cmd.Args = append(cmd.Args, flag.convertFlag()...)
		}
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
