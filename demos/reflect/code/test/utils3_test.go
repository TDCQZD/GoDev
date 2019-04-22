package test
import (
	"testing"
	 "reflect"       
)
/*反射的最佳实践*/


/*
3) 定义了两个函数test1和test2，定义一个适配器函数用作统一处理接口【了解】：
(1) 定义了两个函数
test1 := func(v1 int, v2 int) {
	t.Log(v1, v2)
}
test2 := func(v1 int, v2 int, s string) {
	t.Log(v1, v2, s)
}
(2) 定义一个适配器函数用作统一处理接口, 其大致结构如下:
bridge := func(call interface{}, args ...interface{}) {
 //内容       
}
//实现调用test1对应的函数
bridge(test1, 1, 2) 
//实现调用test2对应的函数
bridge(test2, 1, 2, "test2")

(3) 要求使用反射机制完成 
*/	


func TestReflectFunc(t *testing.T) {
	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}
	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}
	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}
