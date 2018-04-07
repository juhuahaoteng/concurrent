package parser

import (
	"crawler/types"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) types.ParseResult {
	rep := regexp.MustCompile(cityListRe)
	matches := rep.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, types.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
