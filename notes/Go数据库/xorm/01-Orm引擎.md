# Orm引擎
在xorm里面，可以同时存在多个Orm引擎，一个Orm引擎称为Engine，一个Engine一般只对应一个数据库。Engine通过调用xorm.NewEngine生成。

### 创建引擎，driverName, dataSourceName和database/sql接口相同**
```
engine, err := xorm.NewEngine(driverName, dataSourceName)
```
- driverName :数据库名称
- dataSourceName : 数据库连接参数
> 一般情况下如果只操作一个数据库，只需要创建一个engine即可。engine是GoRoutine安全的。

创建完成engine之后，并没有立即连接数据库，此时可以通过engine.Ping()来进行数据库的连接测试是否可以连接到数据库。

engine可以通过engine.Close来手动关闭，但是一般情况下可以不用关闭，在程序退出时会自动关闭。

## 创建Engine组
```
dataSourceNameSlice := []string{masterDataSourceName, slave1DataSourceName, slave2DataSourceName}
engineGroup, err := xorm.NewEngineGroup(driverName, dataSourceNameSlice)
masterEngine, err := xorm.NewEngine(driverName, masterDataSourceName)
slave1Engine, err := xorm.NewEngine(driverName, slave1DataSourceName)
slave2Engine, err := xorm.NewEngine(driverName, slave2DataSourceName)
engineGroup, err := xorm.NewEngineGroup(masterEngine, []*Engine{slave1Engine, slave2Engine})
```
所有使用 engine 都可以简单的用 engineGroup 来替换。

## 日志
日志是一个接口，通过设置日志，可以显示SQL，警告以及错误等，默认的显示级别为INFO。

- engine.ShowSQL(true)，则会在控制台打印出生成的SQL语句；
- engine.Logger().SetLevel(core.LOG_DEBUG)，则会在控制台打印调试及以上的信息；

如果希望将信息不仅打印到控制台，而是保存为文件，那么可以通过类似如下的代码实现，NewSimpleLogger(w io.Writer)接收一个io.Writer接口来将数据写入到对应的设施中。
```
f, err := os.Create("sql.log")
if err != nil {
    println(err.Error())
    return
}
engine.SetLogger(xorm.NewSimpleLogger(f))
```
如果希望将日志记录到syslog中，也可以如下：
```
logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
if err != nil {
	log.Fatalf("Fail to create xorm system logger: %v\n", err)
}

logger := xorm.NewSimpleLogger(logWriter)
logger.ShowSQL(true)
engine.SetLogger(logger)
```

## 连接池
engine内部支持连接池接口和对应的函数。

- 如果需要设置连接池的空闲数大小，可以使用engine.SetMaxIdleConns()来实现。
- 如果需要设置最大打开连接数，则可以使用engine.SetMaxOpenConns()来实现。