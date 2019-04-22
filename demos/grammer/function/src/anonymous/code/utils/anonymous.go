package utils

// 在定义匿名函数时就直接调用
func Anonymous1(n1 ,n2 int) int  {
	/*定义匿名函数
	func (num1 ,num2 int) int  {
		return num1 + num2
	}
	*/
	res := func (num1 ,num2 int) int  {
		return num1 + num2
	}(n1,n2)
	return res
}
// 将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
func Anonymous2(n1 ,n2 int) int   {
	Anony := func (num1 ,num2 int) int  {
		return num1 + num2
	}
	return Anony(n1,n2)
}
// 如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在程序有效。
var Anonymous3 = func (n1 ,n2 int) int   {
	return n1+n2
}