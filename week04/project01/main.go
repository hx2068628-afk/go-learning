package main
import (
	"fmt"
	"time"
)

func test01(){
	for i:=0;i<5;i++{
		go fmt.Println("这是第",i,"个协程")
	}
	time.Sleep(time.Microsecond*20)
	fmt.Println("over")
}

func main(){
	test01()

}