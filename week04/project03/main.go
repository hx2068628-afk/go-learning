package main
import (
	"fmt"
	"sync"
	"time"
)
var (
	wg sync.WaitGroup
	lock sync.Mutex
)
func test01(){
	lock.Lock()
	for i:=0;i<10;i++{
	time.Sleep(time.Second)
	fmt.Printf("第%d个goroutine\n",i+1)		
	}
	wg.Done()
	lock.Unlock()
}
func test02(){
	lock.Lock()
	for i:=0;i<10;i++{
			time.Sleep(time.Second)
	fmt.Printf("第-%d个goroutine\n",i+1)		
	}
	lock.Unlock()
	wg.Done()
}
func test03(){
	fmt.Println("第3个goroutine")
	wg.Done()
}
func test04(nums map[int]int,i int){
		lock.Lock()
		nums[i]=i+1
		lock.Unlock()		
}
func test05(nums map[int]int){
	// lock.Lock()
	for _,v :=range nums{
		fmt.Println(v)
	}
	// lock.Unlock()
}
func main(){
	wg.Add(2)
	go test01()
	go test02()
	fmt.Println("阻塞中.....")
	wg.Wait()
	fmt.Println("无阻塞")
	// var nums map[int]int =make(map[int]int)
	// for i:=0;i<100;i++{
	// 	go test04(nums,i)		
	// }
	// time.Sleep(time.Second*3)
	// test05(nums)
	
	
	
}