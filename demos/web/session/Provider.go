package main

// 抽象出一个Provider接口，用以表征session管理器底层存储结构。
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
	/*
		SessionInit函数实现Session的初始化，操作成功则返回此新的Session变量
		SessionRead函数返回sid所代表的Session变量，如果不存在，那么将以sid为参数调用SessionInit函数创建并返回一个新的Session变量
		SessionDestroy函数用来销毁sid对应的Session变量
		SessionGC根据maxLifeTime来删除过期的数据
	*/
}
