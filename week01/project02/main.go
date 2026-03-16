package main
import (
	"fmt"
)

//1到100的和
func test1(){
	var i int =0
	for j := 1;j<=100;j++{
		i+=j
	}
	fmt.Printf("1加到100的和为：%d",i)
}
//判断一个数是否为偶数
func test2(a int)bool{
	if a%2==0{
		return true
	}else{
		return false
	}
}
//输出1-9的乘法表
func test3(){
	fmt.Println("输出1-9的乘法表")
	for i :=1;i<=9;i++{
		for j :=1;j<=i;j++{
			fmt.Printf("%d*%d=%d\t",i,j,i*j)
		}
		fmt.Println()
	}
}
//判断一个数是否为质数
func test4(a int)bool{
	for i :=2;i<=a/2;i++{
		if a%i==0{
			return false
		}
	}
	return true
}
//输出斐波那契数列的前九个数
func test5(){
	a :=0
	b :=1
	fmt.Printf("%d %d ",a,b)
	for i :=1; i<=8;i++{
		c :=a+b
		a=b
		b=c
		fmt.Printf("%d ",c)
	}
}

func main(){
	//test1()
	// fmt.Println("请输入一个整数:")
	// var a int
	// fmt.Scanf("%d",&a)
	// var b bool=test2(a)
	// fmt.Printf("偶数为true，奇数为false，答案：%t",b)
	//test3()
	// var a int
	// fmt.Println("请输入一个整数:")
	// fmt.Scanf("%d",&a)
	// fmt.Printf("质数为true，不是质数为false，答案：%t",test4(a))
	test5()
}