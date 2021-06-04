// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

package sys

/*
#cgo CFLAGS: -mmacosx-version-min=10.6

typedef long long long_long;

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

#include <libproc.h>
#include <mach/mach.h>
*/
import "C"

type (
	c_short    C.short
	c_int      C.int
	c_int8     C.int8_t
	c_int16    C.int16_t
	c_int32    C.int32_t
	c_int64    C.int64_t
	c_long     C.long
	c_longLong C.long_long
	c_uint     C.uint
	c_uint8    C.uint8_t
	c_uint16   C.uint16_t
	c_uint32   C.uint32_t
	c_uint64   C.uint64_t
	c_char     C.char
	c_float    C.float
	c_double   C.double
)

type kernReturn C.kern_return_t

const (
	kernSuccess                kernReturn = C.KERN_SUCCESS
	kernInvalidAddress         kernReturn = C.KERN_INVALID_ADDRESS
	kernProtectionFailure      kernReturn = C.KERN_PROTECTION_FAILURE
	kernNoSpace                kernReturn = C.KERN_NO_SPACE
	kernInvalidArgument        kernReturn = C.KERN_INVALID_ARGUMENT
	kernFailure                kernReturn = C.KERN_FAILURE
	kernResourceShortage       kernReturn = C.KERN_RESOURCE_SHORTAGE
	kernNotReceiver            kernReturn = C.KERN_NOT_RECEIVER
	kernNoAccess               kernReturn = C.KERN_NO_ACCESS
	kernMemoryFailure          kernReturn = C.KERN_MEMORY_FAILURE
	KernMemoryError            kernReturn = C.KERN_MEMORY_ERROR
	kernAlreadyInSet           kernReturn = C.KERN_ALREADY_IN_SET
	kernNotInSet               kernReturn = C.KERN_NOT_IN_SET
	kernNameExists             kernReturn = C.KERN_NAME_EXISTS
	kernAborted                kernReturn = C.KERN_ABORTED
	kernInvalidName            kernReturn = C.KERN_INVALID_NAME
	kernInvalidTask            kernReturn = C.KERN_INVALID_TASK
	kernInvalidRight           kernReturn = C.KERN_INVALID_RIGHT
	kernInvalidValue           kernReturn = C.KERN_INVALID_VALUE
	kernUrefsOverflow          kernReturn = C.KERN_UREFS_OVERFLOW
	kernInvalidCapability      kernReturn = C.KERN_INVALID_CAPABILITY
	kernRightExists            kernReturn = C.KERN_RIGHT_EXISTS
	kernInvalidHost            kernReturn = C.KERN_INVALID_HOST
	kernMemoryPresent          kernReturn = C.KERN_MEMORY_PRESENT
	kernMemoryDataMoved        kernReturn = C.KERN_MEMORY_DATA_MOVED
	kernMemoryRestartCopy      kernReturn = C.KERN_MEMORY_RESTART_COPY
	kernInvalidProcessorSet    kernReturn = C.KERN_INVALID_PROCESSOR_SET
	kernPolicyLimit            kernReturn = C.KERN_POLICY_LIMIT
	kernInvalidPolicy          kernReturn = C.KERN_INVALID_POLICY
	kernInvalidObject          kernReturn = C.KERN_INVALID_OBJECT
	kernAlreadyWaiting         kernReturn = C.KERN_ALREADY_WAITING
	kernDefaultSet             kernReturn = C.KERN_DEFAULT_SET
	kernExceptionProtected     kernReturn = C.KERN_EXCEPTION_PROTECTED
	kernInvalidLedger          kernReturn = C.KERN_INVALID_LEDGER
	kernInvalidMemoryControl   kernReturn = C.KERN_INVALID_MEMORY_CONTROL
	kernInvalidSecurity        kernReturn = C.KERN_INVALID_SECURITY
	kernNotDepressed           kernReturn = C.KERN_NOT_DEPRESSED
	kernTerminated             kernReturn = C.KERN_TERMINATED
	kernLockSetDestroyed       kernReturn = C.KERN_LOCK_SET_DESTROYED
	kernLockUnstable           kernReturn = C.KERN_LOCK_UNSTABLE
	kernLockOwned              kernReturn = C.KERN_LOCK_OWNED
	kernLockOwnedSelf          kernReturn = C.KERN_LOCK_OWNED_SELF
	kernSemaphoreDestroyed     kernReturn = C.KERN_SEMAPHORE_DESTROYED
	kernRPCServerTerminated    kernReturn = C.KERN_RPC_SERVER_TERMINATED
	kernRPCTerminateOrphan     kernReturn = C.KERN_RPC_TERMINATE_ORPHAN
	kernRPCContinueOrphan      kernReturn = C.KERN_RPC_CONTINUE_ORPHAN
	kernNotSupported           kernReturn = C.KERN_NOT_SUPPORTED
	kernNodeDown               kernReturn = C.KERN_NODE_DOWN
	kernNotWaiting             kernReturn = C.KERN_NOT_WAITING
	kernOperationTimedOut      kernReturn = C.KERN_OPERATION_TIMED_OUT
	kernCodesignError          kernReturn = C.KERN_CODESIGN_ERROR
	kernPolicyStatic           kernReturn = C.KERN_POLICY_STATIC
	kernInsufficientBufferSize kernReturn = C.KERN_INSUFFICIENT_BUFFER_SIZE
	kernDenied                 kernReturn = C.KERN_DENIED
	kernMissingKC              kernReturn = C.KERN_MISSING_KC
	kernInvalidKC              kernReturn = C.KERN_INVALID_KC
	kernReturnMax              kernReturn = C.KERN_RETURN_MAX
)
