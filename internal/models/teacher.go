package models

type Teacher struct {
	ID       int       `gorm:"primary_key" json:"id"`
	Email    string    `gorm:"unique" json:"email"`
	Students []Student `gorm:"many2many:teacher_student;" json:"students"`
}
