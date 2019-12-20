package basic

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"

	"unicode"
)

func StudyString() {

	//basicString()
	stringStruct()
	stringToByteSlice()
	stringToRuneSlice()
	//stringJoin()

}

var str = "Hello 世界！"

func basicString() {
	fmt.Printf("%p,%T,%d,%v,%s\n", &str, str, unsafe.Sizeof(""), str, str)
	fmt.Printf("%d,%d\n", len(str), utf8.RuneCountInString(str))
	for i, v := range str {
		fmt.Printf("%d\t%c\t%q\t%d\t%#x \n", i, v, v, v, v)
	}
}

func stringStruct() {

	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	p := unsafe.Pointer(sh.Data)
	fmt.Println(&str, str, len(str), *sh, p)

	bs := (*[15]byte)(p)
	fmt.Printf("%p,%v,%s,%c\n", unsafe.Pointer(bs), *bs, *bs, bs[0])
	for i,v:=range bs{
		fmt.Println(i,v)		//字符串底層數組不能修改
	}




}

func stringToByteSlice()  {
	var bs =[]byte(str)
	fmt.Printf("%p,%T,%d,%d,%s,%v\n",&bs,bs,len(bs),cap(bs),bs,bs)

	rstr := *(*string)(unsafe.Pointer(&bs))
	fmt.Println(rstr)


	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	p := unsafe.Pointer(sh.Data)
	c := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		// 指针类型转换通过unsafe包
		c[i] = *(*byte)(unsafe.Pointer(uintptr(p) + uintptr(i))) // 指针运算只能通过uintptr
	}
	ss := *(*string)(unsafe.Pointer(&c))
	fmt.Println(c, ss)

}


func stringToRuneSlice()  {
	var ss =[]rune(str)
	fmt.Printf("%p,%T,%d,%d,%s,%v\n",&ss,ss,len(ss),cap(ss),(string)(ss),ss)

	rstr := (string)(ss)
	fmt.Println(rstr)

	length := utf8.RuneCountInString(str)
	rs := make([]rune, length)
	for i, v := range ss {
		if !unicode.Is(unicode.Han, v) && i==0 {
			rs[i]=v+32
		}else {
			rs[i] = v
		}
	}
	fmt.Println(rs,string(rs))
}

func stringJoin() {
	fmt.Println(strings.Join([]string{"Hello", "世界！"}, " "))

	var buf bytes.Buffer
	buf.WriteString("Hello ")
	buf.WriteString("世界！")
	fmt.Println(buf.String())

	var sb strings.Builder
	sb.WriteString("Hello ")
	sb.WriteString("世界！")
	fmt.Println(sb.String())
}
