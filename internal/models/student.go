package models

type Student struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Email string `gorm:"unique" json:"email"`
	//TODO: Change Status to enum
	Status   string    `json:"status"`
	Teachers []Teacher `gorm:"many2many:teacher_student;" json:"teachers"`
}
