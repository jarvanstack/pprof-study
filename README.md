# pprof-study 

pprof 性能分析优化, 本项目是一个实战项目, 下面将带你通过 web, 交互界面, 性能图/火焰图 找到性能优化位置并优化代码


## quick start 

```bash
git clone git@github.com:dengjiawen8955/pprof-study.git

cd pprof-study

go mod tidy

go run mian.go
```

访问下面的性能分析地址 

http://localhost:6060/debug/pprof


* allocs	内存分配情况的采样信息	
* blocks	阻塞操作情况的采样信息	
* cmdline	显示程序启动命令及参数	
* goroutine	当前所有协程的堆栈信息	
* heap	堆上内存使用情况的采样信息	
* mutex	锁争用情况的采样信息	
* profile	CPU 占用情况的采样信息	
* threadcreate	系统线程创建情况的采样信息	
* trace	程序运行跟踪信息	


## 交互界面

go tool pprof http://localhost:6060/debug/pprof/allocs

## 

## refrence 

[《golang pprof 实战》](https://blog.wolfogre.com/posts/go-ppof-practice/)代码实验用例。