// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"github.com/go-darwin/hdiutil"
)

func main() {
	img := "test.sparsebundle"

	if err := hdiutil.Create("test", hdiutil.CreateMegabytes(20), hdiutil.CreateHFSPlus, hdiutil.CreateSPARSEBUNDLE); err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(img); err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(img)

	deviceNode, err := hdiutil.Attach(img, hdiutil.AttachMountPoint("./test"), hdiutil.AttachNoVerify, hdiutil.AttachNoAutoFsck)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(hdiutil.RawDeviceNode(deviceNode))
	log.Println(hdiutil.DeviceNumber(deviceNode))

	if err := hdiutil.Detach(deviceNode); err != nil {
		log.Fatal(err)
	}
}
