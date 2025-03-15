package model

type Role string

const (
	RoleUser  = "User"
	RoleAdmin = "Admin"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
