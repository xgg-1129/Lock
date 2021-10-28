package Lock

import (
	"sync/atomic"
	"time"
)

type XggMutex int32

func (l *XggMutex) Lock()  {
	if !atomic.CompareAndSwapInt32((*int32)(l),0,1){
		time.Sleep(time.Millisecond)
	}
}
func (l *XggMutex) Unlock() {
	atomic.StoreInt32((*int32)(l),0)
}


