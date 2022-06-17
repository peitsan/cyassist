package login

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"time"
)

var Key []byte

func Auth(c *gin.Context) (ok bool) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		return false
	}

	token, err := DecryptLogin(strings.Fields(token)[1])
	if err != nil {
		log.Printf("token认证过程中无法解析传进token,err:%v", err.Error())
		return false
	}

	temp := strings.Split(token, ":")
	if len(temp) != 2 {
		log.Println("token认证过程中解密后无法分离user和timestamp")
		return false
	}
	c.Set("mail", temp[0])
	timestamp, err := strconv.ParseInt(temp[1], 10, 64)
	if err != nil {
		return false
	}
	if time.Now().Unix()-timestamp > 2592300 {
		return false
	}
	return true

}

func EncryptLogin(mail string) (string, error) {
	encrypted, err := EncryptByAes([]byte(mail+":"+strconv.FormatInt(time.Now().Unix(), 10)), Key)
	if err != nil {
		log.Printf("在对登录状态加密时，出现错误error:%v", err)
		return "", err
	} else {
		return string(encrypted), nil
	}
}

func DecryptLogin(content string) (string, error) {
	decrypted, err := DecryptByAes(content, Key)
	if err != nil {
		log.Printf("在对登录状态加密时，出现错误error:%v", err)
		return "", err
	} else {
		return string(decrypted), nil
	}
}

//pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

//AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

//AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

//EncryptByAes Aes加密 后 base64 再加
func EncryptByAes(data, key []byte) (string, error) {
	res, err := AesEncrypt(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

//DecryptByAes Aes 解密
func DecryptByAes(data string, key []byte) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(dataByte, key)
}
