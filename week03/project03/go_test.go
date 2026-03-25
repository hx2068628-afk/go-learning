package main
import (
	"testing"
	"os"
	"bufio"
	"fmt"
)

func TestAddUpper(t *testing.T){
	i,j:=1,2
	res := AddUpper(i,j)
	if(res!=3){
		t.Fatalf("测试结果错误，预期结果%v，实际结果%v",3,res)
	}
	t.Logf("测试结果正确，预期结果%v，实际结果%v",3,res)
}
func TestGetSub(t *testing.T){
	i,j:=1,2
	res := GetSub(i,j)
	if(res!=-1){
		t.Fatalf("测试结果错误，预期结果%v，实际结果%v",-1,res)
	}
	t.Logf("测试结果正确，预期结果%v，实际结果%v",-1,res)
}

func TestStore(t *testing.T){
	m :=map[string]string{"name":"张三","age":"18"}
	err:=Store(m)
	if err != nil{
		t.Fatalf("测试结果错误，%v",err)
	}
	t.Logf("测试结果正常")
	file_path:="/home/xr/goproject/test.txt"
	file,_ :=os.OpenFile(file_path,os.O_RDONLY,0666)
	reader:=bufio.NewReader(file)
	for{
		data,err:=reader.ReadString(',')
		fmt.Println(data)
		if err !=nil{
			break
		}
	}
}
func TestIsFile(t *testing.T){
	file_path:="/home/xr/goproject/test1.txt"
	err:=IsFile(file_path)
	if err!=nil{
		t.Fatalf("该文件不存在,%v",err)
	}
	t.Logf("该文件存在")
}