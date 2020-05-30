package models

import "github.com/jinzhu/gorm"

//Article ...轮播图
type Article struct {
	gorm.Model
	Title     string `json:"title" validate:"required"`
	ImageURL  string `json:"image_url" validate:"required"`
	Contents  string `json:"contents" validate:"required"`
	SubjectID int64  `json:"subject_id"`
	Tags      string `json:"tags"`
}
