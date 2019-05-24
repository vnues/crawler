package parse

import (
	"../../engine"
	"fmt"
	"regexp"
)

//我们需要的有价值的信息
//变量在这个包里即使在不同文件也是不能有变量重复（这个包的全局下）
//const cityRe=`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th>`
//
//func ParseCity(contents []byte) engine.ParseResult{
//	 fmt.Println("启动城市解析器---作用获取该城市用户")
//	 re :=regexp.MustCompile(cityListRe)
//	 //实例化这个对象
//	 result :=engine.ParseResult{}
//	 matches :=re.FindAllSubmatch(contents,-1)
//	for _,m :=range matches{
//		 result.Items=append(result.Items,"User:"+string(m[2]))
//		 result.Requests=append(result.Requests,engine.Request{
//		 	Url:string(m[1]),
//		 	ParseFunc:engine.NilParser,
//		 })
//	 }
//	 return result
//}
const cityRe=`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th>`


func ParseCity(contents []byte) engine.ParseResult{
	//fmt.Println(string(contents))
	re :=regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)
	result :=engine.ParseResult{}
	fmt.Println("启动城市详情页解析器")
	for _,m :=range matches {
		name :=string(m[2])
		result.Items = append(result.Items, "User :"+name)

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParseFunc: func(c []byte)engine.ParseResult{
				//这是个闭包
				 return ParseProfile(c,name)
			},
		})
	}
	return result
}
//首先我们在for循环外定义m然后循环后m是最后到值 --但是我们想保留每个循环的变量值所以不能在栈销毁，就得使用闭包了
//我们在for循环内定义一个变量来存姓名信息，为了不然循环后被销毁，我们得return出去，告诉编译器我这个变量还有用