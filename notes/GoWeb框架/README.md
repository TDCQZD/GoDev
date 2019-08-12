# GoWeb框架
* Gin: 一个Go语言写的HTTP Web框架。它提供了Martini风格的API并有更好的性能。 
    * https://github.com/gin-gonic/gin 
    * https://gin-gonic.github.io/gin

* Beego: 一个Go语言下开源的，高性能Web框架 
    * https://github.com/astaxie/beego 
    * https://beego.me

* Buffalo: 一个Go语言下快速Web开发框架 
    * https://github.com/gobuffalo/buffalo 
    * https://gobuffalo.io

* Echo: 一个高性能，极简的Web框架 
    * https://github.com/labstack/echo 
    * https://echo.labstack.com


* Iris: 目前发展最快的Go Web框架。提供完整的MVC功能并且面向未来。 
    * https://github.com/kataras/iris 
    * https://iris-go.com

* Revel: 一个高生产率，全栈Go语言的Web框架。 
    * https://github.com/revel/revel 
    * https://revel.github.io

## beego&gin
- 对mvc的支持
    * beego支持完整的mvc
    * gin不支持完整的mvc，不支持session
- 对路由的支持
    * Beego 支持正则路由和支持restful Controller路由
    * Gin   不支持正则路由
- 适用场景
    * Beego在业务方面较Gin支持更多
        - 在业务更加复杂的项目
        - 在需要快速开发的项目
    * Gin在性能方面较beego更好
        - 如果项目的规模不大，业务相对简单，适用Gin
        - 当某个接口性能遭到较大的挑战，考虑用Gin重写
  