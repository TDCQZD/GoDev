###项目说明
项目功能
        1、实现用户管理功能：验证、数据库操作
        2、实现图书管理功能：book 增删改查
        3、实现购物车管理
        4、实现订单管理
        
###用户管理功能
        说明：UserStatue 维护登录用户信息 （本地）
        Session    保存用户登录账户信息 （保存数据库）

        ###登录
        验证输入账户
                数据库查询账户(name),比对
                        成功：
                                跳转到登录成功界面
                                修改UserStatue
                                设置Cookie——组装Session
                        失败：
                                提升Error
        ###登出
                1、修改UserStatue
                2、删除session
                3、跳转界面到首页
        ###注册
        1、 获取From表单数据并验证
                        验证成功：
                                数据库添加用户
                                        成功：跳转登录界面
                        验证失败:
                                提示Error
                2、AjAx请求验证注册用户名称是否使用并提示
     
###图书管理功能
        图书增删改查操作
###购物车管理
        注意：购物车管理用户状态必须是已登录的状态
        ###购物车添加管理
        注意：每做一次cart_item 操作 需要同步更新cart操作

        1、首页中添加Click事件，ajax请求
        2、新增CartManagementHandler.go 管理购物车
        3、func MangaeCartAdd (){}
                    1、获取bookid from表单
                    2、获取book信息 bookid
                    3、获取用户信息 userid
                    4、添加操作
                        判断用户是否存在购物车   
                            不存在
                                1、创建购物车
                                2、创建购物车中购物项
                            存在
                                1、 查询该book是否已经存在购物项中 
                                        存在    更新count和amount
                                        不存在   创建购物项

        ###购物车查询管理
        1、获取用户信息 userid
        2、通过userid 查询出用户购物车信息
        3、遍历用户购物车中购物项并显示

        ###购物车清空操作
        1、userID
        2、查询CartID 
        3、删除购物项
        4、删除购物车
        ###删除购物车购物项
        1、CartIyemID
        2、数据库删除
        ###跳转支付，生成订单

###订单管理
        ###支付操作
                1、添加订单
                2、添加订单项
                3、清空购物车


        ###发货和完成订单
                 订单状态修改



###项目补充功能——待补充
1、文件上传
book添加时，增加图片上传即文件上传的功能
2、前台分页处理

###开发问题

1、Jquery使用问题

            POST请求中参数不能使用模板字符串,
            如果使用，userName := r.PostFormValue("username")无法获取参数值

            var url = "/verifyName";
            //设置请求参数
            // var params = `{username: ${username} }`;//JQuery 模板字符串使用有问题
            var params = { "username": username };
            //发送Ajax请求
            $.post(url, params, function (res) {
                //将显示提示信息的span元素显示
                $("#msg").show();
                //将响应信息设置到span元素中
                $("#msg").html(res);
            });


