package db

import (
	"database/sql"
	"time"
)

// type Database interface {
// 	GetUsers()
// 	GetPosts()
// 	CreateUser()
// 	CreatePost()
// 	DeletePost()
// }

// структура для таблицы бд с пользователями

type User struct {
	UserId       uint64        `json:"UserId"`
	UserName     string        `json:"UserName"`
	UserEmail    string        `json:"UserEmail"`
	PasswordHash string        `json:"PasswordHash"`
	CreatedAt    sql.NullInt64 `json:"CreatedAt"`
}

type Users struct {
	UserId    uint64        `json:"UserId"`
	UserName  string        `json:"UserName"`
	UserEmail string        `json:"UserEmail"`
	CreatedAt sql.NullInt64 `json:"CreatedAt"`
}

// структура для таблицы с категориями

type Category struct {
	CategoryId   uint64
	CategoryName string
}

// структура для таблицы бд с постами

type Post struct {
	PostId      uint64    `json:"PostId"`
	UserId      uint64    `json:"UserId"`
	CategoryId  uint16    `json:"CategoryId"`
	PostTitle   string    `json:"PostTitle"`
	PostContent string    `json:"PostContent"`
	CreatedAt   time.Time `json:"CreatedAt"`
}
