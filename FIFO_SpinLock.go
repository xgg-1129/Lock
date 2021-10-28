package Lock

import (
	"sync"
	"sync/atomic"
)

type Fifo_SpinLock struct {
	//真正的自旋锁
	mu sync.Mutex
	//当前自旋锁的拥有者
	owner uint32
	//下一个自旋锁的拥有者
	next uint32
}

func (l *Fifo_SpinLock) Lock(){
	l.mu.Lock()
	var self uint32
	atomic.StoreUint32(&self,l.next)
	atomic.AddUint32(&(l.next),1)
	l.mu.Unlock()
	for self != l.owner{
	}
	return
}
func (l *Fifo_SpinLock) UnLock(){
	atomic.AddUint32(&(l.owner),1)
}
