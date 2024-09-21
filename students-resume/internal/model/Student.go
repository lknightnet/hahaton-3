package model

type Student struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email"`
}
