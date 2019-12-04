package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"goTodo/app/controllers"
	"goTodo/app/models"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "app/assets")
	router.LoadHTMLGlob("app/views/*.html")

	// DB初期化処理
	models.DbInit()

	//Index
	router.GET("/", controllers.Index)
	//Regist
	router.GET("/add", controllers.TodoAdd)
	//Regist
	router.POST("/regist", controllers.TodoRegist)
	//Detail
	router.GET("/detail/:id", controllers.TodoDetail)
	//Update
	router.POST("/update/:id", controllers.TodoUpdate)
	//削除確認
	router.GET("/delete_check/:id", controllers.TodoDeleteCheck)
	//Delete
	router.POST("/delete/:id", controllers.TodoDelete)

	//Regist
	router.GET("/tagAdd", controllers.TagAdd)
	//Regist
	router.POST("/tagRegist", controllers.TagRegist)

	router.Run()
}