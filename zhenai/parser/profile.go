package parser

import (
	"crawler/model"
	"regexp"
	"strconv"
	"crawler/types"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<])</td>`)

func parseProfile(contents []byte) types.ParseResult {
	profile := model.Profile{}
	profile.Age, _ = strconv.Atoi(extractString(contents, ageRe))
	result := types.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}
