package Lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type rwLock struct {
	initialValue int32
	value int32
	mu sync.RWMutex
}

func NewRWLock()*rwLock {
	r:=new(rwLock)
	r.initialValue=1024
	r.value=r.initialValue
	return r
}
//获取读锁
func (l *rwLock) WLock() {
	atomic.AddInt32(&l.value,-1)

	for l.value<0{

	}
	return
}
func (l *rwLock) RLock() {
	for !atomic.CompareAndSwapInt32(&l.value,l.initialValue,0){}
	return
}
func (l *rwLock) RUnLock(){
	atomic.AddInt32(&l.value,l.initialValue)
}
func (l *rwLock) WUnLock(){
	atomic.AddInt32(&l.value,1)
	fmt.Printf("value+1后%d\n",l.value)
}