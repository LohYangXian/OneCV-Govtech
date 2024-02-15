package models

type Student struct {
	Email    string
	Status   string
	Teachers []Teacher
}
