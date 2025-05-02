package model

type StudentCourse struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	CourseID  int `json:"course_id"`
}

func (StudentCourse) TableName() string {
	return "student_course"
}
