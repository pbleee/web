package models

import (
	"gorm.io/gorm"
)

// user
type User struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(100);not null" json:"nama"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

// course
type Course struct {
	gorm.Model
	Judul     string `gorm:"type:varchar(100);not null" json:"judul"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}

// enrollment
type Enrollment struct {
	gorm.Model
	UserID   uint `gorm:"not null" json:"user_id"`
	CourseID uint `gorm:"not null" json:"course_id"`
}

// Module
type Module struct {
	gorm.Model
	CourseID uint   `gorm:"not null" json:"course_id"`
	Judul    string `gorm:"type:varchar(255);not null" json:"judul"`
	Konten   string `gorm:"type:text" json:"konten"`
}

// score
type Grade struct {
	gorm.Model
	UserID   uint    `gorm:"not null" json:"user_id"`
	ModuleID uint    `gorm:"not null" json:"module_id"`
	Score    float64 `gorm:"type:decimal(5,2)" json:"score"`
}
