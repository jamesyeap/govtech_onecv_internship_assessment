package models

type Student struct {
	Email	string	`gorm:"primaryKey" json:"email" `
	IsSuspended bool `json:"isSuspended"`
	Teachers []Teacher `gorm:"many2many:student_teachers"`
}

type Teacher struct {
	Email	string	`gorm:"primaryKey" json:"email" `
	Students []Student `gorm:"many2many:teacher_students"`
}