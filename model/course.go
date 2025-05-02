package model

import "time"

type Course struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Capsize int    `json:"capsize"`
	Chooses int
	Begin   time.Time `json:"begin"`
	End     time.Time `json:"end"`
}

func (Course) TableName() string {
	return "course"
}
