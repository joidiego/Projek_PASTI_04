package models

import (
	"komentar/config" // Menggunakan path yang sesuai dengan struktur proyek Anda

	"github.com/jinzhu/gorm"

	"time"
)


type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Content   string    `json:"content"`
	Nama      string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Inisialisasi koneksi database pada package config saat aplikasi dimulai
func init() {
	db = config.Connect() // Memanggil fungsi Connect() dari package config untuk mendapatkan koneksi database
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
	db := db.Where("id = ?", id).First(&comment) // Menggunakan variabel lokal db untuk mencari komentar dengan ID tertentu
	return &comment, db
}

func DeleteComment(id int64) Comment {
	var comment Comment
	db.Where("id = ?", id).Delete(&comment)
	return comment
}
