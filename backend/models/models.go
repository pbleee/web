package models

import (
	"time"

	"gorm.io/gorm"
)

// user
type User struct {
	ID       uint           `gorm:"primaryKey"`
	Nama     string         `gorm:"type:varchar(100);not null"`
	Email    string         `gorm:"type:varchar(100);unique;not null"`
	Password string         `gorm:"type:varchar(255);not null"`
	CreateAt time.Time      `gorm:"autoCreateTime" json:"create_at"`
	UpdateAt time.Time      `gorm:"autoUpdateTime" json:"update_at"`
	DeleteAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

// course
type Course struct {
	ID        uint           `gorm:"primaryKey"`
	Judul     string         `gorm:"type:varchar(100);not null"`
	Deskripsi string         `gorm:"type:text"`
	CreateAt  time.Time      `gorm:"autoCreateTime" json:"create_at"`
	UpdateAt  time.Time      `gorm:"autoUpdateTime" json:"update_at"`
	DeleteAt  gorm.DeletedAt `gorm:"column:delete_at;index" json:"delete_at"`
}

// enrollment
type Enrollment struct {
	ID       uint           `gorm:"primaryKey"`
	UserID   uint           `gorm:"not null"`
	CourseID uint           `gorm:"not null"`
	CreateAt time.Time      `gorm:"column:create_at"`
	UpdateAt time.Time      `gorm:"column:update_at"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;index"`
}

// Module
type Module struct {
	ID       uint           `gorm:"primaryKey"`
	CourseID uint           `gorm:"not null"`
	Judul    string         `gorm:"type:varchar(255);not null"`
	Konten   string         `gorm:"type:text"`
	CreateAt time.Time      `gorm:"column:create_at"`
	UpdateAt time.Time      `gorm:"column:update_at"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;index"`
}

// score
type Grade struct {
	ID       uint           `gorm:"primaryKey"`
	UserID   uint           `gorm:"not null"`
	ModuleID uint           `gorm:"not null"`
	Score    float64        `gorm:"type:decimal(5,2)"`
	CreateAt time.Time      `gorm:"column:create_at"`
	UpdateAt time.Time      `gorm:"column:update_at"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;index"`
}
