package model

import (
	"time"
	
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID string 'json:"id" valid:"uuid"'
	CreatedAt time.time 'json:"created_at" valid:"-"'
	UpdatedAt time.Time 'json:"updated_at" valid:"-"'
}