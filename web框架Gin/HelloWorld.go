package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// db是数据库连接池
	db, err = sql.Open("mysql", "root:123456@tcp(docker.for.mac.localhost:3306)/hello?charset=utf8")
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "helloworld"})
	})
	err := initDB()
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	r.POST("/user/add", func(context *gin.Context) {
		var user User
		var result string
		context.BindJSON(&user)
		rs, err := db.Exec("INSERT INTO users(name, age) VALUES (?, ?)", user.Name, user.Age)
		if err != nil {
			result = "插入数所库失败" + err.Error()
		} else {
			id, _ := rs.LastInsertId()
			result = "插入成功:" + string(id)
		}
		context.JSON(200, gin.H{"msg": result})
	})
	r.GET("/user/query", func(context *gin.Context) {
		users := make([]User, 0)
		rows, _ := db.Query("SELECT id, name, age FROM users")
		defer rows.Close()
		for rows.Next() {
			var user User
			rows.Scan(&user.ID, &user.Name, &user.Age)
			users = append(users, user)
		}
		context.JSON(200, gin.H{"code": 200, "data": users})
	})
	r.Run(":8082")
}
