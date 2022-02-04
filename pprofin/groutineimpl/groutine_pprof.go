package groutineimpl

import (
	"fmt"
	"time"
)

type PprofGoroutine struct {
}

func (p *PprofGoroutine) DoPprof() {
	fmt.Printf("%s\n", "PprofGoroutine")
	//TODO: groutine 泄露
	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
	}
}
