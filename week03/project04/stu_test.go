package main

import(
	"testing"
)

func TestTest01(t *testing.T){
	stus:=test01(3)
	if(len(stus)!=3){
		t.Fatalf("读取数据数量错误,期望值为%v，实际值为%v",3,len(stus))
	}
	t.Logf("读取数据数量正确,期望值为%v，实际值为%v",3,len(stus))
}