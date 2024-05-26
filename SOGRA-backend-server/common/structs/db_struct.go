package common_struct

type User struct {
	User_id           int `gorm:"unique;primaryKey;autoIncrement"`
	Id                string
	Password          string
	Email             string
	Profile_image     string
	Saved_event_list  int
	Saved_course_list int
}

type User_saved_event struct {
	User_id  int `gorm:"unique;primaryKey;autoIncrement"`
	Event_id string
}

type Course struct {
	Course_id  int `gorm:"unique;primaryKey;autoIncrement"`
	Info       int
	Event_list string
}

type User_saved_course struct {
	User_id     int `gorm:"unique;primaryKey;autoIncrement"`
	Course_list string
}

type Event struct {
	Event_id    int `gorm:"unique;primaryKey;autoIncrement"`
	Info        int
	Location    int
	Frame_info  int
	Review_info int
}

type Event_info struct {
	Info_id           int `gorm:"unique;primaryKey;autoIncrement"`
	Start_date        string
	End_date          string
	Weekday_time      string
	Weekend_time      string
	Specific_location string
}

type Info_frame struct {
	Info_frame_id int `gorm:"unique;primaryKey;autoIncrement"`
	Title         string
	Subtle        string
	Context       string
	Thumbnail     string
	Photo         string
}

type Open_area struct {
	Area_id     int
	Area_korean string
}

type Review_frame struct {
	Review_frame_id int `gorm:"unique;primaryKey;autoIncrement"`
	Total_score     int
	Vote_count      int
	Comment_list    string
}

type Review_comment struct {
	Comment_id int `gorm:"unique;primaryKey;autoIncrement"`
	User_id    int
	Context    string
}
