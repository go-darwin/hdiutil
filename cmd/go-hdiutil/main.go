// Copyright 2016 The darwinutils Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/go-darwin/hdiutil"
)

func main() {
	img := "/Users/zchee/.docker/machine/cache/boot2docker.iso"
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
