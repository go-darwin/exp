// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

#include "textflag.h"
#include "abi_amd64.h"

// crosscall2 called by C code generated by cmd/cgo.
// Saves C callee-saved registers and calls cgocallback with three arguments.
// fn is the PC of a func(a unsafe.Pointer) function.
//
// func crosscall2(fn, a unsafe.Pointer, n int32, ctxt uintptr)
TEXT ·crosscall2(SB), NOSPLIT, $0-0
	PUSH_REGS_HOST_TO_ABI0()

	// Make room for arguments to cgocallback.
	ADJSP $0x18
	MOVQ  DI, 0x0(SP) // fn
	MOVQ  SI, 0x8(SP) // arg

	// Skip n in DX.
	MOVQ CX, 0x10(SP) // ctxt

	CALL runtime·cgocallback(SB)

	ADJSP $-0x18
	POP_REGS_HOST_TO_ABI0()
	RET
