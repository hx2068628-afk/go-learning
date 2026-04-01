package main 
import (
	"fmt"
)

func send(intChan chan int){
	for i:=1;i<=3;i++{
		intChan<-i
	}
	close(intChan)

}
func receive(intChan chan int,overChan chan bool){
	for v:=range intChan{
		fmt.Println(v)
	}
	for{
		v,ok:=<-intChan
		if !ok{
			fmt.Println(ok)
			break
		}
			fmt.Println(v)

	}
	overChan<-true
	close(overChan)
}


func main(){
	intChan :=make(chan int,1)
	overChan :=make(chan bool,1)
	go send(intChan)
	go receive(intChan,overChan)
	for{
		_,ok:=<-overChan
		if !ok{
			fmt.Println(ok)
			break
		}
	}
}