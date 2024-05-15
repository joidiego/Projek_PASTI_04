package controllers

import (
	"encoding/json"
	"fmt"
	models "komentar/Models"
	utils "komentar/Utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewComment models.Comment

func GetComment(w http.ResponseWriter, r *http.Request) {
	NewComment := models.GetAllComments()
	res, _ := json.Marshal(NewComment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCommentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["commentId"]
	ID, err := strconv.ParseInt(commentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	commentDetails, _ := models.GetCommentById(ID)
	res, _ := json.Marshal(commentDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	CreateComment := &models.Comment{}
	utils.ParseBody(r, CreateComment)
	b := CreateComment.CreateComment()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["commentId"]
	ID, err := strconv.ParseInt(commentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	comment := models.DeleteComment(ID)
	res, _ := json.Marshal(comment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	// Verifikasi autentikasi pengguna
	user := getCurrentUser(r)
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var updateComment = &models.Comment{}
	utils.ParseBody(r, updateComment)
	vars := mux.Vars(r)
	commentID := vars["commentId"]

	// Parsing ID komentar
	ID, err := strconv.ParseInt(commentID, 0, 0)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	// Mendapatkan komentar dari database berdasarkan ID
	commentDetails, db := models.GetCommentById(ID)
	if commentDetails == nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	// Pengecekan apakah pengguna memiliki hak akses untuk memperbarui komentar
	if user.Role != "customer" || user.ID != commentDetails.UserID {
		http.Error(w, "Forbidden: You are not allowed to update this comment", http.StatusForbidden)
		return
	}

	// Memperbarui konten komentar
	if updateComment.Content != "" {
		commentDetails.Content = updateComment.Content
	}

	// Menyimpan perubahan ke database
	err = db.Save(&commentDetails).Error
	if err != nil {
		http.Error(w, "Failed to update comment", http.StatusInternalServerError)
		return
	}

	// Mengembalikan respons dengan detail komentar yang diperbarui
	res, err := json.Marshal(commentDetails)
	if err != nil {
		http.Error(w, "Error while marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getCurrentUser(r *http.Request) {
	panic("unimplemented")
}