package hidden_slice

import (
	"syscall"
	"unsafe"
)

// IntSlice
func IntSlice(len int) []*int {
	return makeIntSlice(len)
}

// StringSlice
func StringSlice(len int) []*string {
	return makeStringSlice(len)
}

// SliceHeader is slice under the hood representation
// we could also use SliceHeader from reflect package
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func makeIntSlice(len int) []*int {
	fd := -1

	var s *int
	size := unsafe.Sizeof(s)

	data, _, errno := syscall.Syscall6(
		syscall.SYS_MMAP,
		0,
		uintptr(len)*size,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
		uintptr(fd),
		0,
	)

	if errno != 0 {
		panic(errno)
	}

	slice := SliceHeader{
		Data: data,
		Len:  len,
		Cap:  len,
	}

	return *(*[]*int)(unsafe.Pointer(&slice))
}

func makeStringSlice(len int) []*string {
	fd := -1

	var s *string
	size := unsafe.Sizeof(s)

	data, _, errno := syscall.Syscall6(
		syscall.SYS_MMAP,
		0,
		uintptr(len)*size,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
		uintptr(fd),
		0,
	)

	if errno != 0 {
		panic(errno)
	}

	slice := SliceHeader{
		Data: data,
		Len:  len,
		Cap:  len,
	}

	return *(*[]*string)(unsafe.Pointer(&slice))
}

type UserDefined struct {
	a int
	b string
	c bool
}

func makeCustomSlice(len int) []*UserDefined {
	fd := -1

	size := unsafe.Sizeof(&UserDefined{})
	data, _, errno := syscall.Syscall6(
		syscall.SYS_MMAP,
		0,
		uintptr(len)*size,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
		uintptr(fd),
		0,
	)

	if errno != 0 {
		panic(errno)
	}

	slice := SliceHeader{
		Data: data,
		Len:  len,
		Cap:  len,
	}

	return *(*[]*UserDefined)(unsafe.Pointer(&slice))
}
