// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

//go:build !production
// +build !production

package externals

import (
	"regexp"
	"strings"

	libkb "github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	jsonw "github.com/keybase/go-jsonw"
)

// =============================================================================
// Rooter
//

type RooterChecker struct {
	proof libkb.RemoteProofChainLink
}

var _ libkb.ProofChecker = (*RooterChecker)(nil)

func NewRooterChecker(p libkb.RemoteProofChainLink) (*RooterChecker, libkb.ProofError) {
	return &RooterChecker{p}, nil
}

func (rc *RooterChecker) GetTorError() libkb.ProofError { return nil }

func (rc *RooterChecker) CheckStatus(mctx libkb.MetaContext, h libkb.SigHint, _ libkb.ProofCheckerMode,
	pvlU keybase1.MerkleStoreEntry) (*libkb.SigHint, libkb.ProofError) {
	// TODO CORE-8951 see if we can populate verifiedHint with anything useful.
	return nil, CheckProofPvl(mctx, keybase1.ProofType_ROOTER, rc.proof, h, pvlU)
}

//
// =============================================================================

type RooterServiceType struct{ libkb.BaseServiceType }

func (t *RooterServiceType) Key() string { return t.GetTypeName() }

var rooterUsernameRegexp = regexp.MustCompile(`^(?i:[a-z0-9_]{1,20})$`)

func (t *RooterServiceType) NormalizeUsername(s string) (string, error) {
	if !rooterUsernameRegexp.MatchString(s) {
		return "", libkb.NewBadUsernameError(s)
	}
	return strings.ToLower(s), nil
}

func (t *RooterServiceType) NormalizeRemoteName(_ libkb.MetaContext, s string) (string, error) {
	// Allow a leading '@'.
	s = strings.TrimPrefix(s, "@")
	return t.NormalizeUsername(s)
}

func (t *RooterServiceType) GetPrompt() string {
	return "Your username on Rooter"
}

func (t *RooterServiceType) CanMakeNewProofs(mctx libkb.MetaContext) bool {
	return mctx.G().GetRunMode() != libkb.ProductionRunMode
}

func (t *RooterServiceType) ToServiceJSON(un string) *jsonw.Wrapper {
	return t.BaseToServiceJSON(t, un)
}

func (t *RooterServiceType) PostInstructions(un string) *libkb.Markup {
	return libkb.FmtMarkup(`Please toot the following, and don't delete it:`)
}

func (t *RooterServiceType) DisplayName() string   { return "Rooter" }
func (t *RooterServiceType) GetTypeName() string   { return "rooter" }
func (t *RooterServiceType) PickerSubtext() string { return "" }
func (t *RooterServiceType) RecheckProofPosting(tryNumber int, status keybase1.ProofStatus, _ string) (warning *libkb.Markup, err error) {
	return t.BaseRecheckProofPosting(tryNumber, status)
}
func (t *RooterServiceType) GetProofType() string { return "test.web_service_binding.rooter" }

func (t *RooterServiceType) CheckProofText(text string, id keybase1.SigID, sig string) (err error) {
	return t.BaseCheckProofTextShort(text, id, true)
}

func (t *RooterServiceType) MakeProofChecker(l libkb.RemoteProofChainLink) libkb.ProofChecker {
	return &RooterChecker{l}
}

func (t *RooterServiceType) IsDevelOnly() bool { return true }
