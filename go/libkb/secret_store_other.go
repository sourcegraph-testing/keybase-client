// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

//go:build !darwin && !android && !linux
// +build !darwin,!android,!linux

package libkb

func NewSecretStoreAll(m MetaContext) SecretStoreAll {
	g := m.G()
	s := NewSecretStoreFile(g.Env.GetDataDir())
	s.notifyCreate = func(name NormalizedUsername) { notifySecretStoreCreate(m, name) }
	return s
}
