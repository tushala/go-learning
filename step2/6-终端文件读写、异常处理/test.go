// 自定义错误
package main

import (
	"errors"
	"fmt"
)

var PathError2 = errors.New("Path Error")
type PathError struct {
	Op string
	Path string
	Err string
}
func (e * PathError) Error() string{
	return e.Op + " " + e.Path + ": " + e.Err
}

func test() error{
	return &PathError{
		Op:   "op",
		Path: "path",
	}
}
func main() {
	fmt.Println(test())
}