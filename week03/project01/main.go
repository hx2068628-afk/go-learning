package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func test01(){
	file,err:=os.OpenFile("/home/xr/goproject/go_code/test.txt",os.O_CREATE|os.O_RDWR,0666)
	if err !=nil{
		fmt.Println(err)
	}
	defer file.Close()
	writer :=bufio.NewWriter(file)
	reader :=bufio.NewReader(file)
	for i:=0;i<5;i++{
		writer.WriteString("hello world\n")
	}
	writer.Flush()
	file.Seek(0, 0)

	for i:=0;i<5;i++{
		content,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(content)
	}
}


func test02(){
	des_path:="/home/xr/goproject/go_code/test2.txt"
	src_path:="/home/xr/goproject/go_code/test.txt"
	file1,err1:=os.OpenFile(src_path,os.O_RDONLY,0666)
	if err1!=nil{
		fmt.Println(err1)	
	}
	defer file1.Close()
	file2,err2:=os.OpenFile(des_path,os.O_CREATE|os.O_WRONLY,0666)
	if err2!=nil{
		fmt.Println(err2)
	}
	defer file2.Close()
	reader :=bufio.NewReader(file1)
	writer :=bufio.NewWriter(file2)
	io.Copy(writer,reader)
}

func main(){

	test02()

}