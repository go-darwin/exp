// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -srcdir . -godefs -import_runtime_cgo=false -import_syscall=false /Users/zchee/go/src/go-darwin.dev/exp/pkg/libproc/defs_darwin_amd64.go

package libproc

type (
	C_short		int16
	C_int		int32
	C_long		int64
	C_longLong	int64
	C_uint		uint32
	C_char		int8
	C_float		float32
	C_double	float64
)

type (
	KernReturn int32
)

const (
	KernSuccess		KernReturn	= 0x0
	KernAborted		KernReturn	= 0xe
	KernAlreadyInSET	KernReturn	= 0xb
	KernAlreadyWaiting	KernReturn	= 0x1e
	KernCodeSelector	KernReturn	= 0x8
	KernCodesignError	KernReturn	= 0x32
)

const (
	ProcPidpathinfoMaxsize	= 0x1000
	ProcPidpathinfoSize	= 0x400
)

type PID int32
