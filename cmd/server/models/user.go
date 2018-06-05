package models

// User structure
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserModel interface for user repo
type UserModel interface {
	GetData() []User
}

// MockUserModel for debug
type MockUserModel struct{}

// GetData to get mockup data
func (u *MockUserModel) GetData() []User {
	return []User{
		User{
			Username: "Arm",
			Email:    "registbacktown@gmail.com",
		},
	}
}

// ProdUserModel for release
type ProdUserModel struct{}

// GetData to get real data
func (u *ProdUserModel) GetData() []User {
	return []User{
		User{
			Username: "Saran",
			Email:    "saran@mm.co.th",
		},
	}
}
