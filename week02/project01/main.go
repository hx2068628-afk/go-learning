package main
import (
	"fmt"
)

type Person struct{
	Name string
	Age int
	Score float64
}
func (p *Person)String()string{
	return fmt.Sprintf("%s的年龄为%d,成绩为%f",p.Name,p.Age,p.Score)
}

type Rectangle struct{
	length float64
	width float64
}

func (R *Rectangle)Compute()float64{
	return R.length*R.width
}

type Book struct{
	Name string
	Year string
	IsGood bool
}

func (b *Book)isGood()bool{
	b.Name="五万年"
	if(b.IsGood==true){
		return true
	}else{
		return false
	}
}

func main(){
	var p Person =Person{"张三",16,90}
	fmt.Println(&p)
	var R Rectangle
	R.length=5
	R.width=3
	res :=R.Compute()
	fmt.Println("面积为",res)
	var book Book = Book{"三万年","2050",true}
	isGood:=book.isGood()
	fmt.Println(book.Name,"是否值得读:",isGood)
}