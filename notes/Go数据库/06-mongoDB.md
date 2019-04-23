# mongoDB
MongoDB是一个高性能，开源，无模式的文档型数据库，是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。他支持的数据结构非常松散，采用的是类似json的bjson格式来存储数据，因此可以存储比较复杂的数据类型。Mongo最大的特点是他支持的查询语言非常强大，其语法有点类似于面向对象的查询语言，几乎可以实现类似关系数据库单表查询的绝大部分功能，而且还支持对数据建立索引。
## 驱动
* http://labix.org/mgo 【推荐】
* https://github.com/mongodb/mongo-go-driver
* https://github.com/go-mgo/mgo 【停止维护】

### mgo Getting started
To get the package, execute:
```
go get gopkg.in/mgo.v2
```
To import this package, add the following line to your code:
```
import "gopkg.in/mgo.v2"
```
## 示例

## 阿里云ESC 安装mongoDB 设置远程连接
1. 修改mongo.conf 
```
$ vim mongo.conf 
bind_ip = 0.0.0.0
auth = false
```
2. 启动mongo
```
mongod -f mongo.conf 
```