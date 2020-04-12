package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func Fetch(url string) ([] byte, error) {
	/*
	   resp, err := http.Get(url)
	   if err != nil {
	       return nil, err
	   }
	*/
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp, _ := http.DefaultClient.Do(request)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//if resp.StatusCode != http.StatusAccepted  {
		//fmt.Println("resp:", resp)
		return nil, fmt.Errorf("error:status code:%d", resp.StatusCode)
	} else {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return all, err
	}

}
