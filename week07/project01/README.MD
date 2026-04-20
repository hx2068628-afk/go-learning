创建新项目后的步骤：
1 go mod init  (命令行)
2 项目导入:import(
    "github.com/gin-gonic/gin" 
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    )
3 go mod tidy (命令行)
4 air init
5 创建gin r := gin.Default()
6 运行gin r.Run()
7 数据库连接 	db, err1 := sql.Open("mysql", "账号:密码@(127.0.0.1:3306)/test?charset=utf8mb4")
8 接口类型 r.GET,r.POST,r.PUT,r.DELETE