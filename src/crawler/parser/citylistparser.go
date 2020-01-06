package parser

import (
	"regexp"
	"crawler/engine"
	)


const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
var expr=`<a (target="_blank" )?href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-[0-9a-z]{8}>([^<]+)</a>`
type CityListParser struct {

}

func (p CityListParser) Parse (contents []byte) engine.ParseResult  {
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		result.Items = append(result.Items, string(c[2])) //城市名字
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			Parser:  CityParser{},
		})
	}

	return result
}


