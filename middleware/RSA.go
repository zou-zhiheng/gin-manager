package middleware

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

func GeneralRSAKey() ([]byte, []byte) {
	// 指定生成密钥的长度（单位：位）
	keySize := 2048

	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		log.Fatal(err)
	}

	// 将私钥转换为PEM格式
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	//// 将私钥写入文件
	//privateKeyFile, err := os.Create("private.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer privateKeyFile.Close()
	//err = pem.Encode(privateKeyFile, privateKeyPEM)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(privateKeyPEM)

	fmt.Println("Private key generated and saved to private.pem")

	// 获取公钥
	publicKey := &privateKey.PublicKey

	// 将公钥转换为PEM格式
	publicKeyPEM, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("public")
	fmt.Println(publicKeyPEM)

	// 将公钥写入文件
	//publicKeyFile, err := os.Create("public.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer publicKeyFile.Close()
	//err = pem.Encode(publicKeyFile, &pem.Block{
	//	Type:  "PUBLIC KEY",
	//	Bytes: publicKeyPEM,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("Public key generated and saved to public.pem")
	return privateKeyPEM.Bytes, publicKeyPEM
}
