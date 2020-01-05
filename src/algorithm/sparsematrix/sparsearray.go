package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"utils"
	"strings"
	)

//结构体需要 序列化则属性需要公开大写
type ValueNode struct {
	Row int
	Col int
	Val int
}

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, row := range chessMap {
		for _, v := range row {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}
	firstNode := ValueNode{Row: 11, Col: 11, Val: 0}
	var sparseArray []ValueNode
	sparseArray = append(sparseArray, firstNode)

	for i, row := range chessMap {
		for j, v := range row {
			if v != 0 {
				sparseArray = append(sparseArray, ValueNode{i, j, v})
			}
		}
	}
	fmt.Println("当前的稀疏数组：")
	for i, v := range sparseArray {
		fmt.Printf("%d:\t%d \t%d \t%d\n", i, v.Row, v.Col, v.Val)

	}

	//将内容存盘文件
	var fileSerializePath="src/algorithm/sparsematrix/chessmap.SerializeData"
	writeSerializeFile(fileSerializePath,sparseArray)
	result,_ :=readerSerializeFile(fileSerializePath)

	fmt.Println(restore(result))

	var filePath="src/algorithm/sparsematrix/chessmap.Data"
	writeFile(filePath,sparseArray)
	result=readFile(filePath)
	fmt.Println(restore(result))

	var filePathBufio = "src/algorithm/sparsematrix/chessmap.BufioData"
	writeFileBufio(filePathBufio, sparseArray)
	result=readFileBufio(filePathBufio)
	fmt.Println(restore(result))

}

func restore(sparseArray []ValueNode) (data [11][11]int) {
	for i, v := range sparseArray {
		if i != 0 {
			data[v.Row][v.Col] = v.Val
		}
	}
	return
}

func writeSerializeFile(filename string, array []ValueNode) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	data, err := json.Marshal(array)
	if err != nil {
		return err
	}
	e := json.NewEncoder(file)
	return e.Encode(data)

}
func readerSerializeFile(filename string) (result []ValueNode, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	e := json.NewDecoder(file)
	var data []byte
	err = e.Decode(&data)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &result)
	return
}

func writeFile(filename string, array []ValueNode) {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%d %d\n", len(array), len(utils.GetFieldName(ValueNode{})))
	for _, v := range array {
		fmt.Fprintf(file, "%d %d %d\n", v.Row, v.Col, v.Val)
	}

}

func readFile(filename string) (result []ValueNode) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	for i := 0; i < row; i++ {
		var val ValueNode
		fmt.Fscanf(file, "%d %d %d\n", &val.Row, &val.Col, &val.Val)
		result = append(result, val)
	}
	return
}

func writeFileBufio(filename string, array []ValueNode) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := bufio.NewWriter(file)
	for _, v := range array {
		buf.WriteString(strconv.FormatInt(int64(v.Row), 10))
		buf.WriteString(" ")
		buf.WriteString(strconv.FormatInt(int64(v.Col), 10))
		buf.WriteString(" ")
		buf.WriteString(strconv.FormatInt(int64(v.Val), 10))
		buf.WriteString("\n")
	}
	buf.Flush()

}

func readFileBufio(filename string) (result []ValueNode){
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line,_, err := buf.ReadLine()
		if err == io.EOF {
			fmt.Println("read the file finished")
			break
		}
		lines :=strings.Split(string(line)," ")
		var node ValueNode
		row,_:= strconv.ParseInt(lines[0] ,10,64)
		node.Row=int(row)
		col,_:= strconv.ParseInt(lines[1] ,10,64)
		node.Col=int(col)
		val,_:= strconv.ParseInt(lines[2] ,10,64)
		node.Val=int(val)
		result=append(result,node)
	}
	return

}
