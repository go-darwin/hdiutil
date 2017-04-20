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

func (e EncryptionType) attachFlag() string { return e.String() }

// Plist provide result output in plist format.
// Other programs invoking hdiutil are expected to use -plist rather than try to parse the human-readable output.
//
// The usual output is consistent but generally unstructured.
type Plist bool

func (p Plist) attachFlag() []string { return boolFlag(bool(p), "plist") }

// Puppetstrings provide progress output that is easy for another program to parse.
// PERCENTAGE outputs can include the value -1 which means hdiutil is performing an operation that will take an indeterminate amount of time to complete.
//
// Any program trying to interpret hdiutil's progress should use -puppetstrings.
type Puppetstrings bool

func (p Puppetstrings) attachFlag() []string { return boolFlag(bool(p), "puppetstrings") }

// Srcimagekey specify a key/value pair for the disk image recognition system. (-imagekey is normally a synonym)
type Srcimagekey map[string]string

// Tgtimagekey specify a key/value pair for any image created. (-imagekey is only a synonym if there is no input image).
type Tgtimagekey map[string]string

// Encryption specify a particular type of encryption or, if not specified, the default encryption algorithm.
//
// As of OS X 10.7, the default algorithm is the AES cipher running in CBC mode on 512-byte blocks with a 128-bit key.
type Encryption EncryptionType

// Stdinpass read a null-terminated passphrase from standard input.
// If the standard input is a tty, the passphrase will be read with readpassphrase(3).
// Otherwise, the password is read from stdin.
//
// -stdinpass replaces -passphrase which has been deprecated.
// -passphrase is insecure because its argument appears in the output of ps(1) where it is visible to other users and processes on the system.
type Stdinpass bool

func (s Stdinpass) attachFlag() []string { return boolFlag(bool(s), "stdinpass") }

// Agentpass force the default behavior of prompting for a passphrase.
//
// Useful with -pubkey to create an image protected by both a passphrase and a public key.
type Agentpass bool

// Recover specify a keychain containing the secret corresponding to the certificate specified with -certificate when the image was created.
type Recover string

func (r Recover) attachFlag() []string { return stringFlag(string(r), "recover") }

// Certificate specify a secondary access certificate for an encrypted image.
// cert_file must be DER-encoded certificate data, which can be created by Keychain Access or openssl(1).
type Certificate string

// Pubkey specify a list of public keys, identified by their hexadecimal hashes, to be used to protect the encrypted image being created.
type Pubkey []string

// Cacert specify a certificate authority certificate.
// cert can be either a PEM file or a directory of certificates processed by c_rehash(1).
//
// See also --capath and --cacert in curl(1).
type Cacert string

// Insecurehttp ignore SSL host validation failures.
// Useful for self-signed servers for which the appropriate certificates are unavailable or if access to a server is desired when the server name doesn't match what is in the certificate.
type Insecurehttp bool

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

func (s Shadow) attachFlag() []string { return stringFlag(string(s), "shadow") }

// Verbose be verbose: produce extra progress output and error diagnostics.
//
// This option can help the user decipher why a particular operation failed.
// At a minimum, the probing of any specified images will be detailed.
type Verbose bool

func (v Verbose) attachFlag() []string { return boolFlag(bool(v), "verbose") }

// Quiet close stdout and stderr, leaving only hdiutil's exit status to indicate success or failure.
// No /dev entries or mount points will be printed.
//
// -debug and -verbose disable -quiet.
type Quiet bool

func (q Quiet) attachFlag() []string { return boolFlag(bool(q), "quiet") }

// Debug be very verbose.
//
// This option is good if a large amount of progress information is needed.
// As of Mac OS X 10.6, -debug enables -verbose.
type Debug bool

func (d Debug) attachFlag() []string { return boolFlag(bool(d), "debug") }

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
