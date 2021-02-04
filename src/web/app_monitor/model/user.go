package model

import "time"

type UserInfo struct {
	Id        string    `json:"id"`
	Password  string    `json:"password"`
	SessionId string    `json:"session_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DatabaseHandler interface {
	CreateUser(id string, password string, sessionID string) *UserInfo
	ReadUserId(Id string) []*UserInfo
	ReadUserSession(sessionId string) []*UserInfo
	UpdateUser(id string, sessionID string) bool
	DeleteUser(id string) bool
	Close()
}

func NewDatabaseHandlerUser(databaseFilePath string) DatabaseHandler {
	handler := newSqlite_userHandler(databaseFilePath)
	return handler
}
