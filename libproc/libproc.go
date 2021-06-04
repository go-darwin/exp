// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package libproc

import (
	_ "runtime" // used with go:linkname
	"unsafe"

	"golang.org/x/sys/unix"

	"go-darwin.dev/exp/sys"
)

// ProcPidpath given a pid, returns the full executable name including directory
// paths and the longer-than-16-chars executable name.
//
// C implementation:
//
//  int
//  proc_pidpath(int pid, void * buffer, uint32_t buffersize)
//  {
//  	int retval, len;
//
//  	if (buffersize < PROC_PIDPATHINFO_SIZE) {
//  		errno = ENOMEM;
//  		return(0);
//  	}
//  	if (buffersize >  PROC_PIDPATHINFO_MAXSIZE) {
//  		errno = EOVERFLOW;
//  		return(0);
//  	}
//
//  	retval = __proc_info(PROC_INFO_CALL_PIDINFO, pid, PROC_PIDPATHINFO,  (uint64_t)0,  buffer, buffersize);
//  	if (retval != -1) {
//  		len = (int)strlen(buffer);
//  		return(len);
//  	}
//  	return (0);
//  }
func ProcPidpath(pid int) (path string, err error) {
	buffer := make([]byte, ProcPidpathinfoSize)
	_, _, kret := procPidpath(pid, unsafe.Pointer(&buffer[0]), uint32(ProcPidpathinfoMaxsize))
	if kret != 0 {
		return "", sys.Errno(kret)
	}

	path = sys.GoString(&buffer[0])
	return
}

//go:nosplit
func procPidpath(pid int, buffer unsafe.Pointer, buffersize uint32) (r1, r2 uintptr, err unix.Errno) {
	return sys.Syscall(libc_proc_pidpath_trampoline_addr, uintptr(pid), uintptr(buffer), uintptr(buffersize))
}

var libc_proc_pidpath_trampoline_addr uintptr

//go:cgo_import_dynamic libc_proc_pidpath proc_pidpath "/usr/lib/libSystem.B.dylib"
