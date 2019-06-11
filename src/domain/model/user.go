package model

// UserTable table name in db
const UserTable = "user"

// User info
type User struct {
	Base
	Name           string
	Email          string
	HashedPassword string
}
