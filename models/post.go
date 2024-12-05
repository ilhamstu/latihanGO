package models

type Post struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Nik      string `json:"nik"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Mahasiswa struct {
	Id       int      `json:"id" gorm:"primary_key"`
	Nim      string   `json:"nim"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Rekening Rekening `json:"rekening" gorm:"foreignKey:MahasiswaID"`
}

type Rekening struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Norek       string `json:"norek"`
	Bank        string `json:"bank"`
	Name        string `json:"name"`
	MahasiswaID int    `json:"mahasiswa_id"`
}
