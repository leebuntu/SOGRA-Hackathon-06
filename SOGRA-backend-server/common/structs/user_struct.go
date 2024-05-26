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

type GetUsersSaveResponse struct {
	SaveEventList []int `json:"save_event_list"`
}

type GetUserProfileImageResponse struct {
	ProfileImageURL string `json:"profile_image_url"`
}

type GetUserNicknameResponse struct {
	Nickname string `json:"nickname"`
}
