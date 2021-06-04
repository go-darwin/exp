// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

#include "textflag.h"

// func proc_listallpids() (r1, r2 uintptr, err unix.Errno)
GLOBL ·libc_proc_listallpids_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_listallpids_trampoline_addr(SB)/8, $libc_proc_listallpids_trampoline<>(SB)

TEXT libc_proc_listallpids_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_listallpids(SB)

// func procPidpath(pid PID, buffer unsafe.Pointer, buffersize uint32) (r1, r2 uintptr, err unix.Errno)
GLOBL ·libc_proc_pidpath_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_pidpath_trampoline_addr(SB)/8, $libc_proc_pidpath_trampoline<>(SB)

TEXT libc_proc_pidpath_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_pidpath(SB)
