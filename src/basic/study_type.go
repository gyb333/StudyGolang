package basic

import (
	"fmt"
	"unsafe"
		"bytes"
	"encoding/binary"
)


func init() {
	fmt.Println("basic package init")
}

const (
	Monday    = iota //0
	Tuesday          //默认自动加1 1
	Wednesday        //2
	Thursday
	Friday
	Saturday
	Sunday
)
const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
)




const ptrIntSize = unsafe.Sizeof((*int)(nil))
const ptrIntAlign = unsafe.Alignof((*int)(nil))

func unsafeData()  {
	fmt.Println(ptrIntAlign,ptrIntSize)
}


type SliceInt []int

func (s SliceInt) Sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}
func SliceInt_Sum(s SliceInt) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

/**
值类型 深拷贝：基本数据类型  数组
引用类型 浅拷贝: 切片 字典 通道
 */
func BasicType()  {
	unsafeData()
	var s SliceInt = []int{1, 2, 3, 4}
	println(s.Sum())
	println(SliceInt_Sum(s))

}








func toBytes() []byte {
	bb :=bytes.NewBuffer(nil)
	binary.Write(bb, binary.BigEndian, 'h')
	bs := bb.Bytes()
	fmt.Printf("%#X,%d,%d\n",bs,len(bs),cap(bs))
	return bs
}




