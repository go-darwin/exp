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

// GoString ...
func GoString(p *byte) string {
	return runtime_gostring(p)
}

//go:linkname runtime_gostringn runtime.gostringn
//go:nosplit
func runtime_gostringn(p *byte, l int) string

// GoStringN ...
func GoStringN(p *byte, l int) string {
	return runtime_gostringn(p, l)
}

//go:linkname runtime_gobytes runtime.gobytes
func runtime_gobytes(p *byte, n int) (b []byte)

// GoBytes ...
func GoBytes(p *byte, n int) []byte {
	return runtime_gobytes(p, n)
}

//go:noescape
func crosscall2(fn, a unsafe.Pointer, n int32, ctxt uintptr)

// CrossCall2 ...
func CrossCall2(fn, a unsafe.Pointer, n int32, ctxt uintptr) {
	crosscall2(fn, a, n, ctxt)
}
