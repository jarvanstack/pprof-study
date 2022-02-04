package profileimpl

import (
	"fmt"
)

type PprofProfile struct {
}

func (p *PprofProfile) DoPprof() {
	fmt.Printf("%s\n", "PprofProfile")
	//TODO: cpu 性能瓶颈
	for i := 0; i < 1000_000_000; i++ {
		f := 0.003
		f = f*f*f/f + f
		f = f + 1
	}
}
