// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package ldap

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240805001/admin"
)

func TestVerify_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockLDAPConfigurationVerifier(ctrl)

	expected := &atlasv2.LDAPVerifyConnectivityJobRequest{
		RequestId: pointer.Get("5f9b0b4e1f6f7d4e6f7d4e6f"),
		GroupId:   pointer.Get("5f9b0b4e1f6f7d4e6f7d4e6f"),
		Status:    pointer.Get("SUCCESS"),
	}

	opts := &VerifyOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		VerifyLDAPConfiguration(opts.ProjectID, opts.newLDAP()).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, verifyTemplate, *expected)
}

func TestVerifyBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		VerifyBuilder(),
		1,
		[]string{flag.ProjectID, flag.Hostname, flag.Port, flag.BindPassword, flag.BindUsername, flag.CaCertificate, flag.AuthzQueryTemplate},
	)
}
