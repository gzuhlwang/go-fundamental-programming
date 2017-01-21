package main

import(
	"fmt"
	"strconv"
)

func main(){
	f := 1000
	var p *int = &f
	fmt.Println(p)

	//if控制语句
	if f := 10; f > 0 {
		fmt.Println(f)
	}
	fmt.Println(*p)

	//	for循环语句的常用用法
	sum := 0
	for g := 0; g <= 100; g++ {
		sum += g
	}
	fmt.Println("第一种用法:" + strconv.Itoa(sum))

	//for循环的第二种用法类似于while循环
	for sum < 6000 {
		sum++
	}
	fmt.Println("第二种用法:" + strconv.Itoa(sum))

	//for 循环的第三种用法借助类似于do while

	for {
		sum++
		if sum > 6001 {
			break
		}
	}
	fmt.Println("for循环的第三种用法:" + strconv.Itoa(sum))
}
	