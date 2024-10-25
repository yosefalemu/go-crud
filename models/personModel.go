package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID      string   `gorm:"type:uuid;primarykey;default:uuid_generate_v4()" json:"id"`
	Name    string   `gorm:"type:varchar(255);not null" json:"name"`
	Age     int      `gorm:"not null" json:"age"`
	Hobbies []string `gorm:"type:text[]" json:"hobbies"`
}

func (p *Person) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		id, uuidErr := uuid.NewUUID()
		if uuidErr != nil {
			return uuidErr
		}
		p.ID = id.String()
	}
	return nil
}
