// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

//go:build !debug
// +build !debug

package dokan

const isDebug = false //nolint

func debug(...any)          {} // nolint
func debugf(string, ...any) {} // nolint
