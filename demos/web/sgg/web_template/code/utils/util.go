package utils

import (
	"html/template"
	"net/http"
)

/*
使用Go的Web模板引擎需要以下两个步骤:
1、解析模板文件(对文本格式的模板源进行语法分析)
2、执行模板(执行经过语法分析的模板)
*/

//创建单个模板处理器函数
func handleSingleTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../view/index.html") //解析模板文件
	t.Execute(w, "Using Single Template")             //执行
}

//创建多个模板处理器函数
func handleMoreTemplate(w http.ResponseWriter, r *http.Request) {
	//通过Must函数让Go帮我们自动处理异常
	t := template.Must(template.ParseFiles("../view/index.html", "../view/index2.html"))
	//将响应数据在文件中显示
	// t.Execute(w, "Using More Template")
	t.ExecuteTemplate(w, "index2.html", "Using More Template")
	/*
		注意事项：
		t.ExecuteTemplat 中name 只是文件名，不需要路径
	*/

}

func TemplateDemo() {
	http.HandleFunc("/", handleMoreTemplate)
	http.ListenAndServe(":8080", nil)
}
