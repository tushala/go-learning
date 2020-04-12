package parser

import (
	"go-learning/step3/crawler/engine"
	"regexp"
)
const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

//提取所需内容 如网页地址 城市名称
func ParseCityList(contents []byte)engine.ParseResult{
	//设置被提取者所需要的条件
	re :=regexp.MustCompile(citylistRe)
	//从网页文件中提取所需文件
	matches := re.FindAllSubmatch(contents, -1)
	//提取出来的内容需要一个接受者 result定义类型
	result := engine.ParseResult{}
	limit := 10
	for _, m := range  matches{
		//将提取出来的信息按照顺序放入变量result中
		result.Items = append(result.Items,"City" + string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			string(m[1]),
			ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	//返回result result内部是数组
	return result
}