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

