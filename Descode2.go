package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)


func main() {
	key:=[]byte("00000000")//密钥只占8个字节
	arr:="千锋教育"
	fmt.Println("------------DES加密字节数组")
	fmt.Println("加密前",arr,key)
	//resultArr,_:=DesEncrypt()
}

//DES解密字节数，返回字节数组
func DesDecrypt(cipherBytes, key []byte)([]byte ,error){
	block,err:=des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode:=cipher.NewCBCDecrypter(block,key)
	originalText:=make([]byte,len(cipherBytes))
	blockMode.CryptBlocks(originalText,cipherBytes)
	originalText=PKCS5UnPadding(originalText)
	return originalText,nil
}

//DES加密文本返回加密后文本
func DesEmcryptString(originalText string,key []byte)(string ,error){
	cipherArr:=DesEncrypt([]byte(originalText),key)

}

//尾部填充
func PKCS5UnPadding(ciphertext []byte,blockSize int)[]byte{
	padding:=blockSize - len(ciphertext)% blockSize
	padtext:=bytes.Repeat([]byte{byte(padding)},padding)
	return append(ciphertext,padtext...)
}

func PKCSS5UnPadding(origData []byte)[]byte{
	length:=len(origData)
	//去掉最后一个字节unpadding次
	unpadding:=int(origData[length-1])
	return origData[:(length-unpadding)]

}
