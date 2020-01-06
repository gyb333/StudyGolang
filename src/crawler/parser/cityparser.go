package parser

import (
	"crawler/engine"
	"regexp"
)

type CityParser struct {
}

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func (p CityParser) Parse (contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		name:=string(c[2])
		result.Items = append(result.Items, "User:"+name) //用户名字

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			Parser: ProfileParser{Name:name},
		})
	}

	return result
}
