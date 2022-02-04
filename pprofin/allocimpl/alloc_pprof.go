package allocimpl

import "fmt"

type PprofAlloc struct {
	Buf [][]byte
}

func (p *PprofAlloc) DoPprof() {
	fmt.Printf("%s\n", "PprofAlloc")
	//TODO: alloc 内存优化
	for i := 0; i < 1024; i++ {
		//每次申请一个 G
		p.Buf = append(p.Buf, make([]byte, 1024*1024))
	}
}
