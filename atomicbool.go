package multitimer

import (
	"sync/atomic"
)

// Replace with atomic.Bool in Go 1.19+
type atomicBool uintptr

func bool2int(v bool) uintptr {
	if v {
		return 1
	}
	return 0
}

func (b *atomicBool) Store(value bool) {
	atomic.StoreUintptr((*uintptr)(b), bool2int(value))
}

func (b *atomicBool) Load() bool {
	return atomic.LoadUintptr((*uintptr)(b)) != 0
}

func (b *atomicBool) CompareAndSwap(old, new bool) (swapped bool) {
	return atomic.CompareAndSwapUintptr((*uintptr)(b), bool2int(old), bool2int(new))
}

func (b *atomicBool) Swap(value bool) bool {
	return atomic.SwapUintptr((*uintptr)(b), bool2int(value)) != 0
}
