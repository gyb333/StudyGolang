package parser

import (
	"../../engine"
	"../../model"
	"regexp"
	"strconv"
	"fmt"
)

// \d就是所有的数字0-9
//年龄<div data-v-8b1eac0c="" class="m-btn purple">未婚</div>
var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c class="m-btn purple">(\d+)岁</div>`)
//身高
var heightRe = regexp.MustCompile(`<div data-v-[0-9a-z]{8} class="m-btn purple">(\d+)CM</div>`)
//月收入
var incomeRe = regexp.MustCompile(`<div data-v-[0-9a-z]{8} class="m-btn purple">月收入:([^<]+)</div>`)
//体重
var weightRe = regexp.MustCompile(`<div data-v-[0-9a-z]{8} class="m-btn purple">(\d+)KG</span></div>`)
//性别
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
//星座
var constellationRe = regexp.MustCompile(`<div data-v-[0-9a-z]{8} class="m-btn purple">([^<]+座)</span></div>`)
//婚况
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
//学历
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
//职业
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
//籍贯
var nativePlaceRe = regexp.MustCompile(`<div data-v-[0-9a-z]{8} class="m-btn pink">籍贯:([^<]+)</div>`)
//住房条件
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
//是否购车
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

// 用户解析器
// 得到用户详情页的各种结构化数据
func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	fmt.Printf("%s\n",contents)

	profile := model.Profile{}
	profile.Name = name

	profile.Age = extractInt(contents, ageRe)
	profile.Height = extractInt(contents, heightRe)
	profile.Weight = extractInt(contents, weightRe)

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.NativePlace = extractString(contents, nativePlaceRe)
	profile.House = extractString(contents, houseRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Constellation = extractString(contents, constellationRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	//matches := guessRe.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	result.Requests = append(result.Requests,
	//		engine.Request{
	//			Url:        string(m[1]),
	//			ParserFunc: ProfileParser(string(m[2])),
	//		})
	//}

	return result
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(contents)
	str := ""
	if len(match) >= 2 {
		str = string(match[1])
	}
	number, _ := strconv.Atoi(str)
	return number
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}

