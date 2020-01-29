/*实现一个负载均衡调度算法，支持随机、轮训等算法*/
package main

import (
	"fmt"
	"go-learning/step2/5-反射/homework/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i:=0;i<16;i++{
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	var balanceName = "roundrobin"
	if len(os.Args) > 1{
		balanceName = os.Args[1]
	}
	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil{
			fmt.Fprintf(os.Stdout, "do balance err\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}