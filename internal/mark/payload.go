package mark

import "gorm.io/gorm"

type MarkRequest struct {
	gorm.Model
	Mark string `json:"mark"`
}
