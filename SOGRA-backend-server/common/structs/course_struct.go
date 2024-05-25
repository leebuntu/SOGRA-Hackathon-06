package common_struct

type CreateCourseRequest struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ThumbnailURL string `json:"thumbnail"`
}

type GetPublicCourseResponse struct {
	CourseList []int `json:"course_list"`
}

type GetCourseOutlineResponse struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	CreateUserID string `json:"create_user_id"`
	ThumbnailURL string `json:"thumbnail"`
}

type GetCourseFullResponse struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	CreateUserID string `json:"create_user_id"`
	ThumbnailURL string `json:"thumbnail"`
}
