// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package logger

import "golang.org/x/net/context"

type Null struct{}

func NewNull() *Null {
	return &Null{}
}

// Verify Null fully implements the Logger interface.
var _ Logger = (*Null)(nil)

func (l *Null) Debug(format string, args ...any)                       {}
func (l *Null) Info(format string, args ...any)                        {}
func (l *Null) Warning(format string, args ...any)                     {}
func (l *Null) Notice(format string, args ...any)                      {}
func (l *Null) Errorf(format string, args ...any)                      {}
func (l *Null) Critical(format string, args ...any)                    {}
func (l *Null) CCriticalf(ctx context.Context, fmt string, arg ...any) {}
func (l *Null) Fatalf(fmt string, arg ...any)                          {}
func (l *Null) CFatalf(ctx context.Context, fmt string, arg ...any)    {}
func (l *Null) Profile(fmts string, arg ...any)                        {}
func (l *Null) CDebugf(ctx context.Context, fmt string, arg ...any)    {}
func (l *Null) CInfof(ctx context.Context, fmt string, arg ...any)     {}
func (l *Null) CNoticef(ctx context.Context, fmt string, arg ...any)   {}
func (l *Null) CWarningf(ctx context.Context, fmt string, arg ...any)  {}
func (l *Null) CErrorf(ctx context.Context, fmt string, arg ...any)    {}
func (l *Null) Error(fmt string, arg ...any)                           {}
func (l *Null) Configure(style string, debug bool, filename string)    {}

func (l *Null) CloneWithAddedDepth(depth int) Logger { return l }

func (l *Null) SetExternalHandler(handler ExternalHandler) {}
func (l *Null) Shutdown()                                  {}
