package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
		"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
	} else {

		r,e:=determineEncoding(resp.Body)
		utf8Reader := transform.NewReader(r, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s \n", all)

		printCityList(all)
	}

}

func determineEncoding(r io.Reader) (reader *bufio.Reader,e encoding.Encoding) {
	reader =bufio.NewReader(r)
	bytes, err := reader.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ = charset.DetermineEncoding(bytes, "")
	return
}

func printCityList(contents []byte){
	//[^>]*>[^<]+</a>
	var expr=`<a (target="_blank" )?href="http://www.zhenai.com/zhenghun/[0-9a-z]+" data-v-[0-9a-z]{8}>[^<]+</a>`
	re :=regexp.MustCompile(expr)
	matches :=re.FindAll(contents,-1)
	for i,v:=range matches{
		fmt.Printf("%d,%s \n",i,v)
	}
}