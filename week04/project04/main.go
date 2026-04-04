package main
import(
	"fmt"
	"sync"
)

func test01Send(intChan chan int){
	for i:=1;i<=100;i++{
		intChan<-i*i
	}
	close(intChan)
}
func test01Receive(intChan chan int,exitChan chan bool){
	for v:=range intChan{
		fmt.Println(v)
	}
	exitChan<-true
	close(exitChan)
}
func test02Send(intChan chan int){
	for i:=1;i<=100;i++{
		intChan<-i
	}
	close(intChan)
}
func test02Prime(intChan chan int,primeChan chan int){
	for v:=range intChan{
		var i int
		for i=2;i<=v/2;i++{
			if(v%i==0){
				break
			}
		}
		if (i==v/2+1){
			primeChan<-v
			fmt.Println("输入素数",v)
		}
	}
	close(primeChan)
}
func test02Receive(primeChan chan int,exitChan chan bool){
	for v:=range primeChan{
		fmt.Println("输出素数",v)
	}
	exitChan<-true
	close(exitChan)
}


var (
	wg sync.WaitGroup
)
func Test03Send(intChan chan int){
	for i:=1;i<=100;i++{
		intChan<-i
	}
	close(intChan)
	wg.Done()
}
func Test03Prime(intChan chan int,primeChan chan int){
	for v:=range intChan{
		var i int
		for i=2;i<=v/2;i++{
			if(v%i==0){
				break
			}
		}
		if (i==v/2+1){
			primeChan<-v
			fmt.Println("输入素数",v)
		}
	}
	close(primeChan)
	wg.Done()
}
func Test03Receive(primeChan chan int){
	for v:=range primeChan{
		fmt.Println("输出素数",v)
	}
	wg.Done()
}
func main(){
	intChan :=make(chan int,50)
	// exitChan :=make(chan bool)
	primeChan :=make(chan int,100)

	// go test01Send(intChan)
	// go test01Receive(intChan,exitChan)
	// for _=range exitChan{

	// }
	// fmt.Println("执行完...")
	// go test02Send(intChan)
	// go test02Prime(intChan,primeChan)
	// go test02Receive(primeChan,exitChan)
	// for _=range exitChan{
	// }
	// fmt.Println("执行完...")
	wg.Add(3)
	go Test03Send(intChan)
	go Test03Prime(intChan,primeChan)
	go Test03Receive(primeChan)
	wg.Wait()
	fmt.Println("执行完...")
}