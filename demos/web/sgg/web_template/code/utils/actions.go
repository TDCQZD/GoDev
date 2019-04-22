package utils

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

/*条件动作*/
func handlerIF(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../view/if.html")
	//声明一个变量
	age := 16
	//执行模板
	t.Execute(w, age > 18)

}

/*迭代动作
1、数组
2、结构体
*/
func handlerRange(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../view/range.html")
	stars := []string{"马蓉", "李小璐", "白百何"}
	//执行模板
	t.Execute(w, stars)

}

/*设置动作*/
func handlerWith(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../view/with.html")

	//执行模板
	t.Execute(w, "狸猫")

}

/*包含动作*/
func handlerTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../view/template.html", "../view/template2.html"))

	t.Execute(w, "Testing Template")

}

/*定义动作
1、在一个模板文件中定义一个模板
2、一个模板文件中定义多个模板
3、在不同的模板文件中定义同名的模板
*/
func handlerDefine(w http.ResponseWriter, r *http.Request) {
	/*单个模板*/
	// t, _ := template.ParseFiles("../view/define.html")

	/*多个模板*/
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(5) > 2 {
		//解析模板文件
		t = template.Must(template.ParseFiles("../view/defineMore.html", "../view/defineMore_Content1.html"))
	} else {
		//解析模板文件
		t = template.Must(template.ParseFiles("../view/defineMore.html", "../view/defineMore_Content2.html"))
	}
	t.ExecuteTemplate(w, "model", "定义模板")
}

/*块动作*/
func handlerBlock(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(5) > 2 {
		//解析模板文件
		t = template.Must(template.ParseFiles("../view/block.html", "../view/define.html"))
	} else {
		//解析模板文件
		t = template.Must(template.ParseFiles("../view/block.html"))
	}
	//执行模板
	t.ExecuteTemplate(w, "model", "")

}

func ActionsDemo() {
	http.HandleFunc("/if", handlerIF)
	http.HandleFunc("/range", handlerRange)
	http.HandleFunc("/with", handlerWith)
	http.HandleFunc("/template", handlerTemplate)
	http.HandleFunc("/define", handlerDefine)
	http.HandleFunc("/block", handlerBlock)
	http.ListenAndServe(":8080", nil)
}
