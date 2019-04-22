package test

import (
	"testing"
	"reflect"
)

/*反射的最佳实践*/
// 6、使用反射创建并操作结构体
type User struct {
    UserId string
    Name string
} 
func TestReflectStructPtr(t *testing.T) {
    var (
        model *User 
        st    reflect.Type
        elem  reflect.Value
    )
    st = reflect.TypeOf(model) //获取类型 *user
    t.Log("reflect.TypeOf", st.Kind().String()) // ptr
    st = st.Elem() //st指向的类型
    t.Log("reflect.TypeOf.Elem", st.Kind().String()) //struct
    elem = reflect.New(st) //New返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针
    t.Log("reflect.New", elem.Kind().String()) // ptr
    t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct
    //model就是创建的user结构体变量(实例)
    model = elem.Interface().(*User) //model 是 *user 它的指向和elem是一样的.
    elem = elem.Elem() //取得elem指向的值
    elem.FieldByName("UserId").SetString("12345678") //赋值..
    elem.FieldByName("Name").SetString("nickname")
    t.Log("model model.Name", model, model.Name)
}
