// 反射：对变量动态分析
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}
func (self Student) Getname(){
	fmt.Println(self.Name)
}
func (self Student) Getage(){
	fmt.Println(self.Age)
}
func practice1(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("value: ", v)
	//switch v.Kind() {
	//case reflect.Float64:
	//	fmt.Println("type: ", v.Type())
	//	fmt.Println("kind: ", v.Kind())
	//	fmt.Println("value:", v.Float())
	//default:
	//	fmt.Println(123)
	//}
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())
	fmt.Println("value:", v.Float())

	fmt.Println(v.Interface()) // 转换成interface
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
func practice2(a interface{}){

	fv := reflect.ValueOf(a)
	fv.Elem().SetFloat(3.3)
	// 传地址 改值
	//fmt.Println(a) // reflect.Value.SetFloat using unaddressable value
}
func practice3(a interface{}){
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct{
		fmt.Println("exepect struct")
		return
	}
	num := val.NumField()
	numofmethod := val.NumMethod()
	fmt.Printf("NumField %d, NumMethod %d", num, numofmethod)
}
func practice4(t interface{}){
	// 通过反射操作结构体
	//val := reflect.ValueOf(s).Elem()
	//fmt.Println(123, val)
	//
	//for i:=0;i<val.NumField();i++{
	//	//reflect.TypeOf()
	//	fmt.Printf("%v\n",val.Field(i))
	//}
	//val.Field(0).Elem().SetString("kobe")
	//val.Field(1).Elem().SetInt(41)
	s := reflect.ValueOf(t).Elem()
	fmt.Println(123, s)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetString("kobe")
	s.Field(1).SetInt(41)

}
func main() {
	//var x float64 = 3.4
	//practice1(x)

	s := Student{
		Name:  "tushala",
		Age:   22,
		Score: 99.9,
	}
	//v := reflect.ValueOf(s)
	//fmt.Println(v.Kind())
	//fmt.Println(v.Type())
	//iv := v.Interface()
	//stu, ok := iv.(Student)
	//fmt.Println(stu, ok)*/
	//var a float64 = 2
	//fmt.Println(a)
	//practice2(&a)
	//fmt.Println(a)
	//practice3(s)
	fmt.Println(s)
	practice4(&s)
	fmt.Println(s)

}
