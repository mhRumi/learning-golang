package models

import "time"

type Education struct {
	Id           int64     `json:"id"`
	Degree       string    `json:"degree"`
	Institute    string    `json:"institute"`
	Subject      string    `json:"subject"`
	Location     string    `json:"location"`
	Website_Link string    `json:"website_link"`
	Joining_Date time.Time `json:"joining_date"`
	Leaving_Date time.Time `json:"leaving_date"`
	Description  string    `json:"description"`
	Created_At   time.Time `json:"created_at"`
	Updated_At   time.Time `json:"updated_at"`
	Reg_No       int64     `json:"reg_no"`
}
