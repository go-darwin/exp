// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package libproc

import (
	_ "runtime" // for go:linkname
	"unsafe"

	"golang.org/x/sys/unix"

	"go-darwin.dev/exp/sys"
)

// ProcListallpids retursn the list all pids.
//
// C implementation:
//
//  int
//  proc_listallpids(void * buffer, int buffersize)
//  {
//  	int numpids;
//  	numpids = proc_listpids(PROC_ALL_PIDS, (uint32_t)0, buffer, buffersize);
//
//  	if (numpids == -1) {
//  		return -1;
//  	} else {
//  		return numpids / sizeof(int);
//  	}
//  }
func ProcListallpids() (numPids int, err error) {
	r1, _, kret := proc_listallpids()
	if sys.KernReturn(kret) != sys.KernSuccess {
		return 0, sys.KernErrno(sys.KernReturn(kret))
	}

	numPids = int(r1 / sys.Sizeof_C_int)
	return
}

//go:nosplit
func proc_listallpids() (r1, r2 uintptr, err unix.Errno) {
	return sys.Syscall(libc_proc_listallpids_trampoline_addr, uintptr(0), uintptr(0), 0)
}

var libc_proc_listallpids_trampoline_addr uintptr

//go:cgo_import_dynamic libc_proc_listallpids proc_listallpids "/usr/lib/libSystem.B.dylib"

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
func ProcPidpath(pid PID) (path string, err error) {
	buffer := make([]byte, ProcPidpathinfoSize)
	_, _, kret := procPidpath(pid, unsafe.Pointer(&buffer[0]), uint32(ProcPidpathinfoMaxsize))
	if sys.KernReturn(kret) != sys.KernSuccess {
		return "", sys.KernErrno(sys.KernReturn(kret))
	}

	path = sys.GoString(&buffer[0])
	return
}

//go:nosplit
func procPidpath(pid PID, buffer unsafe.Pointer, buffersize uint32) (r1, r2 uintptr, err unix.Errno) {
	return sys.Syscall(libc_proc_pidpath_trampoline_addr, uintptr(pid), uintptr(buffer), uintptr(buffersize))
}

var libc_proc_pidpath_trampoline_addr uintptr

//go:cgo_import_dynamic libc_proc_pidpath proc_pidpath "/usr/lib/libSystem.B.dylib"
