package common

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	crand "crypto/rand"
	"fmt"
)

const (
	VerifyKey     = "goobeer1goobeer2goobeer3"
	VerifyFormKey = "gvk"
)

var (
	RandSeed = "abcdefghjkmnpqrstuvwxyz23456789"
)

//生成随机字符串:strLen:随机字符串中字符长度,numLen:数字个数,upper:区分大小写,默认小写
func CreateRandom(strLen, numLen int, upper ...bool) (randStr string, err error) {
	b := make([]byte, strLen)
	n := make([]byte, numLen)

	_, err = crand.Read(b)
	if err != nil {
		return
	}
	_, err = crand.Read(n)
	if err != nil {
		return
	}
	ciFlag := byte('a')
	if len(upper) > 0 {
		if upper[0] {
			ciFlag = byte('A')
		}
	}

	for i := range b {
		b[i] = ciFlag + b[i]%26
	}
	niFlag := byte('0')
	for i := range n {
		t := niFlag + n[i]%10
		b = append(b, t)
	}

	allNum := strLen + numLen
	kvData := make(map[int]byte, allNum)
	allNum--

	for ; allNum >= 0; allNum-- {
		kvData[allNum] = 0
	}
	allNum = 0
	for k, _ := range kvData {
		//swap b[k] && b[allNum]
		t := b[k]
		b[k] = b[allNum]
		b[allNum] = t
		allNum++
	}
	randStr = string(b)
	return
}

//从种子字符串中生成随机数
func CreateRandomFromSeed(strLen int, randSeed string) (randStr string, err error) {
	b := make([]byte, strLen)
	_, err = crand.Read(b)
	if err != nil {
		return
	}
	randSeedBytes := []byte(randSeed)
	randstrLen := byte(len(randSeed))
	for i := range b {
		b[i] %= randstrLen
		b[i] = randSeedBytes[b[i]]
	}
	randStr = string(b)
	return
}

//生成guid
func CreateGuid() (guid string, err error) {
	b := make([]byte, 16)
	_, err = crand.Read(b)
	if err != nil {
		return
	}
	guid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

func TripleDesEncrypt(data, key []byte) ([]byte, error) {
	if key == nil {
		key = []byte(VerifyKey)
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	data = PKCS5Padding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	return crypted, nil
}

func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	if key == nil {
		key = []byte(VerifyKey)
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
