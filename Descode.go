package main

import (
	"bytes"
	"crypto/des"
	"fmt"
) //cryptogram   crypto:加密

func main() {
	/*
	des块加密，分组加密
	块加密不能直接加密，直接加密只会加密一部分
	 des:key,data,mode

	key :密钥
	data：明文，即将加密的明文
	mode：两种模式，加密和解密
	*/

	/*key:=[]byte("c1906041")

	data:="遇贵人先立业，遇良人先成家"

	//cipher 密码

	block,err:=des.NewCipher(key)
	if err != nil {
		panic("初始化密钥错误，请重试！第一处")
	}
	//dst,src
	dst:=make([]byte,len([]byte(data)))
	//加密过程
	block.Encrypt(dst,[]byte(data))
	//fmt.Printf("%x",dst)
	fmt.Println("加密后的内容",dst)
	//解密
	deBlock,err:=des.NewCipher(key)
	if err != nil {
		panic("初始化密钥错误，请重试！第二处")
	}
*/
	key:=[]byte("c1906041")

	data:="遇贵人先立业，遇良人先成家"

	//1,得到cipher
	block,_:=des.NewCipher()

	//2,对数据明文进行结尾块填充

	//3，选择模式
	//NewCBCEncrypter 分组块加密
	fmt.Println("待续。。。")

	//二，对数据进行解密
	//DES三元素：key,data,mode
}

//明文数据尾部填充
func EndPadding(text []byte,blockSize int)[]byte{
	//计算要填充的块内容的大小
	paddingSize :=blockSize- len(text)%blockSize
	//Repeat 平铺，html中图片总是重复，总要no-repeat
	paddingText:=bytes.Repeat()

}
