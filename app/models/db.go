package models

import "github.com/jinzhu/gorm"

//DB初期化
func DbInit() {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("データベースとの接続に失敗しました")
	}
	db.AutoMigrate(&Todo{}, &Tag{})
}

//DB初期化
func DbConnect() (db *gorm.DB) {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("データベースとの接続に失敗しました")
	}
	return db
}