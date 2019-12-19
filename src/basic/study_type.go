package basic

import (
	"fmt"
	"unsafe"
		"bytes"
	"encoding/binary"
)

/**
值类型 深拷贝：基本数据类型  数组
引用类型 浅拷贝: 切片 字典 通道
 */
func BasicType()  {
	unsafeData()
	//arrayType()
	runeType()
	//stringType()
}



const ptrIntSize = unsafe.Sizeof((*int)(nil))
const ptrIntAlign = unsafe.Alignof((*int)(nil))

func unsafeData()  {
	fmt.Println(ptrIntAlign,ptrIntSize)
}









//复合数据类型：指针 数组 切片 字典 通道 结构 接口

func arrayType()  {
	var arr [10]int
	fmt.Printf("%T,%v,%p,%d,%d\n",arr,arr,&arr,len(arr),cap(arr))
	for i:=0;i<len(arr);i++{
		arr[i]=i
	}
	fmt.Println(&arr,unsafe.Pointer(&arr),&arr[1],"\n")
	arr2:= [10]int{}
	fmt.Printf("%T,%v,%p,%d,%d\n",arr2,arr2,&arr2,len(arr2),cap(arr2))
	for i,v :=range arr{
		pb := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr2))+uintptr(i)*ptrIntAlign))
		*pb =v+1
	}
	fmt.Println(&arr2,unsafe.Pointer(&arr2),&arr2[1],"\n")

	arr3 :=arr2
	fmt.Printf("%T,%v,%p,%d,%d\n",arr2,arr2,&arr2,len(arr2),cap(arr2))
	for i,v :=range arr3{
		arr3[i]= v-1
	}
	fmt.Println(&arr3,unsafe.Pointer(&arr3),&arr3[1],"\n")

}

func toBytes() []byte {
	bb :=bytes.NewBuffer(nil)
	binary.Write(bb, binary.BigEndian, 'h')
	bs := bb.Bytes()
	fmt.Printf("%#X,%d,%d\n",bs,len(bs),cap(bs))
	return bs
}


//append copy
func sliceType()  {

}

