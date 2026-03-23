package main
import (
	"fmt"
	"sort"
)
type Student struct{
	Name string
	Age int
	Score float64
}
type stuslice []Student

func (stu stuslice) Len() int{
	return len(stu)
}
func (stu stuslice) Less(i, j int) bool{
	return stu[i].Score<stu[j].Score
}
func (stu stuslice) Swap(i, j int){
	stu[i],stu[j]=stu[j],stu[i]
}



func main(){
	var stus stuslice =make(stuslice,5)
	for stu :=range stus{
		stus[stu].Name=fmt.Sprintf("张%d",stu)
		stus[stu].Age=17+stu
		stus[stu].Score=80.0+float64(stu)
	}
	fmt.Println(stus)
	sort.Sort(stus)
	for _,stu :=range stus{
		fmt.Println(stu)
	}
}