package main

import(
	"fmt"
	"strconv"
)

func main(){
	f := 1000
	var p *int = &f
	fmt.Println(p)

	//if�������
	if f := 10; f > 0 {
		fmt.Println(f)
	}
	fmt.Println(*p)

	//	forѭ�����ĳ����÷�
	sum := 0
	for g := 0; g <= 100; g++ {
		sum += g
	}
	fmt.Println("��һ���÷�:" + strconv.Itoa(sum))

	//forѭ���ĵڶ����÷�������whileѭ��
	for sum < 6000 {
		sum++
	}
	fmt.Println("�ڶ����÷�:" + strconv.Itoa(sum))

	//for ѭ���ĵ������÷�����������do while

	for {
		sum++
		if sum > 6001 {
			break
		}
	}
	fmt.Println("forѭ���ĵ������÷�:" + strconv.Itoa(sum))
}
	