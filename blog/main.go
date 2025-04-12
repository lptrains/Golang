package main

import (
	"blog/controller/user"
	"blog/datebase"
	"blog/entity"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func m1() gin.HandlerFunc { //定义m1中间件 计时
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "username")
		fmt.Println("m1 in")
		c.Next() //调用后续处理函数
		cost := time.Since(start)
		fmt.Printf("cost: %v\n", cost) //服务器到客户端时在输出
	}

}
func main() {
	// 创建一个默认的Gin路由器
	r := gin.Default()
	r.Use(m1())                      //注册全局中间件
	r.Static("/css", "./static")     //静态文件()
	r.LoadHTMLGlob("templates/**/*") //模版解析(*为ant通配符)

	// 定义一个GET请求的路由 和 处理函数（匿名函数)
	// r.RouterGroup("/user")

	user.RouterRegistry(r)
	//连接数据库
	err := datebase.InitMySql()
	if err != nil {
		panic(err)
	}
	//退出程序关闭数据库
	defer datebase.Close()
	//绑定模型
	datebase.SqlSession.AutoMigrate(&entity.User{})
	// 启动Gin服务器，默认监听8080端口
	r.Run(":8080")
}

/*
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.Form)
		fmt.Fprintln(w, "helloworld")
		fmt.Fprintln(w, "ld")
		fmt.Fprintln(w, "password")
		fmt.Fprintln(w, r.Header)
	})

	/*server :=&http.Server{//自定义服务器配置
		Addr:   "localhost:8080",
		MaxHeaderBytes:  1<<20,
	}*/
/*http.ListenAndServe("localhost:8080", nil)
}*/
