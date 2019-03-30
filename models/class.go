package models

//Class 课程
type Class struct {
	TotalElements int            `json:"totalElements"`
	TotalPages    int            `json:"totalPages"`
	Content       []ClassContent `json:"content"`
}

//ClassContent .
type ClassContent struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
