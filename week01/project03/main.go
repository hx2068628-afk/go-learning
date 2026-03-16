package main
import(
	"fmt"
	"math/rand"
)

//冒泡排序
func test1(){
	var arr []int
	arr = make([]int,5)
	for i:=0;i<5;i++{
		arr[i]=rand.Intn(100)
	}
	for i:=0;i<5;i++{
		for j:=0;j<5-i-1;j++{
			if(arr[j]>arr[j+1]){
				var tep = arr[j+1]
				arr[j+1]=arr[j]
				arr[j]=tep
			}
		}
	}
	fmt.Println(arr)
}
//统计字符串中每个字符出现的次数
func test2(){
	var str string
	fmt.Println("请输入一个字符串:")
	fmt.Scanf("%s",&str)
	str1 :=[]byte(str)
	var m map[byte]int =make(map[byte]int)
	for i:=0;i<len(str1);i++{
		m[str1[i]]+=1
	}
	for k,v :=range m{
		fmt.Printf("%c出现了%d次\n",k,v)
	}
}
//返回两个数的最大值和最小值
func test3(a ,b int)(max int,min int){
	if(a>=b){
		max=a
		min=b
	}else{
		max=b
		min=a
	}
	return
}
//根据输入的运算符进行简单的加减乘除运算
func test4(a,b int,op byte){
	switch op{
	case '+': fmt.Printf("%d+%d=%d",a,b,a+b) 
	case '-': fmt.Printf("%d-%d=%d",a,b,a-b) 
	case '*': fmt.Printf("%d*%d=%d",a,b,a*b) 
	case '/': fmt.Printf("%d/%d=%d",a,b,a/b) 
	default: fmt.Println("输入的运算符有误")
	}
}
//入栈
func push(arr []int)[]int{
	var i int
	fmt.Println("入栈\n请输入一个整数：")
	fmt.Scanf("%d",&i)
	arr =append(arr,i)
	return arr
}
//出栈
func pop(arr []int)[]int{
	fmt.Println("出栈元素为",arr[len(arr)-1])
	arr = arr[:len(arr)-1]
	return arr
}
//栈的应用
func test5(){
	var arr []int =make([]int,0)
	var i=1
	for i!=0{
		fmt.Println("输入1为入栈，输入2为出栈，输入0为结束")
		fmt.Scanf("%d",&i)
		if i==1{
			arr=push(arr)
		}else if i==2{
			arr=pop(arr)
		}
		fmt.Println(arr)
	}

}

func main(){
	// test1()
	// test2()
	// a,b :=test3(10,20)
	// fmt.Printf("max=%d,min=%d",a,b)
	// test4(10,20,'+')
	test5()
}