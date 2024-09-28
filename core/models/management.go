package models

type Management struct {
	Model
	Name            string
	Accounts        []Account
	Transactions    []Transaction
	ManagementUsers []ManagementUser
}

type ManagementUser struct {
	Model
	ManagementID int
	UserID       int
	User         User
}
