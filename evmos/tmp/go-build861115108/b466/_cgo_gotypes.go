//go:cgo_ldflag "-g"
//go:cgo_ldflag "-O2"
// Code generated by cmd/cgo; DO NOT EDIT.

package cgo

import "unsafe"

func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

