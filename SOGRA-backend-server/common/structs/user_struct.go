package common_struct

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}

type User struct {
	User_id           int
	Id                string
	Password          string
	Email             string
	Profile_image     string
	Saved_event_list  int
	Saved_course_list int
}

type GetUsersSaveResponse struct {
	SaveEventList []int `json:"save_event_list"`
}

type GetUserProfileImageResponse struct {
	ProfileImageURL string `json:"profile_image_url"`
}

type GetUserNicknameResponse struct {
	Nickname string `json:"nickname"`
}
