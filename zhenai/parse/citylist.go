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
	//[][][]三维slice
	result :=engine.ParseResult{}
	limit :=10
	for _,m :=range matches{

		result.Items = append(result.Items,string(m[2]))
		//fmt.Println(m[0])
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:ParseCity,
		})
		limit --
		if(limit==0){
			break
		}
	}
    return result
}