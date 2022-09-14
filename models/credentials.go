package models

import "time"

type Credential struct {
	Id        int
	Role      string
	Email     string
	Reg_No    int64
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
