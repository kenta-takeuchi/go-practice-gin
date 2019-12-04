package models

import (
	"github.com/jinzhu/gorm"
)
type Todo struct {
	gorm.Model
	Text   string
	Status string
	TagID  uint
}

func NewTodo(text, status string) *Todo {
	return &Todo{
		Text:text,
		Status:status,
	}
}

// 新規登録
func (todo *Todo) Insert() {
	db := DbConnect()
	defer db.Close()
	db.Create(&todo)
}

// 更新
func (todo *Todo) Update() {
	db := DbConnect()
	defer db.Close()
	db.Save(&todo)
}

// 削除
func (todo *Todo) Delete() {
	db := DbConnect()
	defer db.Close()
	db.Delete(&todo)
}

//DB全取得
func GetAllTodo() []Todo {
	db := DbConnect()
	defer db.Close()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	return todos
}

//DB一つ取得
func GetOneTodo(id int) *Todo {
	db := DbConnect()
	defer db.Close()
	var todo Todo
	db.First(&todo, id)
	return &todo
}