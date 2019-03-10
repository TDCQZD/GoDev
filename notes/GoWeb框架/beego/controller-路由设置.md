# 路由设置
 beego 存在三种方式的路由:固定路由、正则路由、自动路由，接下来详细的讲解如何使用这三种路由。
 ## 基础路由
从 beego 1.2 版本开始支持了基本的 RESTful 函数式路由,应用中的大多数路由都会定义在 routers/router.go 文件中。最简单的 beego 路由由 URI 和闭包函数组成

### 基本 GET 路由
```
beego.Get("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world"))
})
```
### 基本 POST 路由
```
beego.Post("/alice",func(ctx *context.Context){
     ctx.Output.Body([]byte("bob"))
})
```
### 注册一个可以响应任何 HTTP 的路由
```
beego.Any("/foo",func(ctx *context.Context){
     ctx.Output.Body([]byte("bar"))
})
```
### 所有的支持的基础函数如下所示
```
beego.Get(router, beego.FilterFunc)
beego.Post(router, beego.FilterFunc)
beego.Put(router, beego.FilterFunc)
beego.Patch(router, beego.FilterFunc)
beego.Head(router, beego.FilterFunc)
beego.Options(router, beego.FilterFunc)
beego.Delete(router, beego.FilterFunc)
beego.Any(router, beego.FilterFunc)
```
### 支持自定义的 handler 实现
有些时候我们已经实现了一些 rpc 的应用,但是想要集成到 beego 中,或者其他的 httpserver 应用,集成到 beego 中来.现在可以很方便的集成:
```
s := rpc.NewServer()
s.RegisterCodec(json.NewCodec(), "application/json")
s.RegisterService(new(HelloService), "")
beego.Handler("/rpc", s)
```
beego.Handler(router, http.Handler) 这个函数是关键,第一个参数表示路由 URI, 第二个就是你自己实现的 http.Handler, 注册之后就会把所有 rpc 作为前缀的请求分发到 http.Handler 中进行处理.

这个函数其实还有第三个参数就是是否是前缀匹配,默认是 false, 如果设置了 true, 那么就会在路由匹配的时候前缀匹配,即 /rpc/user 这样的也会匹配去运行


 ## RESTful Controller 路由
 RESTful 是一种目前 API 开发中广泛采用的形式，beego 默认就是支持这样的请求方法，也就是用户 Get 请求就执行 Get 方法，Post 请求就执行 Post 方法。因此默认的路由是这样 RESTful 的请求方式。
 ### 固定路由
 固定路由也就是全匹配的路由，如下所示：
```
beego.Router("/", &controllers.MainController{})
beego.Router("/admin", &admin.UserController{})
beego.Router("/admin/index", &admin.ArticleController{})
beego.Router("/admin/addpkg", &admin.AddController{})
```

 ### 正则路由
    为了用户更加方便的路由设置，beego 参考了 sinatra 的路由实现，支持多种方式的路由：

- beego.Router(“/api/?:id”, &controllers.RController{})

    默认匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”

- beego.Router(“/api/:id”, &controllers.RController{})

    默认匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”，但URL”/api/“匹配失败

- beego.Router(“/api/:id([0-9]+)“, &controllers.RController{})

    自定义正则匹配 //例如对于URL”/api/123”可以匹配成功，此时变量”:id”值为”123”

- beego.Router(“/user/:username([\\w]+)“, &controllers.RController{})

    正则字符串匹配 //例如对于URL”/user/astaxie”可以匹配成功，此时变量”:username”值为”astaxie”

- beego.Router(“/download/*.*”, &controllers.RController{})

    *匹配方式 //例如对于URL”/download/file/api.xml”可以匹配成功，此时变量”:path”值为”file/api”， “:ext”值为”xml”

- beego.Router(“/download/ceshi/*“, &controllers.RController{})

    *全匹配方式 //例如对于URL”/download/ceshi/file/api.json”可以匹配成功，此时变量”:splat”值为”file/api.json”

- beego.Router(“/:id:int”, &controllers.RController{})

    int 类型设置方式，匹配 :id为int 类型，框架帮你实现了正则 ([0-9]+)

- beego.Router(“/:hi:string”, &controllers.RController{})

    string 类型设置方式，匹配 :hi 为 string 类型。框架帮你实现了正则 ([\w]+)

- beego.Router(“/cms_:id([0-9]+).html”, &controllers.CmsController{})

    带有前缀的自定义正则 //匹配 :id 为正则类型。匹配 cms_123.html 这样的 url :id = 123

可以在 Controller 中通过如下方式获取上面的变量：
```
this.Ctx.Input.Param(":id")
this.Ctx.Input.Param(":username")
this.Ctx.Input.Param(":splat")
this.Ctx.Input.Param(":path")
this.Ctx.Input.Param(":ext")
```
 ###  
 [路由设置](https://beego.me/docs/mvc/controller/router.md)