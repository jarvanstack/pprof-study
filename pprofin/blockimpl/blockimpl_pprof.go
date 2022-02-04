package blockimpl

import (
	"fmt"
	"sync"
	"time"
)

type PproBlock struct {
	Buf [][]byte
}

func (p *PproBlock) DoPprof() {
	fmt.Printf("%s\n", "PproBlock")
	//TODO: 锁阻塞优化
	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}
