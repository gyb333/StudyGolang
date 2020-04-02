package main

import (
	"strings"
	"fmt"
	"bytes"
)
//go test -bench=. -benchmem

func stringsBuilder()  {
	var str strings.Builder

	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}

	fmt.Println(str.String())
}

func bytesBuffer(){
	var buffer bytes.Buffer

	for i := 0; i < 1000; i++ {
		buffer.WriteString("a")
	}

	fmt.Println(buffer.String())
}

func copyStrings(){
	bs := make([]byte, 1000)
	bl := 0
	for n := 0; n < 1000; n++ {
		bl += copy(bs[bl:], "a")
	}
	fmt.Println(string(bs))
}

func appendStrings()  {
	bs := make([]byte, 1000)
	for n := 0; n < 1000; n++ {
		bs = append(bs,'a')
	}
	fmt.Println(string(bs))
}

func stringsQuate(){
	var result string

	for i := 0; i < 1000; i++ {
		result += "a"
	}

	fmt.Println(result)
}

func stringsRepeat(){
	fmt.Println(strings.Repeat("x",1000))
}

func main() {

}
