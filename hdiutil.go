// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdiutil

import (
	"fmt"
	"strconv"
	"strings"
)

// Usage: hdiutil <verb> <options>
// <verb> is one of the following:
// help                    flatten
// attach                  imageinfo
// detach                  internet-enable
// eject                   isencrypted
// verify                  makehybrid
// create                  mount
// compact                 mountvol
// convert                 unmount
// burn                    plugins
// info                    resize
// checksum                segment
// chpass                  pmap
// erasekeys               udifderez
// unflatten               udifrez
// help                       display more detailed help

// EncryptionType specify a particular type of encryption.
type EncryptionType int

const (
	// AES128 use AES cipher running in CBC mode on 512-byte blocks with a 128-bit key (recommended).
	AES128 EncryptionType = 1 << iota
	// AES256 use AES cipher running in CBC mode on 512-byte blocks with a 256-bit key (more secure, but slower).
	AES256
)

func (e EncryptionType) String() string {
	switch e {
	case AES128:
		return "AES-128"
	case AES256:
		return "AES-256"
	}
	return fmt.Sprintf("EncryptionType(%d)", e)
}

func (e EncryptionType) attachFlag() []string     { return stringFlag("encryption", e.String()) }
func (e EncryptionType) verifyFlag() []string     { return stringFlag("encryption", e.String()) }
func (e EncryptionType) convertFlag() []string    { return stringFlag("encryption", e.String()) }
func (e EncryptionType) makehybridFlag() []string { return stringFlag("encryption", e.String()) }

type plist bool

func (p plist) attachFlag() []string  { return boolFlag("plist", bool(p)) }
func (p plist) verifyFlag() []string  { return boolFlag("plist", bool(p)) }
func (p plist) convertFlag() []string { return boolFlag("plist", bool(p)) }

type puppetstrings bool

func (p puppetstrings) attachFlag() []string     { return boolFlag("puppetstrings", bool(p)) }
func (p puppetstrings) verifyFlag() []string     { return boolFlag("puppetstrings", bool(p)) }
func (p puppetstrings) convertFlag() []string    { return boolFlag("puppetstrings", bool(p)) }
func (p puppetstrings) makehybridFlag() []string { return boolFlag("puppetstrings", bool(p)) }

// Srcimagekey specify a key/value pair for the disk image recognition system. (-imagekey is normally a synonym)
type Srcimagekey map[string]string

func (s Srcimagekey) commonFlag() []string {
	var arg string
	for k, v := range s {
		arg = k + "=" + v
	}
	return stringFlag("srcimagekey", arg)
}
func (s Srcimagekey) attachFlag() []string     { return s.commonFlag() }
func (s Srcimagekey) createFlag() []string     { return s.commonFlag() }
func (s Srcimagekey) convertFlag() []string    { return s.commonFlag() }
func (s Srcimagekey) makehybridFlag() []string { return s.commonFlag() }

// Tgtimagekey specify a key/value pair for any image created. (-imagekey is only a synonym if there is no input image).
type Tgtimagekey map[string]string

func (t Tgtimagekey) commonFlag() []string {
	var arg string
	for k, v := range t {
		arg = k + "=" + v
	}
	return stringFlag("tgtimagekey", arg)
}
func (t Tgtimagekey) attachFlag() []string  { return t.commonFlag() }
func (t Tgtimagekey) createFlag() []string  { return t.commonFlag() }
func (t Tgtimagekey) convertFlag() []string { return t.commonFlag() }

// Imagekey is normally a synonym to Srcimagekey, only a synonym Tgtimagekey if there is no input image.
type Imagekey map[string]string

func (i Imagekey) commonFlag() []string {
	var arg string
	for k, v := range i {
		arg = k + "=" + v
	}
	return stringFlag("imagekey", arg)
}
func (i Imagekey) attachFlag() []string { return i.commonFlag() }
func (i Imagekey) createFlag() []string { return i.commonFlag() }

// Encryption specify a particular type of encryption or, if not specified, the default encryption algorithm.
//
// As of OS X 10.7, the default algorithm is the AES cipher running in CBC mode on 512-byte blocks with a 128-bit key.
type Encryption EncryptionType

type stdinpass bool

func (s stdinpass) attachFlag() []string     { return boolFlag("stdinpass", bool(s)) }
func (s stdinpass) verifyFlag() []string     { return boolFlag("stdinpass", bool(s)) }
func (s stdinpass) convertFlag() []string    { return boolFlag("stdinpass", bool(s)) }
func (s stdinpass) makehybridFlag() []string { return boolFlag("stdinpass", bool(s)) }

type agentpass bool

// Recover specify a keychain containing the secret corresponding to the certificate specified with -certificate when the image was created.
type Recover string

func (r Recover) attachFlag() []string { return stringFlag("recover", string(r)) }

// Certificate specify a secondary access certificate for an encrypted image.
// cert_file must be DER-encoded certificate data, which can be created by Keychain Access or openssl(1).
type Certificate string

func (c Certificate) convertFlag() []string { return stringFlag("certificate", string(c)) }

// Pubkey specify a list of public keys, identified by their hexadecimal hashes, to be used to protect the encrypted image being created.
type Pubkey []string

// Cacert specify a certificate authority certificate.
// cert can be either a PEM file or a directory of certificates processed by c_rehash(1).
//
// See also --capath and --cacert in curl(1).
type Cacert string

type insecurehttp bool

// Shadow use a shadow file in conjunction with the data in the primary image file.
// This option prevents modification of the original image and allows read-only images to be attached read/write.
//
// When blocks are being read from the image, blocks present in the shadow file override blocks in the base image.
// All data written to an attached device will be redirected to the shadow file.
// If not specified, shadowfile defaults to image.shadow.
// If the shadow file does not exist, it is created.
//
// hdiutil verbs taking images as input accept -shadow, -cacert, and -insecurehttp.
type Shadow string

func (s Shadow) attachFlag() []string     { return stringFlag("shadow", string(s)) }
func (s Shadow) convertFlag() []string    { return stringFlag("shadow", string(s)) }
func (s Shadow) makehybridFlag() []string { return stringFlag("shadow", string(s)) }

type verbose bool

func (v verbose) attachFlag() []string     { return boolFlag("verbose", bool(v)) }
func (v verbose) detachFlag() []string     { return boolFlag("verbose", bool(v)) }
func (v verbose) createFlag() []string     { return boolFlag("verbose", bool(v)) }
func (v verbose) convertFlag() []string    { return boolFlag("verbose", bool(v)) }
func (v verbose) makehybridFlag() []string { return boolFlag("verbose", bool(v)) }

type quiet bool

func (q quiet) attachFlag() []string     { return boolFlag("quiet", bool(q)) }
func (q quiet) detachFlag() []string     { return boolFlag("quiet", bool(q)) }
func (q quiet) createFlag() []string     { return boolFlag("quiet", bool(q)) }
func (q quiet) makehybridFlag() []string { return boolFlag("quiet", bool(q)) }

type debug bool

func (d debug) attachFlag() []string     { return boolFlag("debug", bool(d)) }
func (d debug) detachFlag() []string     { return boolFlag("debug", bool(d)) }
func (d debug) createFlag() []string     { return boolFlag("debug", bool(d)) }
func (d debug) convertFlag() []string    { return boolFlag("debug", bool(d)) }
func (d debug) makehybridFlag() []string { return boolFlag("debug", bool(d)) }

const (
	// Plist provide result output in plist format.
	// Other programs invoking hdiutil are expected to use -plist rather than try to parse the human-readable output.
	//
	// The usual output is consistent but generally unstructured.
	Plist plist = true

	// Puppetstrings provide progress output that is easy for another program to parse.
	// PERCENTAGE outputs can include the value -1 which means hdiutil is performing an operation that will take an indeterminate amount of time to complete.
	//
	// Any program trying to interpret hdiutil's progress should use -puppetstrings.
	Puppetstrings puppetstrings = true

	// Stdinpass read a null-terminated passphrase from standard input.
	// If the standard input is a tty, the passphrase will be read with readpassphrase(3).
	// Otherwise, the password is read from stdin.
	//
	// -stdinpass replaces -passphrase which has been deprecated.
	// -passphrase is insecure because its argument appears in the output of ps(1) where it is visible to other users and processes on the system.
	Stdinpass stdinpass = true

	// Agentpass force the default behavior of prompting for a passphrase.
	//
	// Useful with -pubkey to create an image protected by both a passphrase and a public key.
	Agentpass agentpass = true

	// Insecurehttp ignore SSL host validation failures.
	// Useful for self-signed servers for which the appropriate certificates are unavailable or if access to a server is desired when the server name doesn't match what is in the certificate.
	Insecurehttp insecurehttp = true

	// Verbose be verbose: produce extra progress output and error diagnostics.
	//
	// This option can help the user decipher why a particular operation failed.
	// At a minimum, the probing of any specified images will be detailed.
	// BUG(zchee): not exit hdiutil command if set.
	Verbose verbose = true

	// Quiet close stdout and stderr, leaving only hdiutil's exit status to indicate success or failure.
	// No /dev entries or mount points will be printed.
	//
	// -debug and -verbose disable -quiet.
	// BUG(zchee): not get the command result such as device node path when attach.
	Quiet quiet = true

	// Debug be very verbose.
	//
	// This option is good if a large amount of progress information is needed.
	// As of Mac OS X 10.6, -debug enables -verbose.
	// BUG(zchee): not exit hdiutil command if set.
	Debug debug = true
)

// RawDeviceNode return the raw device node from the deviceNode.
func RawDeviceNode(deviceNode string) string {
	return strings.Replace(deviceNode, "disk", "rdisk", 1)
}

// DeviceNumber return the device number from the deviceNode.
func DeviceNumber(deviceNode string) int {
	n, err := strconv.Atoi(strings.TrimPrefix(deviceNode, "/dev/disk"))
	if err != nil {
		return 0
	}
	return n
}
