package main

import (
	"./engine"
	"./zhenai/parse"
	"fmt"
)

func main(){
	fmt.Println("开始启动项目")
	//我是这种写法？？
	//engine.Run({
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser
	//})

	//首先我们传这种参数过去是受到js语法才会这样写
	//在go语言中 要指定的参数是什么类型，并且传递过去也是这个具体类型的值
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parse.ParseCityList,
	})
}
