[![Build Status](https://travis-ci.com/48d90782/hidden-slice.svg?branch=master)](https://travis-ci.com/48d90782/hidden-slice)

hidden-slice is a package that can help with allocation of big slices with pointers which hidden from golang GC.
All test and benchmarks are collected in slice_test.go package.  
Main purpose is to define and manage slice which contains a pointers to elements w/o affecting GC time.
As for example, average GC time when create memory mapped slice:  
`GC time average 152.577µs`  
And std slice:  
`GC time average 10.2ms`  
Documentation:  
It uses mmap() syscall to map file into the memory. The program works only in POSIX compatible systems (not Windows).
Links to full documentation:  
http://man7.org/linux/man-pages/man2/mmap.2.html  
https://www.gnu.org/software/libc/manual/html_node/Memory_002dmapped-I_002fO.html  
Description:  
First of all we need to know size of one element to allocate proper amount of memory for all elements:  
-  `size := unsafe.Sizeof(&UserDefined{})`  
We use syscall with 6 parameters: `syscall.Syscall6`  
1. Specify type of syscall  
- `syscall.SYS_MMAP`  
2. Address: according to the documentation:  
If addr is NULL, then the kernel chooses the (page-aligned) address at which to create the mapping; this is the most
portable method of creating a new mapping. If addr is not NULL, then the kernel takes it as a hint about where to
place the mapping; on Linux, the kernel will pick a nearby page boundary (but always above or equal to the value
specified by `/proc/sys/vm/mmap_min_addr`) and attempt to create the mapping there.  If another mapping already exists
there, the kernel picks a new address that may or may not depend on the hint. The address of the new mapping is
returned as the result of the call.  
3. The required size of memory simply calculated by multiplying number of elements by size of one element  
    `uintptr(len)*size`  
4. Prot argument:  
The prot argument describes the desired memory protection of the mapping (and must not conflict with the open mode of the file).
It is either PROT_NONE or the bitwise OR of one or more of the following flags:  
- `PROT_EXEC  Pages may be executed`  
- `PROT_READ  Pages may be read`  
- `PROT_WRITE Pages may be written`  
- `PROT_NONE  Pages may not be accessed`  
We defining, that we can read and write by passing:  
- `syscall.PROT_READ|syscall.PROT_WRITE`  
5. The flags argument determines whether updates to the mapping are visible to other processes mapping the same region,
and whether updates are carried through to the underlying file. You can read about this parameter by link above.  
- `syscall.MAP_ANON|syscall.MAP_PRIVATE`  
6. File descriptor used in case if we map file into memory (the file should have enough free space). But in our case we
 have no any file an map directly to memory:  
- `uintptr(fd)`  
7. And the last parameter, offset. You can check OS parameter by writing simple program in C:
```C
size_t page_size = (size_t) sysconf (_SC_PAGESIZE);
printf("%lu", page_size);
```  
But we provide `0` because we want to map entire file into memory:  
8. Slice converting. We simply provide to slice basic structure (every slice in golang consists of "SliceHeader") our
address to memory, and `len` and `cap` arguments  
```go
	slice := SliceHeader{
		Data: data,
		Len:  len,
		Cap:  len,
	}
```
