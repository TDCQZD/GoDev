
# beego 的 MVC 架构
beego 是一个典型的 MVC 框架，它的整个执行逻辑如下图所示：

![](./detail.png)

1. 在监听的端口接收数据，默认监听在 8080 端口。
2. 用户请求到达 8080 端口之后进入 beego 的处理逻辑。
3. 初始化 Context 对象，根据请求判断是否为 WebSocket 请求，如果是的话设置 Input，同时判断请求的方法是否在标准请求方法中（GET、POST、PUT、DELETE、PATCH、OPTIONS、HEAD），防止用户的恶意伪造请求攻击造成不必要的影响。
4. 执行 BeforeRouter 过滤器，当然在 beego 里面有开关设置。如果用户设置了过滤器，那么该开关打开，这样可以提高在没有开启过滤器的情况下提高执行效率。如果在执行过滤器过程中，responseWriter 已经有数据输出了，那么就提前结束该请求，直接跳转到监控判断。
5. 开始执行静态文件的处理，查看用户的请求 URL 是否和注册在静态文件处理 StaticDir 中的 prefix 是否匹配。如果匹配的话，采用 http 包中默认的 ServeFile 来处理静态文件。
6. 如果不是静态文件开始初始化 session 模块(如果开启 session 的话)，这个里面大家需要注意，如果你的 BeforeRouter 过滤器用到了 session 就会报错，你应该把它加入到 AfterStatic 过滤器中。
7. 开始执行 AfterStatic 过滤器，如果在执行过滤器过程中，responseWriter 已经有数据输出了，那么就提前结束该请求，直接跳转到监控判断。
8. 执行过过滤器之后，开始从固定的路由规则中查找和请求 URL 相匹配的对象。这个匹配是全匹配规则，即如果用户请求的 URL 是 /hello/world，那么固定规则中 /hello 是不会匹配的，只有完全匹配才算匹配。如果匹配的话就进入逻辑执行，如果不匹配进入下一环节的正则匹配。
9. 正则匹配是进行正则的全匹配，这个正则是按照用户添加 beego 路由顺序来进行匹配的，也就是说，如果你在添加路由的时候你的顺序影响你的匹配。和固定匹配一样，如果匹配的话就进行逻辑执行，如果不匹配进入 Auto 匹配。
10. 如果用户注册了 AutoRouter，那么会通过 controller/method 这样的方式去查找对应的 Controller 和他内置的方法，如果找到就开始执行逻辑，如果找不到就跳转到监控判断。
11. 如果找到 Controller 的话，那么就开始执行逻辑，首先执行 BeforeExec 过滤器，如果在执行过滤器过程中，responseWriter 已经有数据输出了，那么就提前结束该请求，直接跳转到监控判断。
12. Controller 开始执行 Init 函数，初始化基本的一些信息，这个函数一般都是 beego.Controller 的初始化，不建议用户继承的时候修改该函数。
13. 是否开启了 XSRF，开启的话就调用 Controller 的 XsrfToken，然后如果是 POST 请求就调用 CheckXsrfCookie 方法。
14. 继续执行 Controller 的 Prepare 函数，这个函数一般是预留给用户的，用来做 Controller 里面的一些参数初始化之类的工作。如果在初始化中 responseWriter 有输出，那么就直接进入 Finish 函数逻辑。
15. 如果没有输出的话，那么根据用户注册的方法执行相应的逻辑，如果用户没有注册，那么就调用 http.Method 对应的方法（Get/Post 等）。执行相应的逻辑，例如数据读取，数据赋值，模板显示之类的，或者直接输出 JSON 或者 XML。
16. 如果 responseWriter 没有输出，那么就调用 Render 函数进行模板输出。
17. 执行 Controller 的 Finish 函数，这个函数是预留给用户用来重写的，用于释放一些资源。释放在 Init 中初始化的信息数据。
18. 执行 AfterExec 过滤器，如果有输出的话就跳转到监控判断逻辑。
19. 执行 Controller 的 Destructor，用于释放 Init 中初始化的一些数据。
20. 如果这一路执行下来都没有找到路由，那么会调用 404 显示找不到该页面。
21. 最后所有的逻辑都汇聚到了监控判断，如果用户开启了监控模块（默认是开启一个 8088 端口用于进程内监控），这样就会把访问的请求链接扔给监控程序去记录当前访问的 QPS，对应的链接访问的执行时间，请求链接等。





## controller 设计
### 参数配置
### 路由设置
### xsrf 过滤
跨站请求伪造(Cross-site request forgery)， 简称为 XSRF，是 Web 应用中常见的一个安全问题。前面的链接也详细讲述了 XSRF 攻击的实现方式。

当前防范 XSRF 的一种通用的方法，是对每一个用户都记录一个无法预知的 cookie 数据，然后要求所有提交的请求（POST/PUT/DELETE）中都必须带有这个 cookie 数据。如果此数据不匹配 ，那么这个请求就可能是被伪造的。

beego 有内建的 XSRF 的防范机制，要使用此机制，你需要在应用配置文件中加上 enablexsrf 设定：
```
enablexsrf = true
xsrfkey = 61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o
xsrfexpire = 3600
```
或者直接在 main 入口处这样设置：
```
beego.EnableXSRF = true
beego.XSRFKEY = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
beego.XSRFExpire = 3600  //过期时间，默认1小时
```
如果开启了 XSRF，那么 beego 的 Web 应用将对所有用户设置一个 _xsrf 的 cookie 值（默认过期 1 小时），如果 POST PUT DELET 请求中没有这个 cookie 值，那么这个请求会被直接拒绝。如果你开启了这个机制，那么在所有被提交的表单中，你都需要加上一个域来提供这个值。你可以通过在模板中使用 专门的函数 XSRFFormHTML() 来做到这一点：

过期时间上面我们设置了全局的过期时间 beego.XSRFExpire，但是有些时候我们也可以在控制器中修改这个过期时间，专门针对某一类处理逻辑：
```
func (this *HomeController) Get(){
    this.XSRFExpire = 7200
    this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
}
```
### session 控制
### flash 数据
### 请求数据处理
### 多种格式数据输出
### 表单数据验证
### 模板输出
### 模板函数
### 错误处理
### 静态文件处理
### 参数配置
###日志处理

## model 设计

## view 设计