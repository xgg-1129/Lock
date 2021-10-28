package Lock

import (
	"sync/atomic"
)

type SpinLock int32

func (l *SpinLock) Lock()  {
	if !atomic.CompareAndSwapInt32((*int32)(l),0,1){}

}
func (l *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(l),0)
}

