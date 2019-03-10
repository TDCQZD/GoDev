# SQL
## 执行SQL查询
也可以直接执行一个SQL查询，即Select命令。在Postgres中支持原始SQL语句中使用 ` 和 ? 符号。
```
sql := "select * from userinfo"
results, err := engine.Query(sql)
```
当调用Query时，第一个返回值results为`[]map[string][]byte`的形式。
## 执行SQL命令
也可以直接执行一个SQL命令，即执行Insert， Update， Delete 等操作。此时不管数据库是何种类型，都可以使用 ` 和 ? 符号。
```
sql = "update `userinfo` set username=? where id=?"
res, err := engine.Exec(sql, "xiaolun", 1) 
```