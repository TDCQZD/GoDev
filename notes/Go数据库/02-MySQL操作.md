# MySQL数据库
MySQL以免费、开源、使用方便为优势成为了很多Web开发的后端数据库存储引擎。
## MySQL驱动
Go中支持MySQL的驱动目前比较多，有如下几种，有些是支持database/sql标准，而有些是采用了自己的实现接口,常用的有如下几种:

* https://github.com/go-sql-driver/mysql 支持database/sql，全部采用go写。
* https://github.com/ziutek/mymysql 支持database/sql，也支持自定义的接口，全部采用go写。
* https://github.com/Philio/GoMySQL 不支持database/sql，自定义接口，全部采用go写。

##  示例
### [go-sql-driver/mysql]()
**安装**
```
go get -u github.com/go-sql-driver/mysql
```
**使用**
```
import _ "github.com/go-sql-driver/mysql"
```
