package model

import (
	"goweb_code/bookstore/utils"
)

type Oreders struct {
	ID          string  //订单号
	Time        string  //生成订单的时间
	TotalCount  int64   //订单中图书的总数量
	TotalAmount float64 //订单中图书的总金额
	State       int64   //订单的状态 0 未发货 1 已发货 2 交易完成
	UserID      string  //订单所属的用户

}

//生成CartID
func (order *Oreders) GetOrederID() string {
	return utils.CreateUUID()
}
//NoSend 未发货
func (order *Oreders) NoSend() bool {
	return order.State == 0
}

//SendComplate 已发货
func (order *Oreders) SendComplate() bool {
	return order.State == 1
}

//Complate 交易完成
func (order *Oreders) Complate() bool {
	return order.State == 2
}