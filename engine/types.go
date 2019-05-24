package engine

//以[]形式声明的都是切片类型
//append用来将元素添加到切片末尾并返回结果。
//定义种子（能放进去执行队列自动执行）对象:这个种子需要自带url,解析函数

type  Request struct {
	Url string
	//匿名函数 解析参数，go函数参数并不需要指定名称
	//ParseFunc func([]byte) ParseResult这样写不是声明函数 声明函数需要参数名称的
	ParseFunc func([]byte) ParseResult
}

//定义函数返回的结果
//函数需要返回解析后的并且封装后的种子对象(肯定是多个种子)
//需要的有用信息items-定义成任意类型
//这个是对象的属性 属性之间村值不用保存同步
type ParseResult struct {
	  Requests []Request
	 //获得的真正有价值的爬虫信息
      Items []interface{}
}

func NilParser([] byte) ParseResult{
	//返回空的东西给它
	return  ParseResult{}
}