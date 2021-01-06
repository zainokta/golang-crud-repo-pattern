package models

import "time"

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Avatar    string     `json:"avatar"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type IUserRepository interface {
	Get() ([]User, error)
	Find(id int64) (User, error)
	Create(user *User) (int64, error)
	Update(id int64, user *User) error
	Delete(id int64) error
	FindUserByUsername(username string) (User, error)
}
