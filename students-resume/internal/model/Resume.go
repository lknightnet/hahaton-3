package model

type Resume struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	UserAge      int    `json:"user_age"`
	City         string `json:"city"`
	University   string `json:"university"`
	Course       string `json:"course"`
	CourseNumber int    `json:"course_number"`
	Status       bool   `json:"status"`
	Description  string `json:"description"`
	DocumentID   int    `json:"document_id"`
}
