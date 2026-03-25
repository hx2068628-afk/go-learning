package main 
import (
	"encoding/json"
	"os"
	"bufio"
	"fmt"
)

func AddUpper(i,j int) int {
	return i+j+1
}
func GetSub(i,j int) int {
	return i-j
}

func Store(m map[string]string)error{
	file_path:="/home/xr/goproject/test.txt"
	file,err :=os.OpenFile(file_path,os.O_CREATE | os.O_WRONLY,0666)
	defer file.Close()
	data,_:=json.Marshal(m)
	fmt.Println(string(data))
	writer:=bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	return err
}

func IsFile(file_path string)error{
	_,err:=os.OpenFile(file_path,os.O_RDONLY,0666)
	return err
}
