package main

import "fmt"

func main() {
	//输出1~50的素数
	var A,a int
	A = 1 //不写人for循环是因为会将A变为1
	LOOP:
		for A < 50 {
               A++
			for a = 2;a<A;a++ {
				if A%a==0 {
					goto LOOP//发现因子则不是素数
				}
			}
			fmt.Printf(" %d \t",  A)
		}

}
