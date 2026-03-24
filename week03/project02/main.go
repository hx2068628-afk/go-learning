package main
import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"

)

type Student struct{
	Name string
	Age int
	Score float64
}

// test01 测试json序列化
func test01()string{
	var slice []Student
	var stu1 Student=Student{"张三",16,90}
	var stu2 Student=Student{"李四",18,88}
	slice=append(slice,stu1,stu2)
	data,err :=json.Marshal(slice)
	if err !=nil{
		fmt.Println(err)
	}
	return string(data)
}
// test02 测试json反序列化
func test02(str string){
	var slice []Student
	err :=json.Unmarshal([]byte(str),&slice)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(slice)
}
// test03 测试json序列化到文件
func test03(str string){
	file_path:="/home/xr/goproject/go_code/test.json"
	file,err := os.OpenFile(file_path,os.O_CREATE | os.O_WRONLY,0666)
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()
	writer :=bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()

}
func test04(){
	file_path:="/home/xr/goproject/go_code/test.json"
	file,err :=os.OpenFile(file_path,os.O_RDONLY,0666)
	if err != nil{
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	for {
		content,err1 := reader.ReadString('}')
		if err1!=nil{
			fmt.Println(err1)
			break
		}
		fmt.Println(content)
	}

}
func main(){
	// str:=test01()
	// test02(str)
	// test03(str)
	test04()
}