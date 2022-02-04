package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"pprof-study/pprofin"
	"pprof-study/pprofin/allocimpl"
	"pprof-study/pprofin/blockimpl"
	"pprof-study/pprofin/groutineimpl"
	"pprof-study/pprofin/profileimpl"
	"runtime"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	//初始化测试内
	pprofs := []pprofin.Pprof{
		&allocimpl.PprofAlloc{Buf: make([][]byte, 0)}, // 内存优化
		&groutineimpl.PprofGoroutine{},                // 协程优化
		&profileimpl.PprofProfile{},                   //CPU优化
		&blockimpl.PproBlock{},                        //锁阻塞优化
	}
	fmt.Printf("%s\n", "---start---")
	for {
		for _, p := range pprofs {
			p.DoPprof()
			time.Sleep(time.Second)
		}
	}
}
