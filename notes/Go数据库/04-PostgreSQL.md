# PostgreSQL
PostgreSQL 是一个自由的对象-关系数据库服务器(数据库管理系统)，它在灵活的 BSD-风格许可证下发行。它提供了相对其他开放源代码数据库系统(比如 MySQL 和 Firebird)，和对专有系统比如 Oracle、Sybase、IBM 的 DB2 和 Microsoft SQL Server的一种选择。

PostgreSQL和MySQL比较，它更加庞大一点，因为它是用来替代Oracle而设计的。所以在企业应用中采用PostgreSQL是一个明智的选择。

## 驱动
Go实现的支持PostgreSQL的驱动也很多，因为国外很多人在开发中使用了这个数据库。

* https://github.com/lib/pq 支持database/sql驱动，纯Go写的
* https://github.com/jbarham/gopgsqldriver 支持database/sql驱动，纯Go写的
* https://github.com/lxn/go-pgsql 支持database/sql驱动，纯Go写的

## 示例
### [lib/pq]()
**安装**
```
go get github.com/lib/pq
```
**使用**
```
 import _ "https://github.com/lib/pq"
```
