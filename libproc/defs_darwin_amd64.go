// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

package libproc

/*
#cgo CFLAGS: -mmacosx-version-min=10.6

#include <libproc.h>
#include <mach/mach.h>
*/
import "C"

const (
	ProcPidpathinfoMaxsize = C.PROC_PIDPATHINFO_MAXSIZE
	ProcPidpathinfoSize    = C.PROC_PIDPATHINFO_SIZE
)

type PID C.pid_t
