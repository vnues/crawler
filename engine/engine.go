package engine

import ("fmt"
	"log"

	"../fetch"
)

func Run(seeds ...Request){
	 fmt.Println("启动种子引擎")
	 var requests []Request

	 for _,r :=range seeds{
	 	 requests=append(requests,r)
	 }
	 //将种子拿出队列执行
	 for len(requests)>0{
	 	  r :=requests[0]
	 	  requests=requests[1:]
	 	  //爬取页面
		 //body html字符串--byte类型
		  body,err :=fetch.Fetcher(r.Url)
	 	  fmt.Printf("我正在爬取这个地址Fetching %s\n",r.Url)
	 	  if err!=nil{
	 	  	 //发生获取错误跳出这次循环
	 	  	 log.Printf("Fetch error "+"fetching url %s:%v",r.Url,err)
	 	  	 continue
		  }
	 	  //解析
	 	  ParseResult :=r.ParseFunc(body)
	 	  //解析出来的结果里面可能包含种子，需要把种子放进去队列
	 	  //append另外一种形式...表示拼接
	 	  requests =append(requests,ParseResult.Requests...)
	 	  for _,item :=range ParseResult.Items{
	 	  	fmt.Printf("爬取所得到到信息：%v\n",item)
		  }

	 }
}










































//go append方式是slice方法 有两种形式追加和拼接(就是打散切片)

//engine这快是把爬取的数据打印出来并且又把种子放进去队列执行

/*
‘…’ 其实是go的一种语法糖。
它的第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
第二个用法是slice可以被打散进行传递
*/
/*
func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func main(){
var strss= []string{
        "qwr",
        "234",
        "yui",
        "cvbc",
    }
    test1(strss...) //切片被打散传入
}
*/


