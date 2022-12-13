package entities

import "time"

type Mahasiswa struct {
	ID        int        `json:"id" gorm:"primary_key"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (m *Mahasiswa) TableName() string {
	return "mahasiswa"
}
