/*
 *   Copyright (c) 2022 Intel Corporation
 *   All rights reserved.
 *   SPDX-License-Identifier: BSD-3-Clause
 */

package constants

const (
	RSAKeyBitLength     = 3072
	LinuxFilePathSize   = 4096
	CLIShortDescription = "Intel® Trust Authority CLI for SEVSNP"
)

// Command Names
const (
	ReportCmd  = "report"
	TokenCmd   = "token"
	RootCmd    = "trustauthority-sevsnp-cli"
	VersionCmd = "version"
	VerifyCmd  = "verify"
)

// Options Names
const (
	PrivateKeyPathOption  = "key-path"
	PublicKeyPathOption   = "pub-path"
	PrivateKeyOption      = "key"
	PolicyIdsOption       = "policy-ids"
	InputOption           = "in"
	UserDataOption        = "user-data"
	NonceOption           = "nonce"
	ConfigOption          = "config"
	RequestIdOption       = "request-id"
	TokenAlgOption        = "token-signing-alg"
	PolicyMustMatchOption = "policy-must-match"
	NoEventLogOption      = "no-eventlog"
	TokenOption           = "token"
)
