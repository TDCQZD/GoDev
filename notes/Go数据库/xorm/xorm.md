# xorm
xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。
> https://github.com/go-xorm/xorm/blob/master/README_CN.md
## 特性
- 支持Struct和数据库表之间的灵活映射，并支持自动同步

- 事务支持

- 同时支持原始SQL语句和ORM操作的混合执行

- 使用连写来简化调用

- 支持使用Id, In, Where, Limit, Join, Having, Table, Sql, Cols等函数和结构体等方式作为条件

- 支持级联加载Struct

- Schema支持（仅Postgres）

- 支持缓存

- 支持根据数据库自动生成xorm的结构体

- 支持记录版本（即乐观锁）

- 内置SQL Builder支持

## 驱动支持
目前支持的Go数据库驱动和对应的数据库如下：

* Mysql: github.com/go-sql-driver/mysql

* MyMysql: github.com/ziutek/mymysql/godrv

* Postgres: github.com/lib/pq

* Tidb: github.com/pingcap/tidb

* SQLite: github.com/mattn/go-sqlite3

* MsSql: github.com/denisenkom/go-mssqldb

* MsSql: github.com/lunny/godbc

* Oracle: github.com/mattn/go-oci8 (试验性支持)

## 安装
```
go get github.com/go-xorm/xorm
```
## xorm操作
[xorm操作快速入门](xorm操作.md)



## 参考资料
* https://github.com/go-xorm/xorm
* https://godoc.org/github.com/go-xorm/xorm
* http://xorm.io/docs