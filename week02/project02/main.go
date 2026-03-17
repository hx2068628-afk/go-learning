package main
import (
	"fmt"
)
//接口实现“不同图形求面积
type area interface{
	compute()float64
}
type rectangle struct{
	length float64
	width float64
}
type circle struct{
	radius float64
}
func (r rectangle )compute()float64{
	return r.length*r.width
}
func (c *circle)compute()float64{
	return 3.14*c.radius*c.radius
}
//写一个函数，分别用值传递和指针传递修改变量，对比效果
func test01(i *int){
	*i=100
}
//定义 Speaker 接口，让 Dog 和 Cat 都实现 Speak()
type Speaker interface{
	Speak()
}
type Dog struct{
	Name string
}
type Cat struct{

}
func (d Dog)Speak(){
	fmt.Println("汪汪汪")
}
func (c Cat)Speak(){
	fmt.Println("喵喵喵")
}
func main(){
	//01
	// var i area
	// var r rectangle
	// r.length=5
	// r.width=3
	// i=r
	// fmt.Println("矩形的面积为:",i.compute())
	// c :=circle{}
	// c.radius=2
	// i=&c
	// fmt.Println("圆的面积为:",i.compute())
	//02
	// i :=0
	// fmt.Println(i)
	// test01(&i)
	// fmt.Println(i)
	//03
	var s Speaker = Dog{"旺财"}
	var c Cat
	s.Speak()
	c.Speak()
}