package main

import (
	"fmt"
	"go-learning/step3/crawler/fetcher"
	"regexp"
)
func printCityList(content []byte){
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/beijing" data-v-5fa74e39>北京</a>
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	submatch := compile.FindAllSubmatch(content, -1) // -1 所有
	for _, matches := range submatch {
		//打印
		fmt.Printf("City:%s URL:%s\n", matches[2], matches[1])
	}

	//可以看到匹配个数为470个
	fmt.Printf("Matches count: %d\n", len(submatch))
}
func main() {
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", resp.StatusCode)
	//	return
	//} else {
	//	all, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		panic(err)
	//	}
	//	printCityList(all)
	//	//fmt.Printf("%s\n", all)
	//}
	all, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	printCityList(all)
}
