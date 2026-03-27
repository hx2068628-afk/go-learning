package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
)

type Student struct{
	Name string
	Age int
	Score float64
}

func test01(length int)[]Student{
	var stus []Student
	for i:=0;i<length;i++{
		var stu Student
		stu.Name=fmt.Sprintf("学生%d",i+1)
		stu.Age=18+i
		stu.Score=90
		stus=append(stus,stu)
	}
	return stus
}
func test02(stus []Student)[]byte{
	data,err1:=json.Marshal(stus)
	if err1!=nil{
		fmt.Println(err1)
	}
	// fmt.Println(string(data))
	file_path:="/home/xr/goproject/go_code/week03/project04/test.json"
	file,err2:=os.OpenFile(file_path,os.O_CREATE|os.O_WRONLY,0666)
	if err2!=nil{
		fmt.Println(err2)
	}
	writer:=bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	return data
}
func test03(){
	file_path:="/home/xr/goproject/go_code/week03/project04/test.json"
	file,err2:=os.OpenFile(file_path,os.O_RDONLY,0666)
	if err2!=nil{
		fmt.Println(err2)
	}	
	reader :=bufio.NewReader(file)
	for {
		data,err:=reader.ReadString('}')
		if(err!=nil){
			fmt.Println(err)
			break
		}
		fmt.Printf("%T ",data)
		fmt.Println(data)
		

	}
}



func main(){
	stus:=test01(5)
	data:= test02(stus)
	// test03()
	err := json.Unmarshal(data,&stus)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%T ",stus)
	fmt.Println(stus)
	fmt.Printf("%T",(data))
	fmt.Println(string(data))
}
