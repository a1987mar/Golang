package mark

import "gorm.io/gorm"

type Mark struct {
	gorm.Model
	Mark string `json:"mark"`
}

func NewMark(mark string) *Mark {
	return &Mark{
		Mark: mark,
	}
}
