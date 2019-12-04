package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"size:255"`
	Todos []Todo
}

func NewTag(name string) *Tag {
	return &Tag{
		Name:name,
	}
}

// 新規登録
func (tag *Tag) Insert() {
	db := DbConnect()
	defer db.Close()
	db.Create(&tag)
}

// 更新
func (tag *Tag) Update() {
	db := DbConnect()
	defer db.Close()
	db.Save(&tag)
}

// 削除
func (tag *Tag) Delete() {
	db := DbConnect()
	defer db.Close()
	db.Delete(&tag)
}

//DB全取得
func GetAllTag() []Tag {
	db := DbConnect()
	defer db.Close()
	var tags []Tag
	db.Order("created_at desc").Find(&tags)
	return tags
}

//DB一つ取得
func GetOneTag(id int) *Tag {
	db := DbConnect()
	defer db.Close()
	var tag Tag
	db.First(&tag, id)
	return &tag
}
