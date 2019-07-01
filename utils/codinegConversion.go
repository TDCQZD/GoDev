package utils

import (
	"bytes"
	"io/ioutil"

	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

/*解决中文乱码
src：字符串
srcCode：字符串当前编码
tagCode：要转换的编码
*/
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

/*使用axgle/mahonia库进行编码转换
编码格式："gbk","UTF-8","GB18030"
*/
// string 字符编码的转换
func StringGBKToUTF8(format, contents string) string {

	return mahonia.NewDecoder(format).ConvertString(contents)
}

/*byte 字节编码转换*/
func ByteGbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func ByteUtf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
