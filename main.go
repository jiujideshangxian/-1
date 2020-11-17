package main

import "fmt"

func main() {
	arr:=[]int64{10,5,9,6,7}
	fmt.Println("初始值为",arr)
	 res:=0
	 for i:=0;i<len(arr);i++{
		 for j:=0;j<len(arr)-1;j++ {
			 if arr[j]>arr[j+1]{
				 arr[j],arr[j+1]=arr[j+1],arr[j]
			 }
		 }
		 res++
		 fmt.Println("冒泡排序的第",res,"次结果为",arr)
		 fmt.Println("无聊1")
		 fmt.Println("无聊2git")
		 fmt.Println("无聊3")
		 fmt.Println("无聊4")
		 fmt.Println("无聊5")
		 fmt.Println("无聊6")
		 fmt.Println("无聊7")
	 }

}
