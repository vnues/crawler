package parse

import (
	"../../engine"
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contens []byte) engine.ParseResult{
	fmt.Println("启动城市列表解析器")
	re :=regexp.MustCompile(cityListRe)
	matches :=re.FindAllSubmatch(contens,-1)
	for _,m :=range matches{
		result :=engine.ParseResult{}
		result.Itmes =append(result.Items,string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:engine.NilParser,
		})
	}
    return result
}