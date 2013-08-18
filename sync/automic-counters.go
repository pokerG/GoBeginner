package main 

import(
	"fmt"
	"time"
	"sync/atomic"
)

func main(){
	var ops uint64 = 0

	for i := 0; i < 50; i++{
		go func(){
			for{
				time.Sleep(time.Microsecond)
				atomic.AddUint64(&ops,1)
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("&ops:" , ops)
	fmt.Println("ops: ", opsFinal)
}