package entity

import (
	//"time"

	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string `json:"-"`
	//ผู้ดูแลระบบ 1 คน สามารถบันทึกข้อมูลนักศึกษาได้หลายคน
	Students []Student `gorm:"foreignKey:OfficerID"`
}

type Faculty struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//1 สำนักวิชา มีนักศึกษาหลายคน
	Students []Student `gorm:"foreignKey:FacultyID"`
	//1 สำนักวิชา มีอาจารย์หลายท่าน
	Teacher []Teacher `gorm:"foreignKey:FacultyID"`
}

type Collegeyear struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//1 ชั้นปี มีนักศึกษาหลายคน
	Students []Student `gorm:"foreignKey:CollegeyearID"`
}

type Teacher struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//Email string `gorm:"uniqueIndex"`
	// TeacherID ทำหน้าที่เป็น FK
	FacultyID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Faculty Faculty `gorm:"references:id"`
	//อาจารย์1คน มีนักศึกษาหลายคน
	Students []Student `gorm:"foreignKey:TeacherID"`
}

type Student struct {
	gorm.Model
	S_ID          string
	Name          string
	Gpax          float32
	Date_of_birth string
	Phone         string
	Parent        string

	OfficerID *uint
	Officer   Officer `gorm:"references:id"`

	CollegeyearID *uint
	Collegeyear   Collegeyear `gorm:"references:id"`

	FacultyID *uint
	Faculty   Faculty `gorm:"references:id"`

	TeacherID *uint
	Teacher   Teacher `gorm:"references:id"`
}
