package Lock

import (
	"sync"
	"testing"
)

func TestFifo_SpinLock_Lock(t *testing.T) {
	var w sync.WaitGroup
	n:=0
	var l = newFifo_SpinLock()
	for i:=0;i<7;i++{
		w.Add(1)
		go routine(	l,&w,1-i%2*2,&n)
	}
	w.Wait()
	println(n)
}
func routine(l *Fifo_SpinLock,w *sync.WaitGroup,i int,n *int){
	defer w.Add(-1)
	for t:=0;t<1000000;t++{
		func(){
			l.Lock()
			defer l.UnLock()
			*n +=i
		}()
	}
}