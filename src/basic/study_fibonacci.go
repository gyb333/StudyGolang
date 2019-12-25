package basic

import (
	"time"
	"math/big"
	"fmt"
	"sync"
)

func FibonacciMain() {
	var n, Number = 1, 100000
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fibonacciTest(n, Number)
		wg.Done()
	}()
	go func() {
		fibonacciClosureTest(n, Number)
		wg.Done()
	}()

	go func() {
		fibonacciChanTest(n, Number)
		wg.Done()
	}()
	wg.Wait()
}


func fibonacciTest(n, Number int) {
	start := time.Now()
	var result *big.Int
	for i := 0; i < n; i++ {
		result = Fibonacci(Number)
	}
	fmt.Println("fibonacci 		  took this amount of time:", time.Since(start).Seconds()/float64(n), "s   ", result)
}

func fibonacciClosureTest(n, Number int) {
	start := time.Now()
	var result *big.Int
	for i := 0; i < n; i++ {
		result = FibonacciClosure(Number)
	}
	fmt.Println("fibonacci Closure took this amount of time:", time.Since(start).Seconds()/float64(n), "s   ", result)
}

func fibonacciChanTest(n, Number int) {
	start := time.Now()
	var result *big.Int
	for i := 0; i <= n; i++ {
		//start := time.Now()
		result = FibonacciChanBig(Number)
		//fmt.Printf("fibonacci(%d) Chan took this amount of time:%fs,result:%d\n", i,time.Since(start).Seconds(),  result)
	}
	fmt.Println("fibonacci Chan    took this amount of time:", time.Since(start).Seconds()/float64(n), "s   ", result)
}

func Fibonacci(n int) *big.Int {
	x, y := big.NewInt(0), big.NewInt(1)
	for i := 0; i < n; i++ {
		x, y = y, x.Add(x, y)
	}
	return x
}

func fibonacciChan(n int) (result int) {
	channel := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < n; i++ {
			result = <-channel
		}
		quit <- true
	}()
	func(channel chan int, quit chan bool) {
		x, y := 0, 1
		for {
			select {
			case channel <- y:
				x, y = y, x+y
			case <-quit:
				return
			}
		}
	}(channel, quit)
	return
}

func FibonacciChanBig(n int) (result *big.Int) {
	result = big.NewInt(0)
	channel := make(chan *big.Int, 30)
	quit := make(chan bool)
	go func() {
		for i := 0; i < n; i++ {
			result = <-channel
		}
		quit <- true
	}()
	func(channel chan *big.Int, quit chan bool) {
		x, y := big.NewInt(0), big.NewInt(1)
		for {
			select {
			case channel <- y:
				x, y = y, x.Add(x, y)
			case <-quit:
				return
			}
		}
	}(channel, quit)
	return
}

func FibonacciClosure(n int) (r *big.Int) {
	if n < 1 {
		return big.NewInt(0)
	}
	f := func() func() *big.Int {
		x, y := big.NewInt(0), big.NewInt(1)
		return func() *big.Int {
			x, y = y, x.Add(x, y)
			return x
		}
	}()
	for i := 0; i < n; i++ {
		r = f()
	}
	return
}

func fibonacciBig(n int) (r *big.Int) {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	f := func() func() *big.Int {
		v, s := big.NewInt(0), big.NewInt(1)
		return func() *big.Int {
			var tmp big.Int
			tmp.Set(s)
			s.Add(s, v)
			v = &tmp
			return s
		}
	}()

	for i := 1; i < n; i++ {
		r = f()
	}
	return

}

func fibonacciRecursion(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciRecursion(n-2) + fibonacciRecursion(n-1)
	}
	return
}
