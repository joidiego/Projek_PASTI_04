package models

import (
	config "komentar/Config"

	"github.com/jinzhu/gorm"

	"time"
)

var db *gorm.DB

type Comment struct {
	ID		uint      `gorm:"primary_key" json:"id"`
	Content	string    `json:"content"`
	Nama	string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Comment{})
}

func (c *Comment) CreateComment() *Comment {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllComments() []Comment {
	var comments []Comment
	db.Find(&comments)
	return comments
}

func GetCommentById(id int64) (*Comment, *gorm.DB) {
	var comment Comment
	db := db.Where("id = ?", id).First(&comment)
	return &comment, db
}

func DeleteComment(id int64) Comment {
	var comment Comment
	db.Where("id = ?", id).Delete(&comment)
	return comment
}
