package basic

import "runtime"

func ChanMain()  {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			j <- i
		}
		close(j)
		i <- struct{}{}
	}(c, ci)
	println("NumGoroutine=", runtime.NumGoroutine())
	<-c
	println("NumGoroutine=", runtime.NumGoroutine())
	for v := range ci {
		print(v)
	}
}
