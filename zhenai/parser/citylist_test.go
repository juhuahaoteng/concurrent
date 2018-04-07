package parser

import (
	"crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, e := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if e != nil {
		panic(e)
	}
	result := ParseCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but had %d", resultSize, len(result.Requests))

	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+"items; but had %d", resultSize, len(result.Items))

	}
}
