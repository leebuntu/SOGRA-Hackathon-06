package common

type LoginRequest struct {
	ID       string `json:"id"`
	password string `json:"password"`
}

type RegisterRequest struct {
	ID       string `json:"id"`
	password string `json:"password"`
	email    string `json:"email"`
	profile  string `json:"profile"`
}

type GetUsersSaveRequest struct {
}

type PostUsersSaveRequest struct {
}

type DeleteUsersSaveRequest struct {
}
