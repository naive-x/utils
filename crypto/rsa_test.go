package cryptoutil

import (
	// "bytes"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"

	// "crypto/x509"
	// "encoding/pem"
	// "fmt"
	"log"
	"testing"
)

func TestD(t *testing.T) {
	bs := "Fd1zA0kja6uFEke+AZhlsMbLfmoOOLlD5xCUau9FV8bedebGj688KY7Tct1A5gUEmGmlxngeJFpLKoZ/WrIYQx20CI8orqP/akrGlU5Ze8nDPd76N3LKC24u0ty1k1mUFrt5P8605w3PP4w9UGQFp8GcgRp1GBrBEkHBrX4R6xy6g9ox3QEwnB0UUbhOCDlbpon4OqQyjVmqJvWkUuTBXalrnu9AvzpDMvCj107tx/4wZgoCz9DmucBLHqm7YBMuzE58nfZn1ILLnb8MXVGdrjM74aDm6f7sWqfGE286XpTQ/jAVWy5rhH99GGwwobfOgPeXJVZjSEB8bVqCbs+RgA=="

	enc, _ := base64.StdEncoding.DecodeString(bs)

	privPem, _ := ioutil.ReadFile("priv.pem")
	blockPriv, _ := pem.Decode(privPem)
	privKey, _ := x509.ParsePKCS1PrivateKey(blockPriv.Bytes)
	dec, _ := rsa.DecryptPKCS1v15(rand.Reader, privKey, enc)
	log.Println("dec: ", string(dec))
}
func TestC(t *testing.T) {
	pubPem, _ := ioutil.ReadFile("pub.pem")

	block, _ := pem.Decode(pubPem)

	//

	// 	ParsePKIXPublicKey parses a public key in PKIX, ASN.1 DER form. The encoded public key is a SubjectPublicKeyInfo structure (see RFC 5280, Section 4.1).

	// It returns a *[rsa.PublicKey], *[dsa.PublicKey], *[ecdsa.PublicKey], [ed25519.PublicKey] (not a pointer), or *[ecdh.PublicKey] (for X25519). More types might be supported in the future.

	// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".

	pubKey, _ := x509.ParsePKIXPublicKey(block.Bytes)

	originalData := "hello world"
	enc, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	log.Println("enc:", string(enc))

	privPem, _ := ioutil.ReadFile("priv.pem")
	blockPriv, _ := pem.Decode(privPem)
	privKey, _ := x509.ParsePKCS1PrivateKey(blockPriv.Bytes)
	dec, _ := rsa.DecryptPKCS1v15(rand.Reader, privKey, enc)
	log.Println("dec: ", string(dec))
}

func TestB(t *testing.T) {
	privPem, _ := ioutil.ReadFile("rsa_1024_priv.pem")
	log.Println(string(privPem))
}

func TestA(t *testing.T) {
	// rsa.GenerateKey(rand.Reader, bitSize)
	var priv crypto.PrivateKey // 声明为接口类型
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln("生成站点 rsa 私钥失败")
	}

	derBytes := x509.MarshalPKCS1PrivateKey(priv.(*rsa.PrivateKey))

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derBytes,
	}

	privKey := pem.EncodeToMemory(block)
	log.Println(string(privKey))

	// priv.Public()

	// Extract public key from private key
	// publicKey := &priv.PublicKey

	// fmt.Println("xx: ", priv.PublicKey)

	// // x509.CreateCertificate
	// // 1.2
	pub := priv.(crypto.Decrypter).Public()

	// 断言 pub 是 rsa 类型的公钥
	pubKey := pub.(*rsa.PublicKey)

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	pem.Encode(&buffer, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	log.Println(string(buffer.Bytes()))

	// var buffer bytes.Buffer
	// // for _, cert := range certs {
	// // 	pem.Encode(&buffer, &pem.Block{
	// // 		Type:  "CERTIFICATE",
	// // 		Bytes: cert.Raw,
	// // 	})
	// // }

	// return buffer.Bytes()
}
