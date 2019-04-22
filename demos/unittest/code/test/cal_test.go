package test

import (
	"testing" //引入go 的testing框架包
)
//编写测试用例，测试getSum是否正确
func TestSum(t *testing.T)  {
	res := getSum(10)
	if res == 55 {
		t.Logf("测试成功")
	}else{
		t.Fatalf("测试失败")
	}
}