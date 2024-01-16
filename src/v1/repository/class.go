package repository

type Class struct {
	Name string `json:"name" pg:",pk"`
	CT   int    `json:"class_teacher_id"`
}

