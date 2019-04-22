package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"log"
	"os"
)

// 判断一个数是否是素数？
// 即只能被1或者自身整除的自然数（不包括1），称为素数/质数。
func numberIsPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

//错误检查-普通
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//错误检查-自定义
func CheckErrorNew(err error, errStr string) {
	if err != nil {
		log.Fatal(err, errors.New(errStr))

	}
}

//panic错误检查-普通
func CheckPanicError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

//panic错误检查-自定义
func CheckPanicErrorNew(err error, errStr string) {
	if err != nil {
		log.Fatal(err, errors.New(errStr))
		panic(err)
	}
}

//判断字符串是否为空
func IsNullString(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	} else {
		return false
	}
}

//序列化数据
func MarshalStruct(v interface{}) (data []byte, err error) {
	data, err = json.Marshal(v)
	return
}

//反序列化数据
func UnmarshalStruct(b []byte, v interface{}) (data interface{}, err error) {
	err = json.Unmarshal(b, v)
	data = v
	return
}

//int转[]byte
func IntToByte(data int) [4]byte {
	datas := uint32(data)
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4], datas)
	return bytes
}

// 文件是否存在
func FileExist(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return true
	}
	return false
}
// byte反转
func ByteSliceReverse(s []byte) {
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}
	return
}
// 结构体序列化为字符串
func (addr *Address) String() string {
	return fmt.Sprintf("%s://%s:%d", addr.Transport, addr.Host, addr.Port)
}
