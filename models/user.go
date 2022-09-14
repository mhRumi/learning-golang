package models

import "time"

type User struct {
	Reg_No        int64     `json:"reg_no"`
	Name          string    `json:"name"`
	Nick_Name     string    `json:"nick_name"`
	Address       string    `json:"address"`
	Biography     string    `json:"biography"`
	Image         string    `json:"image"`
	Fb_Link       string    `json:"fb_link"`
	Batch         int64     `json:"batch"`
	Linkedin_Link string    `json:"linkedin_link"`
	Git_Link      string    `json:"git_link"`
	Phone         string    `json:"phone"`
	Date_of_Birth time.Time `json:"date_of_birth"`
	Skills        []string  `json:"skills"`
}
