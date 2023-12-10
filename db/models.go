package db

import (
	"database/sql"
)

type Database interface {
	GetUsers()
	GetPosts()
	CreateUser()
	CreatePost()
	DeletePost()
}

// структура для таблицы бд с пользователями
type User struct {
	UserId       uint64        `json:"UserId"`
	UserName     string        `json:"UserName"`
	UserEmail    string        `json:"UserEmail"`
	PasswordHash string        `json:"PasswordHash"`
	CreatedAt    sql.NullInt64 `json:"CreatedAt"`
}

// структура для таблицы бд с постами
type Post struct {
	PostId      uint64        `json:"PostId"`
	UserId      sql.NullInt64 `json:"UserId"`
	CategoryId  sql.NullInt64 `json:"CategoryId"`
	PostTitle   string        `json:"PostTitle"`
	PostContent string        `json:"PostContent"`
	CreatedAt   sql.NullInt64 `json:"CreatedAt"`
}
