package basic

import (
	"fmt"
	"unsafe"
	"unicode/utf8"
	"strings"
	"bytes"
)

func StudyString(){

	//basicString()
	stringStruct()
	//stringSlice()
	//stringJoin()

}
var str  ="Hello 世界！"

func basicString()  {
	fmt.Printf("%p,%T,%v,%s\n",&str,str,str,str)
	fmt.Printf("%d,%d,%d\n",unsafe.Sizeof(str),len(str),utf8.RuneCountInString(str))
	for i,v:=range str{
		fmt.Printf("%d\t%c\t%q\t%d\t%#x \n",i,v,v,v,v)
	}
}

func stringStruct()  {
	p:=unsafe.Pointer(&str)
	println(p)
}

func stringSlice()  {
	var s =str[:]
	fmt.Println(s)
	
}





func stringJoin(){
	fmt.Println(strings.Join([]string{"Hello","世界！"}," "))

	b :=bytes.NewBuffer(nil)
	b.WriteString("Hello ")
	b.WriteString("世界！")
	fmt.Println(b.String())
}
