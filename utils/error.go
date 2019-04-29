package utils

import "log"

//错误检查
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
