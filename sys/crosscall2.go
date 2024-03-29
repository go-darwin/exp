// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package sys

import (
	_ "runtime" // for go:linkname
	"unsafe"
)

//go:linkname runtime_cgocall runtime.cgocall
//go:nosplit
func runtime_cgocall(fn unsafe.Pointer, arg uintptr) int32

// CgoCall ...
func CgoCall(fn unsafe.Pointer, arg uintptr) int32 {
	return runtime_cgocall(fn, arg)
}

//go:linkname runtime_gostring runtime.gostring
//go:nosplit
func runtime_gostring(p *byte) string

// C string to Go string.
func GoString(p *byte) string {
	return runtime_gostring(p)
}

//go:linkname runtime_gostringn runtime.gostringn
//go:nosplit
func runtime_gostringn(p *byte, l int) string

// C string, length to Go string.
func GoStringN(p *byte, l int) string {
	return runtime_gostringn(p, l)
}

//go:linkname runtime_gobytes runtime.gobytes
func runtime_gobytes(p *byte, n int) (b []byte)

// C pointer, length to Go []byte,
func GoBytes(p *byte, n int) []byte {
	return runtime_gobytes(p, n)
}

//go:noescape
func crosscall2(fn, a unsafe.Pointer, n int32, ctxt uintptr)

// CrossCall2 called by C code generated by cmd/cgo.
// Saves C callee-saved registers and calls cgocallback with three arguments.
// fn is the PC of a func(a unsafe.Pointer) function.
func CrossCall2(fn, a unsafe.Pointer, n int32, ctxt uintptr) {
	crosscall2(fn, a, n, ctxt)
}
