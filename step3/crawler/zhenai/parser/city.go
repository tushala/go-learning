package parser

import (
	"go-learning/step3/crawler/engine"
	"regexp"
)
//获取用户信息格式
const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
//contents为城市页面地址 从每个城市第一页中筛选信息
func ParseCity(contents []byte) engine.ParseResult{
	//确定要查找的格式
	re := regexp.MustCompile(cityRe)
	//搜索全部与格式相同的信息
	matches := re.FindAllSubmatch(contents, -1)

	//创建结构体进行存放
	result := engine.ParseResult{}
	//把第一张页面中的用户名及地址取出
	for _, m := range matches{
		//m[2]为用户名
		name := string(m[2])

		//在结构体中存入所有昵称名字 并标识为User
		result.Items = append(result.Items, "User "+name)
		//这个函数中最关键的点
		//将函数ParseProfile作为返回值 即确定了昵称 由没有改动结构体
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]), //用户页面地址
			func(c []byte) engine.ParseResult{ //用户信息
				return ParserProfile(c, name)
			},
		})
	}
	return result
}