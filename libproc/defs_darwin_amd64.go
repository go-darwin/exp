// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

package libproc

/*
#cgo CFLAGS: -mmacosx-version-min=10.6

typedef long long long_long;

#include <libproc.h>
#include <mach/mach.h>
#include <stdlib.h>
*/
import "C"

type (
	C_short    C.short
	C_int      C.int
	C_long     C.long
	C_longLong C.long_long
	C_uint     C.uint
	C_char     C.char
	C_float    C.float
	C_double   C.double
)

type (
	KernReturn C.kern_return_t
)

// list of KernReturn errors.
const (
	KernSuccess        KernReturn = C.KERN_SUCCESS
	KernAborted        KernReturn = C.KERN_ABORTED
	KernAlreadyInSET   KernReturn = C.KERN_ALREADY_IN_SET
	KernAlreadyWaiting KernReturn = C.KERN_ALREADY_WAITING
	KernCodeSelector   KernReturn = C.KERN_CODE_SELECTOR
	KernCodesignError  KernReturn = C.KERN_CODESIGN_ERROR
)

const (
	ProcPidpathinfoMaxsize = C.PROC_PIDPATHINFO_MAXSIZE
	ProcPidpathinfoSize    = C.PROC_PIDPATHINFO_SIZE
)

type PID C.pid_t
