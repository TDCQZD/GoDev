package test
import (
	"reflect"
	"testing"
)
/*反射的最佳实践*/
// 5、使用反射操作任意结构体类型：【了解】
type user struct {
    UserId string
    Name string
    
} 

func TestReflectStruct(t *testing.T) {
    var (
        model *user
        sv    reflect.Value
    )
    model = &user{}
    sv = reflect.ValueOf(model)
    t.Log("reflect.ValueOf", sv.Kind().String())
    sv = sv.Elem()
    t.Log("reflect.ValueOf.Elem", sv.Kind().String())
    sv.FieldByName("UserId").SetString("12345678")
    sv.FieldByName("Name").SetString("nickname")
    t.Log("model", model)
}
