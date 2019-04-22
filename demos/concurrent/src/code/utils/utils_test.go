package utils

import (
	"testing"
)

func TestPrimeNumber(t *testing.T)  {
	// res :=PrimeNumber()
	res :=PrimeGoroutine()
	if res {
		t.Logf("测试成功")
	}

}
