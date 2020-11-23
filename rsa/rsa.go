package rsa

import (
	"crypto"
	"crypto/rsa"
	"flag"
	"math/rand"
)

/*

汉森堡
*/
func CratePairKeys([]byte,[]byte){
	//1,先生成私钥
	var bits int
	flag.IntVar(&bits,"b",2048,"密码长度")
	rsa.GenerateKey()
	//2,根据私钥生成公钥

	//3,将私钥和公钥进行返回

}

func RSAEncrypt(key crypto.PublicKey,data []byte)[]byte  {
	rsa.EncryptPKCS1v15(rand.Reader,)
}

/*使用RSA算法对密文数据进行解密，返回加密后的明文*/
funcRSADecrypt(private *rsa,PrivateKey,chiper []byte)([]byte,error){
	return rsa.DecyptPKCS1V15(rand.Reader,Pri)

}