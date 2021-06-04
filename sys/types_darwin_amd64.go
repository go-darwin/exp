// SPDX-FileCopyrightText: Copyright 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package sys

import "unsafe"

// list of common C types.
type (
	C_short    = c_short
	C_int      = c_int
	C_int8     = c_int8
	C_int16    = c_int16
	C_int32    = c_int32
	C_int64    = c_int64
	C_long     = c_long
	C_longLong = c_longLong
	C_uint     = c_uint
	C_uint8    = c_uint8
	C_uint16   = c_uint16
	C_uint32   = c_uint32
	C_uint64   = c_uint64
	C_char     = c_char
	C_float    = c_float
	C_double   = c_double
)

// list of common C types size.
const (
	Sizeof_C_short    = unsafe.Sizeof(new(c_short))
	Sizeof_C_int      = unsafe.Sizeof(new(c_int))
	Sizeof_C_int8     = unsafe.Sizeof(new(c_int8))
	Sizeof_C_int16    = unsafe.Sizeof(new(c_int16))
	Sizeof_C_int32    = unsafe.Sizeof(new(c_int32))
	Sizeof_C_int64    = unsafe.Sizeof(new(c_int64))
	Sizeof_C_long     = unsafe.Sizeof(new(c_long))
	Sizeof_C_longLong = unsafe.Sizeof(new(c_longLong))
	Sizeof_C_uint     = unsafe.Sizeof(new(c_uint))
	Sizeof_C_uint8    = unsafe.Sizeof(new(c_uint8))
	Sizeof_C_uint16   = unsafe.Sizeof(new(c_uint16))
	Sizeof_C_uint32   = unsafe.Sizeof(new(c_uint32))
	Sizeof_C_uint64   = unsafe.Sizeof(new(c_uint64))
	Sizeof_C_char     = unsafe.Sizeof(new(c_char))
	Sizeof_C_float    = unsafe.Sizeof(new(c_float))
	Sizeof_C_double   = unsafe.Sizeof(new(c_double))
)

// KernReturn represents a kern_return_t.
type KernReturn = kernReturn

// list of KernReturn errors.
const (
	KernSuccess KernReturn = kernSuccess

	// KernInvalidAddress specified address is not currently valid.
	KernInvalidAddress KernReturn = kernInvalidAddress

	// KernProtectionFailure specified memory is valid, but does not permit therequired forms of access.
	KernProtectionFailure KernReturn = kernProtectionFailure

	// KernNoSpace is the address range specified is already in use, or
	// no address range of the size specified could be
	// found.
	KernNoSpace KernReturn = kernNoSpace

	// KernInvalidArgument is the function requested was not applicable to this
	// type of argument, or an argument is invalid.
	KernInvalidArgument KernReturn = kernInvalidArgument

	// KernFailure is the function could not be performed. A catch-all.
	KernFailure KernReturn = kernFailure

	// KernResourceShortage a system resource could not be allocated to fulfill
	// this request.
	// This failure may not be permanent.
	KernResourceShortage KernReturn = kernResourceShortage

	// KernNoAccess bogus access restriction.
	KernNoAccess KernReturn = kernNoAccess

	// KernNotReceiver is the task in question does not hold receive rights for the port argument.
	KernNotReceiver KernReturn = kernNotReceiver

	// KernMemoryFailure during a page fault, the target address refers to a memory object that has been destroyed.
	// This failure is permanent.
	KernMemoryFailure KernReturn = kernMemoryFailure

	// kernMemoryError during a page fault, the memory object indicated
	// that the data could not be returned.
	// This failure may be temporary; future attempts to access this
	// same data may succeed, as defined by the memory
	// object.
	kernMemoryError KernReturn = KernMemoryError

	// KernAlreadyInSet is the receive right is already a member of the portset.
	KernAlreadyInSet KernReturn = kernAlreadyInSet

	// KernNotInSet is the receive right is not a member of a port set.
	KernNotInSet KernReturn = kernNotInSet

	// KernNameExists is the name already denotes a right in the task.
	KernNameExists KernReturn = kernNameExists

	// KernAborted is the operation was aborted.
	// Ipc code will catch this and reflect it as a message error.
	KernAborted KernReturn = kernAborted

	// KernInvalidNames is the name doesn't denote a right in the task.
	KernInvalidName KernReturn = kernInvalidName

	// KernInvalidTask target task isn't an active task.
	KernInvalidTask KernReturn = kernInvalidTask

	// kernInvalidRight is the name denotes a right, but not an appropriate right.
	KernInvalidRight KernReturn = kernInvalidRight

	// KernInvalidValue a blatant range error.
	KernInvalidValue KernReturn = kernInvalidValue

	// KernUrefsOverflow operation would overflow limit on user-references.
	KernUrefsOverflow KernReturn = kernUrefsOverflow

	// KernInvalidCapability is the supplied (port) capability is improper.
	KernInvalidCapability KernReturn = kernInvalidCapability

	// KernRightExists is the task already has send or receive rights for the port under another name.
	KernRightExists KernReturn = kernRightExists

	// KernInvalidHost target host isn't actually a host.
	KernInvalidHost KernReturn = kernInvalidHost

	// KernMemoryPresent an attempt was made to supply "precious" data
	// for memory that is already present in a
	// memory object.
	KernMemoryPresent KernReturn = kernMemoryPresent

	// KernMemoryDataMoved a page was requested of a memory manager via
	// memory_object_data_request for an object using
	// a MEMORY_OBJECT_COPY_CALL strategy, with the
	// VM_PROT_WANTS_COPY flag being used to specify
	// that the page desired is for a copy of the
	// object, and the memory manager has detected
	// the page was pushed into a copy of the object
	// while the kernel was walking the shadow chain
	// from the copy to the object. This error code
	// is delivered via memory_object_data_error
	// and is handled by the kernel (it forces the
	// kernel to restart the fault). It will not be
	// seen by users.
	KernMemoryDataMoved KernReturn = kernMemoryDataMoved

	// KernMemoryRestartCopy a strategic copy was attempted of an object
	// upon which a quicker copy is now possible.
	// The caller should retry the copy using
	// vm_object_copy_quickly.
	//
	// This error code is seen only by the kernel.
	KernMemoryRestartCopy KernReturn = kernMemoryRestartCopy

	// KernInvalidProcessorSet an argument applied to assert processor set privilege
	// was not a processor set control port.
	KernInvalidProcessorSet KernReturn = kernInvalidProcessorSet

	// KernPolicyLimit is the specified scheduling attributes exceed the thread's
	// limits.
	KernPolicyLimit KernReturn = kernPolicyLimit

	// KernInvalidPolicy is the specified scheduling policy is not currently
	// enabled for the processor set.
	KernInvalidPolicy KernReturn = kernInvalidPolicy

	// KernInvalidObject is the external memory manager failed to initialize the
	// memory object.
	KernInvalidObject KernReturn = kernInvalidObject

	// KernAlreadyWaiting a thread is attempting to wait for an event for which
	// there is already a waiting thread.
	KernAlreadyWaiting KernReturn = kernAlreadyWaiting

	// KernDefaultSet an attempt was made to destroy the default processor
	// set.
	KernDefaultSet KernReturn = kernDefaultSet

	// KernExceptionProtected an attempt was made to fetch an exception port that is
	// protected, or to abort a thread while processing a
	// protected exception.
	KernExceptionProtected KernReturn = kernExceptionProtected

	// KernInvalidLedger a ledger was required but not supplied.
	KernInvalidLedger KernReturn = kernInvalidLedger

	// KernInvalidMemoryControl is the port was not a memory cache control port.
	KernInvalidMemoryControl KernReturn = kernInvalidMemoryControl

	// KernInvalidSecurity an argument supplied to assert security privilege
	// was not a host security port.
	KernInvalidSecurity KernReturn = kernInvalidSecurity

	// KernNotDepressed thread_depress_abort was called on a thread which
	// was not currently depressed.
	KernNotDepressed KernReturn = kernNotDepressed

	// KernTerminated object has been terminated and is no longer available.
	KernTerminated KernReturn = kernTerminated

	// KernLockSetDestroyed lock set has been destroyed and is no longer available.
	KernLockSetDestroyed KernReturn = kernLockSetDestroyed

	// KernLockUnstable is the thread holding the lock terminated before releasing
	// the lock.
	KernLockUnstable KernReturn = kernLockUnstable

	// KernLockOwned is the lock is already owned by another thread.
	KernLockOwned KernReturn = kernLockUnstable

	// KernLockOwnedSelf is the lock is already owned by the calling thread.
	KernLockOwnedSelf KernReturn = kernLockOwnedSelf

	// KernSemaphoreDestroyed semaphore has been destroyed and is no longer available.
	KernSemaphoreDestroyed KernReturn = kernSemaphoreDestroyed

	// KernRPCServerTerminated return from RPC indicating the target server was
	// terminated before it successfully replied.
	KernRPCServerTerminated KernReturn = kernRPCServerTerminated

	// KernRPCTerminateOrphan terminate an orphaned activation.
	KernRPCTerminateOrphan KernReturn = kernRPCTerminateOrphan

	// KernRPCContinueOrphan allow an orphaned activation to continue executing.
	KernRPCContinueOrphan KernReturn = kernRPCContinueOrphan

	// KernNotSupported empty thread activation (No thread linked to it).
	KernNotSupported KernReturn = kernRPCContinueOrphan

	// KernNodeDown remote node down or inaccessible.
	KernNodeDown KernReturn = kernNodeDown

	// KernNotWaiting a signalled thread was not actually waiting.
	KernNotWaiting KernReturn = kernNotWaiting

	// KernOperationTimedOut some thread-oriented operation (semaphore_wait) timed out.
	KernOperationTimedOut KernReturn = kernOperationTimedOut

	// KernCodesignError during a page fault, indicates that the page was rejected
	// as a result of a signature check.
	KernCodesignError KernReturn = kernCodesignError

	// KernPolicyStatic is the requested property cannot be changed at this time.
	KernPolicyStatic KernReturn = kernPolicyStatic

	// KernInsufficientBufferSize is the provided buffer is of insufficient size for the requested data.
	KernInsufficientBufferSize KernReturn = kernInsufficientBufferSize

	// KernDenied denied by security policy.
	KernDenied KernReturn = kernDenied

	// KernMissingKC is the KC on which the function is operating is missing.
	KernMissingKC KernReturn = kernMissingKC

	// KernInvalidKC is the KC on which the function is operating is invalid.
	KernInvalidKC KernReturn = kernInvalidKC

	// KernReturnMax maximum return value allowable.
	KernReturnMax KernReturn = kernReturnMax
)

// KernReturn Error table.
var errors = [...]string{
	0x1:   "specified address is not currently valid",
	0x2:   "specified memory is valid, but does not permit therequired forms of access",
	0x3:   "address range specified is already in use, or no address range of the size specified could be found",
	0x4:   "function requested was not applicable to this type of argument, or an argument is invalid",
	0x5:   "function could not be performed. A catch-all",
	0x6:   "system resource could not be allocated to fulfill this request",
	0x7:   "bogus access restriction",
	0x8:   "task in question does not hold receive rights for the port argument",
	0x9:   "during a page fault, the target address refers to a memory object that has been destroyed",
	0xa:   "during a page fault, the memory object indicated that the data could not be returned",
	0xb:   "receive right is already a member of the portset",
	0xc:   "receive right is not a member of a port set",
	0xd:   "name already denotes a right in the task",
	0xe:   "operation was aborted",
	0xf:   "name doesn't denote a right in the task",
	0x10:  "target task isn't an active task",
	0x11:  "name denotes a right, but not an appropriate right",
	0x12:  "blatant range error",
	0x13:  "operation would overflow limit on user-references",
	0x14:  "supplied (port) capability is improper",
	0x15:  "task already has send or receive rights for the port under another name",
	0x16:  "target host isn't actually a host",
	0x18:  "attempt was made to supply precious data for memory that is already present in a memory object",
	0x19:  "page was requested of a memory manager via memory_object_data_request for an object using a MEMORY_OBJECT_COPY_CALL strategy",
	0x1a:  "strategic copy was attempted of an object upon which a quicker copy is now possible",
	0x1b:  "argument applied to assert processor set privilege was not a processor set control port",
	0x1c:  "specified scheduling attributes exceed the thread's limits",
	0x1d:  "specified scheduling policy is not currently enabled for the processor set",
	0x1e:  "external memory manager failed to initialize the memory object",
	0x1f:  "thread is attempting to wait for an event for which there is already a waiting thread",
	0x20:  "attempt was made to destroy the default processor set",
	0x21:  "attempt was made to fetch an exception port that is protected, or to abort a thread while processing a protected exception",
	0x22:  "ledger was required but not supplied",
	0x23:  "port was not a memory cache control port",
	0x24:  "argument supplied to assert security privilege was not a host security port",
	0x25:  "thread_depress_abort was called on a thread which was not currently depressed",
	0x26:  "object has been terminated and is no longer available",
	0x27:  "lock set has been destroyed and is no longer available",
	0x28:  "thread holding the lock terminated before releasing the lock",
	0x29:  "lock is already owned by another thread",
	0x2a:  "lock is already owned by the calling thread",
	0x2b:  "semaphore has been destroyed and is no longer available",
	0x2c:  "return from RPC indicating the target server was terminated before it successfully replied",
	0x2d:  "terminate an orphaned activation",
	0x2e:  "allow an orphaned activation to continue executing",
	0x2f:  "empty thread activation (No thread linked to it)",
	0x30:  "remote node down or inaccessible",
	0x31:  "signalled thread was not actually waiting",
	0x32:  "some thread-oriented operation (semaphore_wait) timed out",
	0x33:  "during a page fault, indicates that the page was rejected as a result of a signature check",
	0x34:  "requested property cannot be changed at this time",
	0x35:  "provided buffer is of insufficient size for the requested data",
	0x36:  "KC on which the function is operating is missing",
	0x37:  "KC on which the function is operating is invalid",
	0x100: "maximum return value allowable",
}
