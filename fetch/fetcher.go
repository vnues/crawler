package fetch

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

//声明函数需要参数名 返回的=需要声明类型 error是全局类型
func Fetcher(url sting)([]byte,error){

	   resp,err:=http.Get(url)
	   //假设网址没有不能存在根本就没有请求
	   if err !=nil{
	   	 return nil,err
	   }
	   //defer的写法位置不能靠后---gegoole
	defer resp.Body.Close()
	   //有请求但是不成功
	   if resp.StatusCode !=http.StatusOK{
	   	 return nil,fmt.Errorf("Wrong status code :%d",resp.StatusCode)
	   }
	   bodyReader := bufio.NewReader(resp.Body)
	   /*
	    NewReader returns a new Reader whose buffer has the default size.
	   func NewReader(rd io.Reader) *Reader {
	   	return NewReaderSize(rd, defaultBufSize)
	   }
	   */
       e :=determineEncoding(bodyReader)
       utf8Reader :=transform.NewReader(bodyReader,e.NewDecoder())
       return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	//传的是指针但是不会改变--阅读源代码
	  bytes,err := r.Peek(1024)
      if err !=nil{
      	 log.Printf("Fetcher error:%v",err)
      	 //返回默认的
      	 return unicode.UTF8
	  }
	  e,_,_ :=charset.DetermineEncoding(bytes,"")
	  return e
}