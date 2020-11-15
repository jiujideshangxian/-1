package main

import "fmt"

func main() {
	//输出1~50的素数
	var A,a int
	A=1

	LOOP:
		for A < 50 {
               A++
			for a = 2;a<A;a++ {
				if A%a==0 {
					goto LOOP
				}
			}
			fmt.Printf(" %d \t",  A)
		}

}
