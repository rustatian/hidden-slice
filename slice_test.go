package hidden_slice

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

var sliceLen int = 10000000

func Benchmark_FunSlice(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		a := makeIntSlice(sliceLen)
		_ = a
	}
}

func Benchmark_NoNFunSlice(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		a := make([]*int, sliceLen)
		_ = a
	}
}

func Test_CorrectAlloc(t *testing.T) {
	a := IntSlice(sliceLen)
	b := StringSlice(sliceLen)
	if len(a) == 0 || len(b) == 0 {
		t.Error("couldn't allocate memory for the slice")
	}
}

func Test_GCAverageTimeSyscall(t *testing.T) {
	_ = IntSlice(sliceLen, )

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC time %s\n", time.Since(start))
	}
}

func Test_GCAverageTimeDirect(t *testing.T) {
	slice := make([]*int, sliceLen)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()

		fmt.Printf("GC time %s\n", time.Since(start))
	}

	runtime.KeepAlive(slice)
}

func Test_UserDefinedStructure(t *testing.T) {
	slice := make([]*UserDefined, sliceLen)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()

		fmt.Printf("GC time %s\n", time.Since(start))
	}

	runtime.KeepAlive(slice)
}

func Test_UserDefinedStructure_Syscall(t *testing.T) {
	_ = makeCustomSlice(sliceLen)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()

		fmt.Printf("GC time %s\n", time.Since(start))
	}
}
