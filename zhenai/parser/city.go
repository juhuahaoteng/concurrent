package parser

import (
	"crawler/types"
	"regexp"
)

const cityRe = `"(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a>`

func ParseCity(contents []byte) types.ParseResult {
	rep := regexp.MustCompile(cityRe)
	matches := rep.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, types.Request{
			Url:        string(m[1]),
			ParserFunc: parseProfile,
		})
	}
	return result
}
