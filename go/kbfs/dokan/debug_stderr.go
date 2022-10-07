// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

//go:build debug
// +build debug

package dokan

import (
	"log"
)

const isDebug = true

func debug(args ...any) {
	log.Println(args...)
}

func debugf(s string, args ...any) {
	log.Printf(s, args...)
}
