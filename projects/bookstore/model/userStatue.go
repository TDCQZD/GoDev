package model

type UserStatue struct {
	IsLogin    bool
	UserID     string
	UserName   string
	PageInfor  *Page
	CartInfor  *Cart
	OrderID    string
	OrederItems []*OrederItem
	Oreders    []*Oreders
}
