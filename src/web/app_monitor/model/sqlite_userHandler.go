package model

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite_userHandler struct {
	database *sql.DB
}

func (s *sqlite_userHandler) CreateUser(id string, password string, sessionID string) *UserInfo {
	stmt, err := s.database.Prepare("INSERT INTO user (id, password, session_id, created_at, updated_at) VALUES (?, ?, ?, datetime('now'), datetime('now'))")
	if err != nil {
		log.Println("[CreateUser] error: ", err)
		return nil
	}
	_, err = stmt.Exec(id, password, sessionID)
	if err != nil {
		log.Println("[CreateUser] error: ", err)
		return nil
	}

	var user UserInfo
	user.Id = id
	user.SessionId = sessionID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return &user
}

func (s *sqlite_userHandler) ReadUserId(Id string) []*UserInfo {
	users := []*UserInfo{}
	rows, err := s.database.Query("SELECT id, password, session_id, created_at, updated_at FROM user WHERE id = ?", Id)
	if err != nil {
		log.Println("[ReadUser] error: ", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var user UserInfo
		rows.Scan(&user.Id, &user.Password, &user.SessionId, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, &user)
	}
	return users
}

func (s *sqlite_userHandler) ReadUserSession(sessionId string) []*UserInfo {
	users := []*UserInfo{}
	rows, err := s.database.Query("SELECT id, password, session_id, created_at, updated_at FROM user WHERE session_id = ?", sessionId)
	if err != nil {
		log.Println("[ReadUser] error: ", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var user UserInfo
		rows.Scan(&user.Id, &user.Password, &user.SessionId, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, &user)
	}
	return users
}

func (s *sqlite_userHandler) UpdateUser(id string, sessionId string) bool {
	stmt, err := s.database.Prepare("UPDATE user SET session_id = ? WHERE id = ?")
	if err != nil {
		log.Println("[UpdateUser] error: ", err)
		return false
	}
	rslt, err := stmt.Exec(sessionId, id)
	if err != nil {
		log.Println("[UpdateUser] error: ", err)
		return false
	}
	count, _ := rslt.RowsAffected()
	return count > 0
}

func (s *sqlite_userHandler) DeleteUser(id string) bool {
	stmt, err := s.database.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		log.Println("[DeleteUser] error: ", err)
		return false
	}
	rslt, err := stmt.Exec(id)
	if err != nil {
		log.Println("[DeleteUser] error: ", err)
		return false
	}
	count, _ := rslt.RowsAffected()
	return count > 0
}

func (s *sqlite_userHandler) Close() {
	s.database.Close()
}

func newSqlite_userHandler(databaseFilePath string) DatabaseHandler {
	log.Println("databaseFilePath: ", databaseFilePath)
	database, err := sql.Open("sqlite3", databaseFilePath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS user (
			id			STRING	PRIMARY KEY,
			password	STRING,
			session_id	STRING,
			created_at	DATETIME,
			updated_at	DATETIME
		);
		CREATE INDEX IF NOT EXISTS sessionIDIndexOnUser ON user (
			session_id ASC
		);`)
	statement.Exec()
	return &sqlite_userHandler{database: database}
}
