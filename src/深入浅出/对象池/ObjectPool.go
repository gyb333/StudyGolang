package main

import (
	"sync"
	"fmt"
	"time"
)

/**
对象池能提高内存复用,减少内存申请次数,甚至能降低CPU消耗，是高并发项目优化不可缺少的手法之一
fmt.Sprintf()
 */



// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 1024)
		return &buf
	},
}
var ch = make(chan []byte,1000)

func main() {
	go func(){
		for msg := range ch {
			fmt.Println("recv msg",msg)
			msg = msg[:0]
			bytePool.Put(msg)
		}
	}()

	a := time.Now()
	// 不使用对象池
	for i := 0; i < 100000000; i++{
		lineBuf :=  make([]byte, 1024)
		lineBuf = append(lineBuf, 1)
		lineBuf = append(lineBuf, 2)
		lineBuf = append(lineBuf, 3)
		lineBuf = append(lineBuf, 4)
		_ = lineBuf
	}
	b := time.Now()
	p :=bytePool.Get().(*[]byte)
	// 使用对象池
	for i:=0;i<=100000000;i++ {
		lineBuf := *p
		lineBuf = append(lineBuf, 1)
		lineBuf = append(lineBuf, 2)
		lineBuf = append(lineBuf, 3)
		lineBuf = append(lineBuf, 4)

		if i==0 {
			//fmt.Println(string(lineBuf))
			ch <- lineBuf
		}
	}

	c := time.Now()
	fmt.Println("without pool ", b.Sub(a).Seconds(), "s")
	fmt.Println("with    pool ", c.Sub(b).Seconds(), "s")
}

// without pool  34 s
// with    pool  24 s