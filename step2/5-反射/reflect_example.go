package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	S1, S2, S3 string
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func (self NotknownType) String() string {
	// 1
	return self.S1 + "-" + self.S2 + "-" + self.S3
}
func (self NotknownType) NewString(s1, s2 string) string {
	// 0
	return self.S1 + "-" + self.S2 + "-" + self.S3 + "\n" + s1 + "\n" + s2
}
func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)

	kind := value.Kind()
	fmt.Println(kind)
	fmt.Println(value.NumField(), value.NumMethod())
	for i:=0;i<value.NumField();i++{
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	results1 := value.Method(1).Call(nil)
	fmt.Println(results1)
	param := make([]reflect.Value, 2)
	param[0] = reflect.ValueOf("kobe")
	param[1] = reflect.ValueOf("bryant")
	results2 := value.Method(0).Call(param)
	fmt.Println(results2)

}
