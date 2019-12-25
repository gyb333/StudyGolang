package basic

import (
	"fmt"
	. "unsafe"
	"reflect"
)

/**
类型															大小
bool														1个字节
intN, uintN, floatN, complexN							N/8个字节(例如float64是8个字节)
int, uint, uintptr										1个机器字
*T														1个机器字
string													2个机器字(data,len)
[]T														3个机器字(data,len,cap)
map														1个机器字
func													1个机器字
chan													1个机器字
interface												2个机器字(type,value)
unsafe.Alignof 函数返回对应参数的类型需要对齐的倍数. 和 Sizeof 类似, Alignof 也是返回一个常量表达式, 对应一个常量. 通常情况下布尔和数字类型需要对齐到它们本身的大小(最多8个字节), 其它的类型对齐到机器字大小.

unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
**/




func StudyStruct()  {
	baseStruct()
	structOffset()
}



func baseStruct()  {
	var b bool
	var i8 int8
	var u8 uint8
	var i16 int16
	var u16 uint16
	var i32 int32
	var u32 uint32
	var i int
	var u uint
	var f32 float32
	var f64 float64
	var ptr uintptr
	var p *int
	var str string
	var array [8]byte
	var slice []int
	var m map[string]int
	var c  chan struct{}
	f :=func (){}		//匿名函数
	var s struct{}		//匿名结构体
	var in interface{}	//匿名接口

	fmt.Println(reflect.TypeOf(b),Sizeof(b),Alignof(b))					//bool 1 1
	fmt.Println(reflect.TypeOf(i8),Sizeof(i8),Alignof(i8))				//int8 1 1
	fmt.Println(reflect.TypeOf(u8),Sizeof(u8),Alignof(u8))				//uint8 1 1
	fmt.Println(reflect.TypeOf(i16),Sizeof(i16),Alignof(i16))			//int16 2 2
	fmt.Println(reflect.TypeOf(u16),Sizeof(u16),Alignof(u16))			//uint16 2 2
	fmt.Println(reflect.TypeOf(i32),Sizeof(i32),Alignof(i32))			//int32 4 4
	fmt.Println(reflect.TypeOf(u32),Sizeof(u32),Alignof(u32))			//uint32 4 4
	fmt.Println(reflect.TypeOf(i),Sizeof(i),Alignof(i))					//int 8 8
	fmt.Println(reflect.TypeOf(u),Sizeof(u),Alignof(u))					//uint 8 8
	fmt.Println(reflect.TypeOf(f32),Sizeof(f32),Alignof(f32))			//float32 4 4
	fmt.Println(reflect.TypeOf(f64),Sizeof(f64),Alignof(f64))			//float64 8 8

	fmt.Println(reflect.TypeOf(ptr),Sizeof(ptr),Alignof(ptr))			//uintptr 8 8
	fmt.Println(reflect.TypeOf(p),Sizeof(p),Alignof(p))					//*int 8 8
	fmt.Println(reflect.TypeOf(array),Sizeof(array),Alignof(array))		//[8]uint8 8 1
	fmt.Println(reflect.TypeOf(str),Sizeof(str),Alignof(str))			//string 16 8


	fmt.Println(reflect.TypeOf(slice),Sizeof(slice),Alignof(slice))		//[]int 24 8
	fmt.Println(reflect.TypeOf(m),Sizeof(m),Alignof(m))					//map[string]int 8 8
	fmt.Println(reflect.TypeOf(c),Sizeof(c),Alignof(c))					//chan struct {} 8 8

	fmt.Println(reflect.TypeOf(f),Sizeof(f),Alignof(f))					//func() 8 8
	fmt.Println(reflect.TypeOf(s),Sizeof(s),Alignof(s))					//struct {} 0 1
	fmt.Println(reflect.TypeOf(in),Sizeof(in),Alignof(in))				//<nil> 16 8

}



var x struct {
	a bool
	b int16
	c []int
}

func structOffset() {
	fmt.Println(Sizeof(x), Alignof(x))
	fmt.Println(Sizeof(x.a), Alignof(x.a), Offsetof(x.a))
	fmt.Println(Sizeof(x.b), Alignof(x.b), Offsetof(x.b))
	fmt.Println(Sizeof(x.c), Alignof(x.c), Offsetof(x.c))
	// 和 pb := &x.b 等价
	pb := (*int16)(Pointer(uintptr(Pointer(&x)) + Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
}