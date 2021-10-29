package Lock

import (
	"sync"
	"sync/atomic"
)

type Fifo_SpinLock struct {
	owner uint64
	next uint64

	//mu用来模拟cpu的原子操作
	mu sync.Mutex
}
func newFifo_SpinLock()*Fifo_SpinLock{
	res:=new(Fifo_SpinLock)
	res.owner = 0
	res.next = 0
	return res
}
func (l *Fifo_SpinLock) Lock(){
	l.mu.Lock()
	//获取next和owner
	owner:=l.next
	next:=l.owner
	l.next++
	l.mu.Unlock()
	//获取next和owner
	for owner != next{
		next=l.owner
	}
	return
}
func (l *Fifo_SpinLock) UnLock(){
	atomic.AddUint64(&(l.owner),1)
}
