package utils

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"bytes"
	"io"
	"crypto/rand"
	"encoding/hex"
)

const (
	key = "909dfef0d4242e7ed94abcc2be9cc747"
	iv  = "37cf1d51a2192ad5"
)

func main() {
	str := "我勒个去"
	es, _ := AesEncrypt(str, []byte(key))
	fmt.Println(es)

	ds, _ := AesDecrypt(es, []byte(key))
	fmt.Println(string(ds))

}

func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	fmt.Println("encode byte", blockSize)
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)
	//补全16位
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	//向量
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AECCBCEncrypter(encoding string) {
	key := []byte(key) //秘钥长度需要时AES-128(16bytes)或者AES-256(32bytes)
	iv := []byte(iv)
	plaintext := []byte(encoding) //原文必须填充至blocksize的整数倍，填充方法可以参见https://tools.ietf.org/html/rfc5246#section-6.2.3.2

	if len(plaintext)%aes.BlockSize != 0 { //块大小在aes.BlockSize中定义
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key) //生成加密用的block
	if err != nil {
		panic(err)
	}

	// 对IV有随机性要求，但没有保密性要求，所以常见的做法是将IV包含在加密文本当中
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	//随机一个block大小作为IV
	//采用不同的IV时相同的秘钥将会产生不同的密文，可以理解为一次加密的session

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 谨记密文需要认证(i.e. by using crypto/hmac)

	fmt.Printf("%x\n", ciphertext)
}

func AECCBCDecrypter(deconding string) string {
	key := []byte(key)
	ciphertext, _ := hex.DecodeString(deconding)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := []byte(iv)
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks可以原地更新
	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("%s\n", ciphertext)
	return string(ciphertext)
	// Output: exampleplaintext
}
