package utils

import (
	"errors"
)
// 定义好我们需要使用到的error 
var (
	FILE_NOT_EXIST = errors.New("file not exist")
	FILE_IO_ERROR = errors.New("file io errr")
)