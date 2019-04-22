package utils

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func ErrFormat(format string, args ...interface{}) {
	logs.Error(fmt.Sprintf(format, args...))
}

func ErrLog(comment string, uid int64, err error) string {
	return fmt.Sprintf("%v uid = %v | err = %v", comment, uid, err)
}

func FmtLog(comment string) {
	fmt.Println(comment)
}
