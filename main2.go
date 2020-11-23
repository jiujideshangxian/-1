package main

import "fmt"

func main() {
	var  a,b int
	for a = 2;a<=50;a++{
		for b = 2;b<=(a/b);b++{
			if a%b==0 {
				break
			}
			//fmt.Printf("%d\t", b)
		}
		//fmt.Printf("%d\t", a)

		fmt.Printf("%d\t", b)

		/*if b > (a/b) {
			fmt.Printf("%d\t", a)
		}*/
		//减少循环次数
		//eg:24   1,2,3,4,6,8,12,24
		//1*24   2*12  3*8  4*6
		//
		/*
		if b < (a/b) {
			fmt.Printf("%d\t", a)
		}*/
	}
}
