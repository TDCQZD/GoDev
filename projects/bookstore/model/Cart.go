package model

import (
	"goweb_code/bookstore/utils"
)

type Cart struct {
	CartID      string      //购物车的id,uuid
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //购物车中图书的总数量，通过计算得到
	TotalAmount float64     //购物车中图书的总金额，通过计算得到
	UserID      string      //当前购物车所属的用户
}

//生成CartID
func (cart *Cart) GetCartID() string {
	return utils.CreateUUID()
}

//GetTotalCount 获取购物车中图书的总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历购物车中的购物项切片
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalAmount 获取购物车中图书的总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	//遍历购物车中的购物项切片
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.Amount
	}
	return totalAmount
}



//Data 结构
type Data struct {
	Amount      float64
	TotalAmount float64
	TotalCount  int64
}
