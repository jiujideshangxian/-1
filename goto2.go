package main

import "fmt"

func main() {
	 c:=hanshu1(50,2)
	 fmt.Println(c)
}

func hanshu1(A,a int )[]int{
c:=make([]int,50)
LOOP:
	for A < 50 {
		A++
		for a = 2;a<A;a++ {
			if A%a==0 {
				goto LOOP
			}
		}
		c=append(c,A)

	}
	return  c


}
