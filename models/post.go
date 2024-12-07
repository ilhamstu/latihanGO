package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Nik       string `json:"nik"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Mahasiswa struct {
	Id        int        `json:"id" gorm:"primary_key"`
	Nim       string     `json:"nim"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Rekening  Rekening   `json:"rekening" gorm:"foreignKey:MahasiswaID"`
	Presensi  []Presensi `json:"presensi" gorm:"foreignKey:MahasiswaID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Rekening struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Norek       string `json:"norek"`
	Bank        string `json:"bank"`
	Name        string `json:"name"`
	MahasiswaID int    `json:"mahasiswa_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Presensi struct {
	Id          int       `json:"id" gorm:"primary_key"`
	MahasiswaID int       `json:"mahasiswa_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
