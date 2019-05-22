package engine

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