package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          int64  `json:"_id"`
	UUID        string `json:"_uuid"`
	UserName    string `json:"user_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber int64  `json:"phone_number"`
	Active      bool   `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   int64  `json:"created_by"`
	UpdatedBy   int64  `json:"updated_by"`
}

func (user *User) Create(db *sql.DB) (*User,error){
	user.UUID = uuid.New().String()

	query := `INSERT INTO users (_uuid, user_name, first_name, last_name, email, phone_number, active, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result,err:= db.Exec(query,user.UUID,user.UserName,user.FirstName,user.LastName,user.Email,user.PhoneNumber,user.Active,user.CreatedBy,user.UpdatedBy)

	if err!=nil{
		return nil,err
	}

	user.ID, _ = result.LastInsertId()
	return user,nil
}

// func (user *User) GetById(db *sql.DB)(*User,error){
	
// }