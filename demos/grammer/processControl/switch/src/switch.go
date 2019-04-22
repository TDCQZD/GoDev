package utils

import (
	"fmt"
)

func main() {
	/*
		var char byte
		fmt.Println("请输入字符")
		fmt.Scanf("%c",&char)
		utils.Test(char)

		utils.Details()

		utils.SwitchDemo1(char)

		var grade float64
		fmt.Println("请输入学生成绩")
		fmt.Scanf("%f",&grade)
		utils.SwitchDemo2(grade)

		var mouth int
		fmt.Println("请输入月份")
		fmt.Scanf("%d",&mouth)
		utils.SwitchDemo3(mouth)
	*/
	var week string
	fmt.Println("请输入星期时间")
	fmt.Scanf("%s", &week)
	SwitchDemo4(week)

}
func Test(char byte) {
	switch char {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("输入有误，请重新输入")
	}

}
func test(n1 int, n2 int) int {
	return n1 + n2
}
func Details() {
	//1. case后是一个表达式( 即：常量值、变量、一个有返回值的函数, 也可以是计算式等都可以
	var key int = 10
	switch key {
	case test(3, 7):
		fmt.Println("ok~")
	default:
		fmt.Println("default")
	}

	//case后的各个表达式的值的数据类型，必须和 switch 的表达式数据类型一致

	var key2 float64 = 10
	switch key2 {
	case 12.3:
		fmt.Println("ok~")
	default:
		fmt.Println("default")
	}

	//case后面可以带多个表达式，使用逗号间隔。比如 case 表达式1, 表达式2
	var key3 string = "jack"
	switch key3 {
	case "tom", "jack":
		fmt.Println("tom or jack")
	case "smith", "scott":
		fmt.Println("smith or scott")
	default:
		fmt.Println("default")
	}

	//case后面的表达式如果是常量值(字面量)，则要求不能重复

	var age int = 10
	var age2 int = 10
	switch age {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	case age2:
		fmt.Println("10~")
	default:
		fmt.Println("default")
	}

	//5. case后面不需要带break , 程序匹配到一个case后就会执行对应的代码块，然后退出switch，如果一个都匹配不到，则执行 default

	//6. default 语句不是必须

	//7. switch 后也可以不带表达式，类似 if --else分支来使用。【案例演示, 判断 score 范围】

	var score float64 = 89.5
	switch {
	case score > 90 && score <= 100:
		fmt.Println("成绩优秀")
	case score > 70 && score <= 90:
		fmt.Println("成绩优良")
	}

	//8. switch 后也可以直接声明/定义一个变量，分号结束，不推荐

	switch month := 10; month {
	case 10:
		fmt.Println("10月份")
	default:
		fmt.Println("其它月份")
	}

	//9. switch 穿透-fallthrough ，如果在case语句块后增加fallthrough ,
	//则会继续执行下一个case，也叫switch穿透

	var season string = "winter"
	switch season {
	case "winter":
		fmt.Println("冬天")
		fallthrough //穿透， 默认只是穿透一层
	case "spring":
		fmt.Println("春天")
	case "summer":
		fmt.Println("夏天")
	case "autumn":
		fmt.Println("秋天")

	}

	//10 Type Switch：switch 语句还可以被用于
	//type-switch 来判断某个 interface 变量中实际指向的变量类型

	//这个 type-switch 又叫 类型断言，它的作用是可以判断某个接口实际指向的数据类型
	// 可以简单的理解，一个空接口，可以接受任意的数据类型【多态的核心】

	var t interface{} // 这个是一个空接口
	var n int32 = 80  //
	t = n             // 将  n 赋值给 t

	//写一代码，判断 t究竟指向什么类型 type-switch 又叫 类型断言
	// i := t.(type)  有两个作用
	// (1) 先尝试将 t 转成实际指向的数据类型
	// (2) 如果转成功，就会将实际指向的值，赋给 i
	// (3)
	switch i := t.(type) { //
	case float32:
		fmt.Printf("i的类型%T i的值%v\n", i, i)
	case float64:
		fmt.Printf("i的类型%T i的值%v\n", i, i)
	case int32:
		fmt.Printf("i的类型%T i的值%v\n", i, i)
	case bool:
		fmt.Printf("i的类型%T i的值%v\n", i, i)
	case string:
		fmt.Printf("i的类型%T i的值%v\n", i, i)
	default:
		fmt.Println("类型不确定")
	}
}

/*
1) 使用 switch 把小写类型的 char型转为大写(键盘输入)。只转换 a, b, c, d, e. 其它的输出 “other”。
分析：
1、变量 char byte
2、switch 比较 char
3、输出
*/
func SwitchDemo1(char byte) {

	switch char {
	case 'a':
		fmt.Println("A")
		fmt.Printf("%c", char-32)
	case 'b':
		fmt.Println("B")
		fmt.Printf("%c", char-32)
	case 'c':
		fmt.Println("C")
		fmt.Printf("%c", char-32)
	case 'd':
		fmt.Println("D")
		fmt.Printf("%c", char-32)
	case 'e':
		fmt.Println("E")
		fmt.Printf("%c", char-32)
	default:
		fmt.Println("other")
	}
}

/*
2) 对学生成绩大于60分的，输出“合格”。低于60分的，输出“不合格”。(注：输入的成绩不能大于100)
1、变量 grade float64
2、case后使用表达式
2、case后使用int 常量
*/
func SwitchDemo2(grade float64) {

	switch {

	case grade >= 60 && grade <= 100:
		fmt.Println("合格")
	case grade < 60 && grade > 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误")
	}
	//常量的做法,101~119 会出错 % 不能使用在float上
	switch int(grade / 60) {
	case 1:
		fmt.Println("合格")
	case 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误")
	}
}

/*
3) 根据用户指定月份，打印该月份所属的季节。
3,4,5 春季 6,7,8 夏季  9,10,11 秋季 12, 1, 2 冬季

分析：
1、变量 mouth int
2、switch 后面跟多个变量
3、输出
*/
func SwitchDemo3(mouth int) {
	switch mouth {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("输入有误")
	}
}

/*
4) 根据用户输入显示对应的星期时间(string)，
如果“星期一”，显示“干煸豆角”如果“星期二”，显示“醋溜土豆”
如果“星期三”，显示“红烧狮子头”如果“星期四”，显示“油炸花生米”
如果“星期五”，显示“蒜蓉扇贝”如果“星期六”，显示“东北乱炖”如果“星期日”，显示“大盘鸡”
分析：
1、变量 week string
2、switch 后面跟单个个变量
3、输出
*/
func SwitchDemo4(week string) {
	switch week {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("冬油炸花生米季")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")

	default:
		fmt.Println("输入有误")
	}
}
