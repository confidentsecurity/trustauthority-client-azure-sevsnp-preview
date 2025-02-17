//go:build test

/*
 *   Copyright (c) 2022 Intel Corporation
 *   All rights reserved.
 *   SPDX-License-Identifier: BSD-3-Clause
 */
package tdx

import (
	"github.com/confidentsecurity/trustauthority-client-azure-sevsnp-preview/go-connector"
)

type mockAdapter struct {
	uData       []byte
	EvLogParser EventLogParser
}

func NewEvidenceAdapter(udata []byte, evLogParser EventLogParser) (connector.EvidenceAdapter, error) {
	return &mockAdapter{
		uData:       udata,
		EvLogParser: evLogParser,
	}, nil
}

func (adapter *mockAdapter) CollectEvidence(nonce []byte) (*connector.Evidence, error) {

	return &connector.Evidence{
		Type:     1,
		Quote:    nil,
		UserData: nil,
		EventLog: nil,
	}, nil
}
