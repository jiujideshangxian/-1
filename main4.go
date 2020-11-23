package main

import "fmt"

func main() {
	var i int
	for ( i=2;i<=50;i++) {
		var j = 2;
		while(i%j!=0){
			j++
		}
		fmt.Println()
	}
}
