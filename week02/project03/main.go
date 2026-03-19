package main
import(
	"fmt"
	"errors"
)
//写一个除法函数，除数为 0 时返回错误
func test01(a , b float64)error{
	if(b!=0){
		fmt.Println(a/b)
		return nil
	}else{
		return errors.New("除数为0，报错")
	}
}

//用 defer 打印程序执行结束提示
func test02(){
	var i,j int
	i++
	j++
	defer func (i int,j int){
		fmt.Println(i+j,"函数执行完")
	}(i,j)
	i++
	j++
	fmt.Println(i+j)
}

//模拟一个简单登录函数，对空用户名或空密码做错误返回
func test03()(err error){
	var user string
	var password string
	defer func(){
		if (user=="user"){
			if (password=="123"){
				fmt.Println("successful")
				err =nil
			}else{
				err= errors.New("密码错误")
			}
		}else{
			err= errors.New("账号错误")
		}
	}()
	fmt.Scanf("%s %s",&user,&password)
	return nil
}


func main(){
	// a,b:=1.0,0.0
	// res:=test01(a,b)
	// fmt.Println(res)
	// test02()
	fmt.Println(test03())
}