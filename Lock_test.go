package Lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestFifo_SpinLock_Lock(t *testing.T) {
	var w sync.WaitGroup
	n:=0
	var l = NewRWLock()
	for i:=0;i<5;i++{
		w.Add(1)
		go routine(	l,&w,1-i%2*2,&n)
	}
	for i:=0;i<5;i++{
		go func() {
			w.Add(1)
			for j:=0;j<10000;j++{
				func(){
					l.WLock()
					defer l.WUnLock()
					fmt.Printf("n=%d\n",n)
				}()
			}
			w.Done()
		}()
	}
	w.Wait()
	println(n)
}
func routine(l *rwLock,w *sync.WaitGroup,i int,n *int){
	defer w.Add(-1)
	for t:=0;t<1000000;t++{
		func(){
			l.RLock()
			defer l.RUnLock()
			*n +=i
		}()
	}
}