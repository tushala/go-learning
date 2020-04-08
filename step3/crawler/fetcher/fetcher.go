package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("Error: status code", resp.StatusCode)
	} else {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("%s\n", all)
		return all, err
	}

}