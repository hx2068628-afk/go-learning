package main
import(
	"testing"
	"fmt"
)


func TestTest(t *testing.T){
	intChan :=make(chan int,50)
	primeChan :=make(chan int,100)
	wg.Add(3)
	go Test03Send(intChan)
	go Test03Prime(intChan,primeChan)
	go Test03Receive(primeChan)
	wg.Wait()
	fmt.Println("已执行完")
	t.Logf("已执行完")
}