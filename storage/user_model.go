package storage

type User struct {
	ID          int64
	FirstName   string
	LastName    string
	Age         int32
	PhoneNumber string
}

type GetAllUsersParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}

type UserStorageI interface {
	Create(u *User) (*User)
	Get(id int64) (*User)
	GetAll(params *GetAllUsersParams) (*GetAllUsersResult)
	Update(u *User) (*User)
	Delete(id int64) error
}
