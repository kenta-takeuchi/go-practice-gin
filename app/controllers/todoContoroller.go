package controllers

import (
	"github.com/gin-gonic/gin"
	"goTodo/app/models"
	"strconv"
)

//Index
func Index(ctx *gin.Context) {
	todos := models.GetAllTodo()
	tags := models.GetAllTag()
	ctx.HTML(200, "index.html", gin.H{
		"tags": tags,
		"todos": todos,
	})
}

//Create
func TodoAdd(ctx *gin.Context) {
	ctx.HTML(200, "addTodo.html", gin.H{})
}

//Create
func TodoRegist(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	todo := models.NewTodo(text, status)
	todo.Insert()
	ctx.Redirect(302, "/")
}

//Detail
func TodoDetail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := models.GetOneTodo(id)
	ctx.HTML(200, "detailTodo.html", gin.H{"todo": todo})
}

//Update
func TodoUpdate(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	todo := models.GetOneTodo(id)
	todo.Text = text
	todo.Status = status
	todo.Update()
	ctx.Redirect(302, "/")
}

//削除確認
func TodoDeleteCheck(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todo := models.GetOneTodo(id)
	ctx.HTML(200, "deleteTodo.html", gin.H{"todo": todo})
}

//Delete
func TodoDelete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todo := models.GetOneTodo(id)
	todo.Delete()
	ctx.Redirect(302, "/")
}

//Create
func TagAdd(ctx *gin.Context) {
	ctx.HTML(200, "addTag.html", gin.H{})
}

//Create
func TagRegist(ctx *gin.Context) {
	name := ctx.PostForm("name")
	tag := models.NewTag(name)
	tag.Insert()
	ctx.Redirect(302, "/")
}

//Detail
func TagDetail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := models.GetOneTag(id)
	ctx.HTML(200, "detailTag.html", gin.H{"todo": todo})
}